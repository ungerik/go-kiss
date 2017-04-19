package kiss

import (
	"io/ioutil"
	"os"

	"github.com/ungerik/go-kiss/contenttype"
	"github.com/ungerik/go-kiss/httperr"
)

// LoadFile returns the content of a file
// or a httperr.NotFound if the file does not exist
func LoadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, httperr.NotFound(filename)
	}
	return ioutil.ReadAll(f)
}

// LoadPlainTextFile returns the content of a file as contenttype.PlainText
// or a httperr.NotFound if the file does not exist
func LoadPlainTextFile(filename string) (contenttype.PlainText, error) {
	return LoadFile(filename)
}

// LoadPDFFile returns the content of a file as contenttype.PDF
// or a httperr.NotFound if the file does not exist
func LoadPDFFile(filename string) (contenttype.PDF, error) {
	return LoadFile(filename)
}

// LoadXMLFile returns the content of a file as contenttype.XML
// or a httperr.NotFound if the file does not exist
func LoadXMLFile(filename string) (contenttype.XML, error) {
	return LoadFile(filename)
}

// LoadJSONFile returns the content of a file as contenttype.JSON
// or a httperr.NotFound if the file does not exist
func LoadJSONFile(filename string) (contenttype.JSON, error) {
	return LoadFile(filename)
}
