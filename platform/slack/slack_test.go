package slack

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/slack-go/slack/slackevents"
)

func TestStripAppMentionText(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "strips bot mention prefix",
			in:   "<@U0BOT123> run tests",
			want: "run tests",
		},
		{
			name: "empty mention becomes empty text",
			in:   "<@U0BOT123> ",
			want: "",
		},
		{
			name: "plain text remains unchanged",
			in:   "run tests",
			want: "run tests",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripAppMentionText(tt.in); got != tt.want {
				t.Fatalf("stripAppMentionText(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestDownloadSlackFile_HTMLDetection(t *testing.T) {
	// Test that we detect HTML responses (Slack login page) and return an error
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate Slack returning HTML login page when auth is missing
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<!DOCTYPE html><html><body>Please login</body></html>"))
	}))
	defer ts.Close()

	p := &Platform{botToken: "xoxb-test-token"}
	_, err := p.downloadSlackFile(ts.URL)
	if err == nil {
		t.Fatal("expected error for HTML response, got nil")
	}
	// Should detect HTML prefix
	if err != nil && err.Error() == "" {
		t.Fatal("expected non-empty error message")
	}
}

func TestDownloadSlackFile_MissingAuth(t *testing.T) {
	// Test that we return an error for non-200 status codes
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
	}))
	defer ts.Close()

	p := &Platform{botToken: "xoxb-test-token"}
	_, err := p.downloadSlackFile(ts.URL)
	if err == nil {
		t.Fatal("expected error for 401 response, got nil")
	}
}

func TestDownloadSlackFile_Success(t *testing.T) {
	// Test successful binary download
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify Authorization header is set
		auth := r.Header.Get("Authorization")
		if auth != "Bearer xoxb-test-token" {
			t.Errorf("expected Authorization header 'Bearer xoxb-test-token', got %q", auth)
		}
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("\x89PNG\r\n\x1a\n")) // PNG magic bytes
	}))
	defer ts.Close()

	p := &Platform{botToken: "xoxb-test-token"}
	data, err := p.downloadSlackFile(ts.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(data) != 8 {
		t.Errorf("expected 8 bytes, got %d", len(data))
	}
}

func TestDownloadSlackFile_EmptyURL(t *testing.T) {
	p := &Platform{botToken: "xoxb-test-token"}
	_, err := p.downloadSlackFile("")
	if err == nil {
		t.Fatal("expected error for empty URL, got nil")
	}
}

func TestParseSlackInnerEventFiles(t *testing.T) {
	raw := json.RawMessage(`{"type":"app_mention","user":"U1","text":"<@B> hi","files":[{"id":"F1","name":"a.pdf","mimetype":"application/pdf","url_private_download":"http://example/f"}]}`)
	files := parseSlackInnerEventFiles(&raw)
	if len(files) != 1 {
		t.Fatalf("len(files) = %d, want 1", len(files))
	}
	if files[0].Name != "a.pdf" || files[0].Mimetype != "application/pdf" {
		t.Fatalf("unexpected file: %+v", files[0])
	}
}

func TestProcessSlackFileShares_GenericFile(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("%PDF-1.4 minimal"))
	}))
	defer ts.Close()

	p := &Platform{botToken: "xoxb-test"}
	images, audio, docs := p.processSlackFileShares([]slackevents.File{
		{
			ID:                 "Fpdf",
			Name:               "doc.pdf",
			Mimetype:           "application/pdf",
			URLPrivateDownload: ts.URL,
		},
	})
	if len(images) != 0 || audio != nil {
		t.Fatalf("expected only doc file, got images=%d audio=%v", len(images), audio)
	}
	if len(docs) != 1 {
		t.Fatalf("len(docs) = %d, want 1", len(docs))
	}
	if docs[0].FileName != "doc.pdf" || docs[0].MimeType != "application/pdf" {
		t.Fatalf("unexpected doc: %+v", docs[0])
	}
	if string(docs[0].Data) != "%PDF-1.4 minimal" {
		t.Fatalf("unexpected data %q", docs[0].Data)
	}
}

func TestProcessSlackFileShares_ImageVsDoc(t *testing.T) {
	imgSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("fakepng"))
	}))
	defer imgSrv.Close()
	txtSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	}))
	defer txtSrv.Close()

	p := &Platform{botToken: "xoxb-test"}
	images, audio, docs := p.processSlackFileShares([]slackevents.File{
		{ID: "1", Name: "x.png", Mimetype: "image/png", URLPrivateDownload: imgSrv.URL},
		{ID: "2", Name: "n.txt", Mimetype: "text/plain", URLPrivateDownload: txtSrv.URL},
	})
	if audio != nil {
		t.Fatal("unexpected audio")
	}
	if len(images) != 1 || len(docs) != 1 {
		t.Fatalf("want 1 image 1 doc, got images=%d docs=%d", len(images), len(docs))
	}
	if images[0].MimeType != "image/png" {
		t.Errorf("image mime: %q", images[0].MimeType)
	}
	if docs[0].MimeType != "text/plain" || string(docs[0].Data) != "hello" {
		t.Errorf("unexpected text file: %+v", docs[0])
	}
}

func TestProcessSlackFileShares_EmptyMimeBecomesOctetStream(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{0, 1, 2})
	}))
	defer ts.Close()

	p := &Platform{botToken: "xoxb-test"}
	_, _, docs := p.processSlackFileShares([]slackevents.File{
		{ID: "z", Name: "blob.bin", Mimetype: "", URLPrivateDownload: ts.URL},
	})
	if len(docs) != 1 || docs[0].MimeType != "application/octet-stream" {
		t.Fatalf("got %+v", docs)
	}
}

func TestShouldDropSender(t *testing.T) {
	tests := []struct {
		name      string
		allowBots bool
		chans     map[string]bool
		selfID    string
		botID     string
		userID    string
		channel   string
		want      bool
	}{
		{"human, bots off", false, nil, "UME", "", "UHUMAN", "CPUB", false},
		{"human, bots on", true, nil, "UME", "", "UHUMAN", "CPUB", false},
		{"missing user is always dropped", true, nil, "UME", "B1", "", "CPUB", true},
		{"bot, bots off", false, nil, "UME", "B1", "UQM", "CPUB", true},
		{"bot, bots on, no channel restriction", true, nil, "UME", "B1", "UQM", "CPUB", false},
		{"bot, bots on, in allowed channel", true, map[string]bool{"CPRIV": true}, "UME", "B1", "UQM", "CPRIV", false},
		{"bot, bots on, outside allowed channel", true, map[string]bool{"CPRIV": true}, "UME", "B1", "UQM", "CPUB", true},
		{"human unaffected by channel restriction", true, map[string]bool{"CPRIV": true}, "UME", "", "UHUMAN", "CPUB", false},
		{"own reply echoed back, bots on", true, nil, "UME", "B1", "UME", "CPRIV", true},
		{"own reply echoed back, in allowed channel", true, map[string]bool{"CPRIV": true}, "UME", "B1", "UME", "CPRIV", true},
		{"bot, bots on, self id unknown", true, nil, "", "B1", "UQM", "CPRIV", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Platform{allowBots: tt.allowBots, allowBotsChans: tt.chans, selfUserID: tt.selfID}
			if got := p.shouldDropSender(tt.botID, tt.userID, tt.channel); got != tt.want {
				t.Fatalf("shouldDropSender(%q,%q,%q)=%v want %v", tt.botID, tt.userID, tt.channel, got, tt.want)
			}
		})
	}
}

func TestIsSelfMention(t *testing.T) {
	p := &Platform{selfUserID: "UBOT"}
	if !p.isSelfMention("<@UBOT> list files") {
		t.Fatal("expected mention of self to be detected")
	}
	if p.isSelfMention("<@UOTHER> hi <@UBOTX>") {
		t.Fatal("other user IDs must not match")
	}
	if (&Platform{}).isSelfMention("<@UBOT> hi") {
		t.Fatal("unknown self ID must never match")
	}
}

func TestShouldIgnoreUnmentioned(t *testing.T) {
	p := &Platform{requireMention: true}
	p.markThreadActive("CPRIV", "1.000")
	cases := []struct {
		name                       string
		channelType, channel, thTS string
		want                       bool
	}{
		{"DM never gated", "im", "DPRIV", "", false},
		{"top-level channel message dropped", "group", "CPRIV", "", true},
		{"reply in unknown thread dropped", "group", "CPRIV", "9.999", true},
		{"reply in active thread passes", "group", "CPRIV", "1.000", false},
		{"same ts in other channel dropped", "channel", "CPUB", "1.000", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := p.shouldIgnoreUnmentioned(c.channelType, c.channel, c.thTS); got != c.want {
				t.Fatalf("got %v want %v", got, c.want)
			}
		})
	}
	off := &Platform{}
	if off.shouldIgnoreUnmentioned("group", "CPRIV", "") {
		t.Fatal("require_mention off must never gate")
	}
}

func TestFinalBelowCard(t *testing.T) {
	var m sync.Map
	c := &slackStreamingCard{lastPost: &m, channel: "C1", threadTS: "1.0"}
	if c.finalBelowCard() {
		t.Fatal("unposted card must edit in place")
	}
	c.ts = "1.1"
	recordPost(&m, "C1", "1.0", "1.1")
	if c.finalBelowCard() {
		t.Fatal("card is the latest post; edit in place")
	}
	recordPost(&m, "C1", "1.0", "1.2") // e.g. a permission prompt posted below
	if !c.finalBelowCard() {
		t.Fatal("something was posted after the card; final must go below")
	}
}

func TestActiveThreadsPersistence(t *testing.T) {
	dir := t.TempDir()
	p := &Platform{requireMention: true, activeThreads: map[string]int64{}, activeThreadsPath: activeThreadsPath(dir)}
	p.markThreadActive("C1", "1.000")
	// expired entry must be dropped on reload
	p.activeMu.Lock()
	p.activeThreads["C1:0.001"] = time.Now().Add(-activeThreadRetention - time.Hour).Unix()
	p.saveActiveThreadsLocked()
	p.activeMu.Unlock()

	q := &Platform{requireMention: true, activeThreads: map[string]int64{}, activeThreadsPath: p.activeThreadsPath}
	n, err := q.loadActiveThreads()
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("loaded %d entries, want 1 (expired one dropped)", n)
	}
	if q.shouldIgnoreUnmentioned("group", "C1", "1.000") {
		t.Fatal("thread persisted across restart must stay active")
	}
	if !q.shouldIgnoreUnmentioned("group", "C1", "0.001") {
		t.Fatal("expired thread must not be active")
	}
	if activeThreadsPath("") != "" {
		t.Fatal("empty data dir must disable persistence")
	}
}
