// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// MpercolateOption is a non-required Mpercolate option that gets applied to an HTTP request.
type MpercolateOption func(r *transport.Request)

// WithMpercolateIndex - the index of the document being count percolated to use as default.
func WithMpercolateIndex(index string) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// WithMpercolateType - the type of the document being percolated to use as default.
func WithMpercolateType(documentType string) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// WithMpercolateAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithMpercolateAllowNoIndices(allowNoIndices bool) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// MpercolateExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type MpercolateExpandWildcards int

const (
	// MpercolateExpandWildcardsOpen can be used to set MpercolateExpandWildcards to "open"
	MpercolateExpandWildcardsOpen = iota
	// MpercolateExpandWildcardsClosed can be used to set MpercolateExpandWildcards to "closed"
	MpercolateExpandWildcardsClosed = iota
	// MpercolateExpandWildcardsNone can be used to set MpercolateExpandWildcards to "none"
	MpercolateExpandWildcardsNone = iota
	// MpercolateExpandWildcardsAll can be used to set MpercolateExpandWildcards to "all"
	MpercolateExpandWildcardsAll = iota
)

// WithMpercolateExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithMpercolateExpandWildcards(expandWildcards MpercolateExpandWildcards) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// WithMpercolateIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithMpercolateIgnoreUnavailable(ignoreUnavailable bool) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// WithMpercolateErrorTrace - include the stack trace of returned errors.
func WithMpercolateErrorTrace(errorTrace bool) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// WithMpercolateFilterPath - a comma-separated list of filters used to reduce the respone.
func WithMpercolateFilterPath(filterPath []string) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// WithMpercolateHuman - return human readable values for statistics.
func WithMpercolateHuman(human bool) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// WithMpercolateIgnore - ignores the specified HTTP status codes.
func WithMpercolateIgnore(ignore []int) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// WithMpercolatePretty - pretty format the returned JSON response.
func WithMpercolatePretty(pretty bool) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// WithMpercolateSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithMpercolateSourceParam(sourceParam string) MpercolateOption {
	return func(r *transport.Request) {
	}
}

// Mpercolate - for indices created on or after version 5.0.0-alpha1 the percolator automatically indexes the query terms with the percolator queries. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-percolate.html for more info.
//
// body: the percolate request definitions (header & body pair), separated by newlines.
//
// options: optional parameters.
func (a *API) Mpercolate(body []interface{}, options ...MpercolateOption) (*MpercolateResponse, error) {
	req := a.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &MpercolateResponse{resp}, err
}

// MpercolateResponse is the response for Mpercolate.
type MpercolateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *MpercolateResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}