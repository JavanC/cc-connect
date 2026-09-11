package slack

import (
	"regexp"
	"strings"

	"github.com/slack-go/slack"
)

// Rich-text rendering for Slack.
//
// cc-connect historically sent every reply as a plain mrkdwn `text` payload.
// Slack renders that acceptably for inline styles, but lists, nested lists,
// quotes and code blocks end up as space-padded text rather than native list
// elements. Slack's Block Kit `rich_text` block gives real bullets, ordered
// lists with indentation, quotes and preformatted blocks, so the agent's
// Markdown is converted into rich_text elements here and attached as blocks.
// The mrkdwn `text` is still sent alongside as the notification/fallback body.
//
// The converter deliberately handles the subset of Markdown the agent is
// prompted to emit (headings, paragraphs, -/*/• and 1. lists with indentation,
// > quotes, ``` fences, tables, inline bold/italic/strike/code/links, Slack
// <@user> / <#channel> mentions). Anything else passes through as literal text.

// maxRichTextIndent is Slack's documented ceiling for rich_text_list.indent.
const maxRichTextIndent = 8

var (
	reRTFence     = regexp.MustCompile("^\\s*```")
	reRTHeading   = regexp.MustCompile(`^#{1,6}\s+(.*)$`)
	reRTListItem  = regexp.MustCompile(`^(\s*)(?:([-*•])|(\d+)[.)])\s+(.*)$`)
	reRTCheckbox  = regexp.MustCompile(`^\[([ xX])\]\s+(.*)$`)
	reRTQuote     = regexp.MustCompile(`^>\s?(.*)$`)
	reRTRule      = regexp.MustCompile(`^\s*(-{3,}|\*{3,}|_{3,})\s*$`)
	reRTTableLine = regexp.MustCompile(`^\s*\|.*\|\s*$`)

	// reRTInline finds the next inline construct. Group numbers:
	//  1 code span            `x`
	//  2 bold                 **x**
	//  3 bold (mrkdwn)        *x*
	//  4 italic               __x__
	//  5 italic               _x_
	//  6 strike               ~~x~~
	//  7 strike (mrkdwn)      ~x~
	//  8,9  markdown link     [text](url)
	// 10,11 slack link        <url|text> / <url>
	// 12 user mention         <@U123>
	// 13 channel mention      <#C123|name>
	// 14 bare URL
	reRTInline = regexp.MustCompile(
		"(`[^`\\n]+`)" +
			`|(\*\*[^*\n]+?\*\*)` +
			`|(\*[^*\n]+?\*)` +
			`|(__[^_\n]+?__)` +
			`|(_[^_\n]+?_)` +
			`|(~~[^~\n]+?~~)` +
			`|(~[^~\n]+?~)` +
			`|\[([^\]\n]+)\]\((https?://[^)\s]+)\)` +
			`|<(https?://[^|>\s]+)(?:\|([^>\n]*))?>` +
			`|<@([UW][A-Z0-9]+)>` +
			`|<#([CG][A-Z0-9]+)(?:\|[^>\n]*)?>` +
			`|(https?://[^\s<>]+)`,
	)
)

// MarkdownToRichTextBlocks converts agent Markdown into Slack Block Kit blocks.
// It returns nil when the input has no renderable content so callers can fall
// back to a text-only message.
func MarkdownToRichTextBlocks(md string) []slack.Block {
	elements := markdownToRichTextElements(md)
	if len(elements) == 0 {
		return nil
	}
	return []slack.Block{slack.NewRichTextBlock("", elements...)}
}

type rtList struct {
	style  slack.RichTextListElementType
	indent int
	items  []slack.RichTextElement
}

// rtBuilder accumulates block-level elements while scanning lines.
type rtBuilder struct {
	elements  []slack.RichTextElement
	paragraph []string
	quote     []string
	list      *rtList
	code      []string
	inCode    bool
	table     []string
}

func markdownToRichTextElements(md string) []slack.RichTextElement {
	b := &rtBuilder{}
	lines := strings.Split(strings.ReplaceAll(md, "\r\n", "\n"), "\n")
	for _, line := range lines {
		if b.inCode {
			if reRTFence.MatchString(line) {
				b.flushCode()
				continue
			}
			b.code = append(b.code, line)
			continue
		}
		if reRTFence.MatchString(line) {
			b.flushAll()
			b.inCode = true
			continue
		}
		trimmed := strings.TrimSpace(line)
		switch {
		case trimmed == "":
			b.flushAll()
		case reRTRule.MatchString(line):
			b.flushAll()
		case reRTTableLine.MatchString(line):
			b.flushParagraph()
			b.flushQuote()
			b.flushList()
			b.table = append(b.table, trimmed)
		case reRTHeading.MatchString(trimmed):
			b.flushAll()
			m := reRTHeading.FindStringSubmatch(trimmed)
			text := strings.TrimSpace(m[1])
			b.elements = append(b.elements, slack.NewRichTextSection(
				slack.NewRichTextSectionTextElement(text+"\n", &slack.RichTextSectionTextStyle{Bold: true}),
			))
		case reRTQuote.MatchString(trimmed):
			b.flushParagraph()
			b.flushList()
			b.flushTable()
			m := reRTQuote.FindStringSubmatch(trimmed)
			b.quote = append(b.quote, m[1])
		case reRTListItem.MatchString(line):
			b.flushParagraph()
			b.flushQuote()
			b.flushTable()
			m := reRTListItem.FindStringSubmatch(line)
			indent := listIndent(m[1])
			style := slack.RTEListBullet
			if m[3] != "" {
				style = slack.RTEListOrdered
			}
			item := m[4]
			if cb := reRTCheckbox.FindStringSubmatch(item); cb != nil {
				mark := "☐ "
				if cb[1] != " " {
					mark = "☑ "
				}
				item = mark + cb[2]
			}
			if b.list == nil || b.list.style != style || b.list.indent != indent {
				b.flushList()
				b.list = &rtList{style: style, indent: indent}
			}
			b.list.items = append(b.list.items, slack.NewRichTextSection(parseInline(item)...))
		default:
			b.flushQuote()
			b.flushList()
			b.flushTable()
			b.paragraph = append(b.paragraph, trimmed)
		}
	}
	if b.inCode {
		// Unterminated fence (e.g. streaming mid-block): render what we have.
		b.flushCode()
	}
	b.flushAll()
	trimTrailingNewline(b.elements)
	return b.elements
}

func listIndent(leading string) int {
	n := 0
	for _, r := range leading {
		if r == '\t' {
			n += 2
		} else {
			n++
		}
	}
	indent := n / 2
	if indent > maxRichTextIndent {
		indent = maxRichTextIndent
	}
	return indent
}

func (b *rtBuilder) flushAll() {
	b.flushParagraph()
	b.flushQuote()
	b.flushList()
	b.flushTable()
}

func (b *rtBuilder) flushParagraph() {
	if len(b.paragraph) == 0 {
		return
	}
	text := strings.Join(b.paragraph, "\n")
	b.paragraph = nil
	els := parseInline(text)
	if len(els) == 0 {
		return
	}
	// Paragraph spacing: rich_text_section elements render back-to-back, so
	// close each paragraph with a newline. trimTrailingNewline removes the last.
	els = append(els, slack.NewRichTextSectionTextElement("\n", nil))
	b.elements = append(b.elements, slack.NewRichTextSection(els...))
}

func (b *rtBuilder) flushQuote() {
	if len(b.quote) == 0 {
		return
	}
	text := strings.Join(b.quote, "\n")
	b.quote = nil
	els := parseInline(text)
	if len(els) == 0 {
		return
	}
	q := slack.RichTextQuote(slack.RichTextSection{Type: slack.RTEQuote, Elements: els})
	b.elements = append(b.elements, &q)
}

func (b *rtBuilder) flushList() {
	if b.list == nil {
		return
	}
	if len(b.list.items) > 0 {
		b.elements = append(b.elements, slack.NewRichTextList(b.list.style, b.list.indent, b.list.items...))
	}
	b.list = nil
}

func (b *rtBuilder) flushCode() {
	b.inCode = false
	if len(b.code) == 0 {
		return
	}
	text := strings.Join(b.code, "\n")
	b.code = nil
	b.elements = append(b.elements, &slack.RichTextPreformatted{
		RichTextSection: slack.RichTextSection{
			Type:     slack.RTEPreformatted,
			Elements: []slack.RichTextSectionElement{slack.NewRichTextSectionTextElement(text, nil)},
		},
	})
}

// flushTable renders a Markdown table as a preformatted block so column
// alignment survives; Slack has no native table element in rich_text.
func (b *rtBuilder) flushTable() {
	if len(b.table) == 0 {
		return
	}
	text := strings.Join(b.table, "\n")
	b.table = nil
	b.elements = append(b.elements, &slack.RichTextPreformatted{
		RichTextSection: slack.RichTextSection{
			Type:     slack.RTEPreformatted,
			Elements: []slack.RichTextSectionElement{slack.NewRichTextSectionTextElement(text, nil)},
		},
	})
}

// trimTrailingNewline drops the paragraph-spacing newline from the final
// section so the message does not end with a blank line.
func trimTrailingNewline(elements []slack.RichTextElement) {
	if len(elements) == 0 {
		return
	}
	sec, ok := elements[len(elements)-1].(*slack.RichTextSection)
	if !ok || len(sec.Elements) == 0 {
		return
	}
	last, ok := sec.Elements[len(sec.Elements)-1].(*slack.RichTextSectionTextElement)
	if !ok {
		return
	}
	if last.Text == "\n" {
		sec.Elements = sec.Elements[:len(sec.Elements)-1]
	} else {
		last.Text = strings.TrimSuffix(last.Text, "\n")
	}
}

// parseInline converts one run of inline Markdown/mrkdwn into section elements.
func parseInline(text string) []slack.RichTextSectionElement {
	var out []slack.RichTextSectionElement
	var plain strings.Builder
	flushPlain := func() {
		if plain.Len() > 0 {
			out = append(out, slack.NewRichTextSectionTextElement(plain.String(), nil))
			plain.Reset()
		}
	}
	rest := text
	for rest != "" {
		loc := reRTInline.FindStringSubmatchIndex(rest)
		if loc == nil {
			plain.WriteString(rest)
			break
		}
		start, end := loc[0], loc[1]
		group := func(i int) string {
			if loc[2*i] < 0 {
				return ""
			}
			return rest[loc[2*i]:loc[2*i+1]]
		}
		// Underscore emphasis must not fire inside identifiers (snake_case).
		if (group(4) != "" || group(5) != "") && !emphasisBoundaryOK(rest, start, end) {
			plain.WriteString(rest[:start+1])
			rest = rest[start+1:]
			continue
		}
		plain.WriteString(rest[:start])
		flushPlain()
		switch {
		case group(1) != "":
			out = append(out, slack.NewRichTextSectionTextElement(strings.Trim(group(1), "`"), &slack.RichTextSectionTextStyle{Code: true}))
		case group(2) != "":
			out = append(out, slack.NewRichTextSectionTextElement(strings.TrimSuffix(strings.TrimPrefix(group(2), "**"), "**"), &slack.RichTextSectionTextStyle{Bold: true}))
		case group(3) != "":
			out = append(out, slack.NewRichTextSectionTextElement(strings.Trim(group(3), "*"), &slack.RichTextSectionTextStyle{Bold: true}))
		case group(4) != "":
			out = append(out, slack.NewRichTextSectionTextElement(strings.TrimSuffix(strings.TrimPrefix(group(4), "__"), "__"), &slack.RichTextSectionTextStyle{Italic: true}))
		case group(5) != "":
			out = append(out, slack.NewRichTextSectionTextElement(strings.Trim(group(5), "_"), &slack.RichTextSectionTextStyle{Italic: true}))
		case group(6) != "":
			out = append(out, slack.NewRichTextSectionTextElement(strings.TrimSuffix(strings.TrimPrefix(group(6), "~~"), "~~"), &slack.RichTextSectionTextStyle{Strike: true}))
		case group(7) != "":
			out = append(out, slack.NewRichTextSectionTextElement(strings.Trim(group(7), "~"), &slack.RichTextSectionTextStyle{Strike: true}))
		case group(9) != "":
			out = append(out, slack.NewRichTextSectionLinkElement(group(9), group(8), nil))
		case group(10) != "":
			label := group(11)
			if label == "" {
				label = group(10)
			}
			out = append(out, slack.NewRichTextSectionLinkElement(group(10), label, nil))
		case group(12) != "":
			out = append(out, slack.NewRichTextSectionUserElement(group(12), nil))
		case group(13) != "":
			// slack-go v0.16.0's constructor sets Type to "text" by mistake;
			// Slack rejects that, so set the correct element type explicitly.
			ch := slack.NewRichTextSectionChannelElement(group(13), nil)
			ch.Type = slack.RTSEChannel
			out = append(out, ch)
		case group(14) != "":
			url := strings.TrimRight(group(14), ".,;:!?)")
			trail := group(14)[len(url):]
			out = append(out, slack.NewRichTextSectionLinkElement(url, url, nil))
			plain.WriteString(trail)
		default:
			plain.WriteString(rest[start:end])
		}
		rest = rest[end:]
	}
	flushPlain()
	return out
}

// emphasisBoundaryOK reports whether the underscore span rest[start:end] is
// delimited by non-word characters on both sides, so snake_case_names are
// left alone.
func emphasisBoundaryOK(s string, start, end int) bool {
	isWord := func(r byte) bool {
		return r == '_' || (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
	}
	if start > 0 && isWord(s[start-1]) {
		return false
	}
	if end < len(s) && isWord(s[end]) {
		return false
	}
	return true
}

// messageOptions builds the Slack message payload for agent content: mrkdwn
// text (fallback/notification body) plus, when rich is set and the content
// converts, a rich_text block for native rendering.
func messageOptions(mrkdwn string, blocks []slack.Block) []slack.MsgOption {
	opts := []slack.MsgOption{slack.MsgOptionText(mrkdwn, false)}
	if len(blocks) > 0 {
		opts = append(opts, slack.MsgOptionBlocks(blocks...))
	}
	return opts
}
