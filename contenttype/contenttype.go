package contenttype

import "net/http"

// Data is an interface for returning content-data
// together with a content-type.
type Data interface {
	// ContentType returns the content-type value
	ContentType() string

	// Content returns the data
	Content() []byte
}

// WrapData wraps content together with a content-type
func WrapData(content []byte, contentType string) Data {
	return &wrapper{content, contentType}
}

type wrapper struct {
	content     []byte
	contentType string
}

// ContentType returns the content-type value
func (c *wrapper) ContentType() string {
	return c.contentType
}

// Content returns the data
func (c *wrapper) Content() []byte {
	return c.content
}

// Detect uses http.DetectContentType
type Detect []byte

// ContentType returns the content-type value
func (d Detect) ContentType() string {
	return http.DetectContentType(d)
}

// Content returns the data
func (d Detect) Content() []byte {
	return d
}

// XML content-type
type XML []byte

// ContentType returns the content-type value
func (XML) ContentType() string {
	return "application/xml; charset=utf-8"
}

// Content returns the data
func (x XML) Content() []byte {
	return x
}

// JSON content-type
type JSON []byte

// ContentType returns the content-type value
func (JSON) ContentType() string {
	return "application/json; charset=utf-8"
}

// Content returns the data
func (j JSON) Content() []byte {
	return j
}

// HTML content-type
type HTML []byte

// ContentType returns the content-type value
func (HTML) ContentType() string {
	return "text/html; charset=utf-8"
}

// Content returns the data
func (h HTML) Content() []byte {
	return h
}

// PlainText content-type
type PlainText []byte

// ContentType returns the content-type value
func (PlainText) ContentType() string {
	return "text/plain; charset=utf-8"
}

// Content returns the data
func (t PlainText) Content() []byte {
	return t
}

// JPEG content-type
type JPEG []byte

// ContentType returns the content-type value
func (JPEG) ContentType() string {
	return "image/jpeg"
}

// Content returns the data
func (j JPEG) Content() []byte {
	return j
}

// PDF content-type
type PDF []byte

// ContentType returns the content-type value
func (PDF) ContentType() string {
	return "application/pdf"
}

// Content returns the data
func (p PDF) Content() []byte {
	return p
}

// CSV content-type
type CSV []byte

// ContentType returns the content-type value
func (CSV) ContentType() string {
	return "text/csv; charset=utf-8"
}

// Content returns the data
func (t CSV) Content() []byte {
	return t
}
