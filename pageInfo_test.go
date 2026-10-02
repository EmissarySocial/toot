package toot

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPageInfo_NoPages(t *testing.T) {

	header := make(http.Header)
	PageInfo{}.SetHeader(header, "/api/v1/notifications")

	if got := header.Get("Link"); got != "" {
		t.Errorf("expected no Link header, got %q", got)
	}
}

func TestPageInfo_NextOnly(t *testing.T) {

	header := make(http.Header)
	PageInfo{MaxID: "ZZZ"}.SetHeader(header, "/api/v1/notifications")

	expected := `</api/v1/notifications?max_id=ZZZ>; rel="next"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestPageInfo_PreviousOnly(t *testing.T) {

	header := make(http.Header)
	PageInfo{MinID: "AAA"}.SetHeader(header, "/api/v1/notifications")

	expected := `</api/v1/notifications?min_id=AAA>; rel="prev"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

// Both links must survive. A "Set" per link would leave only the last one.
func TestPageInfo_BothPagesSurvive(t *testing.T) {

	header := make(http.Header)
	PageInfo{MinID: "AAA", MaxID: "ZZZ"}.SetHeader(header, "/api/v1/notifications")

	expected := `</api/v1/notifications?max_id=ZZZ>; rel="next", </api/v1/notifications?min_id=AAA>; rel="prev"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}

	if count := len(header.Values("Link")); count != 1 {
		t.Errorf("expected exactly 1 Link field, got %d", count)
	}
}

func TestPageInfo_OffsetIsNextPage(t *testing.T) {

	header := make(http.Header)
	PageInfo{Offset: "40"}.SetHeader(header, "/api/v1/directory")

	expected := `</api/v1/directory?offset=40>; rel="next"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

// MaxID wins over Offset, because a cursor is more precise than a count.
func TestPageInfo_MaxIDBeatsOffset(t *testing.T) {

	page := PageInfo{MaxID: "ZZZ", Offset: "40"}

	if got := page.GetNextPage("/x"); got != "/x?max_id=ZZZ" {
		t.Errorf("expected max_id to win, got %q", got)
	}
}

// Each link writes its own paging cursor, replacing the request's; other parameters are kept.
func TestPageInfo_QueryStringIsReplaced(t *testing.T) {

	header := make(http.Header)
	PageInfo{MaxID: "ZZZ"}.SetHeader(header, "/api/v1/notifications?max_id=OLD&limit=20")

	expected := `</api/v1/notifications?max_id=ZZZ&limit=20>; rel="next"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

// An ID carrying URL metacharacters must not break out of the link.
func TestPageInfo_IDsAreEscaped(t *testing.T) {

	header := make(http.Header)
	PageInfo{MaxID: `a b>"&c`}.SetHeader(header, "/api/v1/notifications")

	expected := `</api/v1/notifications?max_id=a+b%3E%22%26c>; rel="next"`

	if got := header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestPageInfo_NilHeader(_ *testing.T) {
	PageInfo{MaxID: "ZZZ"}.SetHeader(nil, "/x") // must not panic
}

func TestPageInfo_SetHeadersNilGuards(_ *testing.T) {

	PageInfo{MaxID: "ZZZ"}.SetHeaders(nil)
	PageInfo{MaxID: "ZZZ"}.SetHeaders(&http.Response{Header: make(http.Header)})
	PageInfo{MaxID: "ZZZ"}.SetHeaders(&http.Response{
		Header:  make(http.Header),
		Request: &http.Request{},
	})
}

func TestPageInfo_SetHeadersFromClientResponse(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "https://example.com/api/v1/notifications?limit=20", nil)
	response := &http.Response{Header: make(http.Header), Request: request}

	PageInfo{MinID: "AAA", MaxID: "ZZZ"}.SetHeaders(response)

	// The request's own limit stays on each link, as Mastodon's does
	expected := `<https://example.com/api/v1/notifications?max_id=ZZZ&limit=20>; rel="next", ` +
		`<https://example.com/api/v1/notifications?min_id=AAA&limit=20>; rel="prev"`

	if got := response.Header.Get("Link"); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

// TestPageInfo_SetHeader_KeepsFilters confirms a request's filters carry onto the next and previous
// links, while its own paging cursors are replaced.
func TestPageInfo_SetHeader_KeepsFilters(t *testing.T) {

	header := make(http.Header)
	PageInfo{MaxID: "ZZZ", MinID: "AAA"}.SetHeader(header, "https://example.com/api/v1/accounts/1/statuses?only_media=true&exclude_replies=true&max_id=OLD&limit=20")

	link := header.Get("Link")

	expectedNext := `<https://example.com/api/v1/accounts/1/statuses?max_id=ZZZ&exclude_replies=true&limit=20&only_media=true>; rel="next"`
	expectedPrev := `<https://example.com/api/v1/accounts/1/statuses?min_id=AAA&exclude_replies=true&limit=20&only_media=true>; rel="prev"`

	if !strings.Contains(link, expectedNext) {
		t.Errorf("next link should keep the filters\nwant: %s\ngot:  %s", expectedNext, link)
	}

	if !strings.Contains(link, expectedPrev) {
		t.Errorf("previous link should keep the filters\nwant: %s\ngot:  %s", expectedPrev, link)
	}

	if strings.Contains(link, "max_id=OLD") {
		t.Errorf("the request's own cursor should be replaced, got: %s", link)
	}
}
