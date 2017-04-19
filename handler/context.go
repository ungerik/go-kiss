package handler

import (
	"net/http"

	"github.com/ungerik/go-kiss/contenttype"
)

// Context is the context for a request handler
type Context struct {
	RequestPath struct {
		LastPart    PathPart
		StringParts []string
	}
	Request  *http.Request
	Response http.ResponseWriter
}

// NewContext creates a new context for a request handler
func NewContext(lastPathPart PathPart, pathStringParts []string, writer http.ResponseWriter, request *http.Request) *Context {
	ctx := &Context{
		Request:  request,
		Response: writer,
	}
	ctx.RequestPath.LastPart = lastPathPart
	ctx.RequestPath.StringParts = pathStringParts
	return ctx
}

// GetRequestArgs from partStrings leading up to part
func (ctx *Context) GetRequestArgs() []string {
	lastPart := ctx.RequestPath.LastPart
	stringParts := ctx.RequestPath.StringParts
	if lastPart.HandlerFunc() == nil {
		panic("Context.GetRequestArgs must be called by a Handler")
	}
	// Start with lastPart.Parent() because this will be called by a Handler
	// with depth of len(partStrings)
	return getStringArgsRecursive(lastPart.Parent(), stringParts, len(stringParts)-1, nil)
}

func getStringArgsRecursive(part PathPart, partStrings []string, partDepth int, baseArgs []string) []string {
	if part == nil {
		// We have reached to parent of the root part
		return baseArgs
	}
	baseArgs = getStringArgsRecursive(part.Parent(), partStrings, partDepth-1, baseArgs)
	if part.IsArg() {
		baseArgs = append(baseArgs, partStrings[partDepth])
	}
	return baseArgs
}

// WriteResponse writes Data to a http.ResponseWriter
func (ctx *Context) WriteResponse(data contenttype.Data) {
	ctx.Response.Header().Add("Content-Type", data.ContentType())
	_, err := ctx.Response.Write(data.Content())
	if err != nil {
		// Don't return the error because it should only
		// haven in very rare cases (out of memory or canceled request)
		panic(err)
	}
}
