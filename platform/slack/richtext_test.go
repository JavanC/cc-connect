package slack

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/slack-go/slack"
)

// elementsOf renders the block to JSON and returns the top-level element types
// in order, which is the most robust way to assert structure across slack-go's
// custom marshalers.
func elementsOf(t *testing.T, blocks []slack.Block) []map[string]any {
	t.Helper()
	if len(blocks) != 1 {
		t.Fatalf("expected exactly one rich_text block, got %d", len(blocks))
	}
	raw, err := json.Marshal(blocks[0])
	if err != nil {
		t.Fatal(err)
	}
	var b struct {
		Type     string           `json:"type"`
		Elements []map[string]any `json:"elements"`
	}
	if err := json.Unmarshal(raw, &b); err != nil {
		t.Fatal(err)
	}
	if b.Type != "rich_text" {
		t.Fatalf("block type = %q, want rich_text", b.Type)
	}
	return b.Elements
}

func types(els []map[string]any) []string {
	out := make([]string, 0, len(els))
	for _, e := range els {
		out = append(out, e["type"].(string))
	}
	return out
}

func TestMarkdownToRichTextBlocks_Empty(t *testing.T) {
	if got := MarkdownToRichTextBlocks(""); got != nil {
		t.Fatalf("empty input should yield nil blocks, got %v", got)
	}
	if got := MarkdownToRichTextBlocks("\n\n"); got != nil {
		t.Fatalf("blank input should yield nil blocks, got %v", got)
	}
}

func TestMarkdownToRichTextBlocks_Structure(t *testing.T) {
	md := strings.Join([]string{
		"# Summary",
		"First paragraph with *bold* and `code`.",
		"",
		"- item one",
		"- item two",
		"  - nested",
		"1. first",
		"2. second",
		"> quoted line",
		"```",
		"go build ./...",
		"```",
		"| a | b |",
		"|---|---|",
		"Closing line.",
	}, "\n")
	els := elementsOf(t, MarkdownToRichTextBlocks(md))
	got := strings.Join(types(els), ",")
	want := "rich_text_section,rich_text_section,rich_text_list,rich_text_list,rich_text_list,rich_text_quote,rich_text_preformatted,rich_text_preformatted,rich_text_section"
	if got != want {
		t.Fatalf("element types\n got: %s\nwant: %s", got, want)
	}
	// Nested bullet gets indent 1; ordered list uses "ordered".
	if els[3]["indent"].(float64) != 1 || els[3]["style"] != "bullet" {
		t.Fatalf("nested list = %v", els[3])
	}
	if els[4]["style"] != "ordered" {
		t.Fatalf("ordered list style = %v", els[4]["style"])
	}
	// Heading is bold and last section has no trailing newline.
	head := els[0]["elements"].([]any)[0].(map[string]any)
	if head["style"].(map[string]any)["bold"] != true {
		t.Fatalf("heading should be bold: %v", head)
	}
	last := els[len(els)-1]["elements"].([]any)
	if txt := last[len(last)-1].(map[string]any)["text"].(string); strings.HasSuffix(txt, "\n") {
		t.Fatalf("final section must not end with newline: %q", txt)
	}
}

func TestMarkdownToRichTextBlocks_Checklist(t *testing.T) {
	els := elementsOf(t, MarkdownToRichTextBlocks("- [ ] todo\n- [x] done"))
	items := els[0]["elements"].([]any)
	first := items[0].(map[string]any)["elements"].([]any)[0].(map[string]any)["text"].(string)
	second := items[1].(map[string]any)["elements"].([]any)[0].(map[string]any)["text"].(string)
	if !strings.HasPrefix(first, "☐ ") || !strings.HasPrefix(second, "☑ ") {
		t.Fatalf("checkbox rendering: %q / %q", first, second)
	}
}

func TestParseInline(t *testing.T) {
	els := parseInline("see **bold**, _it_, ~~gone~~, `x`, [doc](https://a.b/c), <https://s.io|Site>, <@U123>, <#C456|gen>, https://raw.example/path, snake_case_name.")
	var kinds []string
	for _, e := range els {
		raw, _ := json.Marshal(e)
		var m map[string]any
		_ = json.Unmarshal(raw, &m)
		k := m["type"].(string)
		if st, ok := m["style"].(map[string]any); ok {
			for _, s := range []string{"bold", "italic", "strike", "code"} {
				if st[s] == true {
					k += ":" + s
				}
			}
		}
		kinds = append(kinds, k)
	}
	got := strings.Join(kinds, " ")
	want := "text text:bold text text:italic text text:strike text text:code text link text link text user text channel text link text"
	if got != want {
		t.Fatalf("inline kinds\n got: %s\nwant: %s", got, want)
	}
	// snake_case must survive untouched in the trailing text element.
	raw, _ := json.Marshal(els[len(els)-1])
	if !strings.Contains(string(raw), "snake_case_name") {
		t.Fatalf("snake_case was mangled: %s", raw)
	}
}

func TestMessageOptions_TextOnlyWhenNoBlocks(t *testing.T) {
	if n := len(messageOptions("hi", nil)); n != 1 {
		t.Fatalf("expected 1 option without blocks, got %d", n)
	}
	if n := len(messageOptions("hi", MarkdownToRichTextBlocks("hi"))); n != 2 {
		t.Fatalf("expected text+blocks options, got %d", n)
	}
}
