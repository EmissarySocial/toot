package toot

import (
	"net/http"
	"net/url"
	"strings"
)

// PageInfo is used to add pagination headers to a response
type PageInfo struct {
	MinID  string // If present, then a <prev> link header will be created
	MaxID  string // If present, then a <next> link header will be created
	Offset string // If present, then a <next> link header will be created
}

// SetHeader writes this page's "Link" header into the provided header map, using
// path as the base URL for each link. Nothing is written when there is neither a
// next nor a previous page. This is the method a SERVER wants.
func (p PageInfo) SetHeader(header http.Header, path string) {

	// RULE: There is nowhere to write the header.
	if header == nil {
		return
	}

	// Each link supplies its own paging cursor, but keeps the request's other parameters (such as
	// only_media or exclude_replies) so that the next page is filtered the same way as this one.
	path, rawQuery, _ := strings.Cut(path, "?")
	filters, _ := url.ParseQuery(rawQuery)

	for _, cursor := range []string{"max_id", "min_id", "since_id", "offset"} {
		filters.Del(cursor)
	}

	links := make([]string, 0, 2)

	if next := p.GetNextPage(path); next != "" {
		links = append(links, "<"+withFilters(next, filters)+`>; rel="next"`)
	}

	if prev := p.GetPreviousPage(path); prev != "" {
		links = append(links, "<"+withFilters(prev, filters)+`>; rel="prev"`)
	}

	// RULE: A header with no links at all is not written.
	if len(links) == 0 {
		return
	}

	// RFC 8288 carries several link-values in ONE field, separated by commas.  Writing
	// them one at a time with Set would leave only the last one standing.
	header.Set("Link", strings.Join(links, ", "))
}

// SetHeaders adds pagination headers to a response.
//
// Note that http.Request.Response is only populated for a CLIENT following a
// redirect; it is always nil on the server side, where this call does nothing.
// Server code should call SetHeader instead.
func (p PageInfo) SetHeaders(response *http.Response) {

	// RULE: There is nowhere to write the header.
	if response == nil {
		return
	}

	// RULE: The originating request supplies the base URL of every link.
	if response.Request == nil {
		return
	}

	if response.Request.URL == nil {
		return
	}

	p.SetHeader(response.Header, response.Request.URL.String())
}

// GetPreviousPage returns a URL for the previous page of results.
// If there is no previous page, then an empty string is returned.
func (p PageInfo) GetPreviousPage(path string) string {

	if p.MinID == "" {
		return ""
	}

	return path + "?min_id=" + url.QueryEscape(p.MinID)
}

// GetNextPage returns a URL for the next page of results.
// If there is no next page, then an empty string is returned.
func (p PageInfo) GetNextPage(path string) string {

	if p.MaxID != "" {
		return path + "?max_id=" + url.QueryEscape(p.MaxID)
	}

	if p.Offset != "" {
		return path + "?offset=" + url.QueryEscape(p.Offset)
	}

	return ""
}

// withFilters adds query parameters to a link that already has its paging cursor.
func withFilters(link string, filters url.Values) string {

	if len(filters) == 0 {
		return link
	}

	return link + "&" + filters.Encode()
}
