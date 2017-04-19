package httperr

import (
	"net/http"
)

// Error is an interface for HTTP errors
type Error interface {
	error
	StatusCode() int
}

// NotFound 404
type NotFound string

func (e NotFound) Error() string {
	return "404 Not Found: " + string(e)
}

func (NotFound) StatusCode() int {
	return http.StatusNotFound
}

// Redirect is an interface for HTTP redirects returned as error for a request
type Redirect interface {
	error
	URL() string
	StatusCode() int
}

// TemporaryRedirect 307
type TemporaryRedirect string

func (url TemporaryRedirect) Error() string {
	return "TemporaryRedirect: " + string(url)
}

func (url TemporaryRedirect) URL() string {
	return string(url)
}

func (TemporaryRedirect) StatusCode() int {
	return http.StatusTemporaryRedirect
}

// PermanentRedirect 308
type PermanentRedirect string

func (url PermanentRedirect) Error() string {
	return "PermanentRedirect: " + string(url)
}

func (url PermanentRedirect) URL() string {
	return string(url)
}

func (PermanentRedirect) StatusCode() int {
	return http.StatusPermanentRedirect
}
