package kiss

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ungerik/go-kiss/handler"
	"github.com/ungerik/go-kiss/httperr"
)

var (
	// ShowInternalErrors determines if internal errors are
	// shown in public server responses.
	// Enable it for debugging.
	ShowInternalErrors = true
)

// Serve starts a server for pathTree listening at addr.
func Serve(pathTree handler.PathPart, addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: Handler(pathTree),
	}
	return server.ListenAndServe()
}

// ServeTLS starts a TLS server for pathTree listening at addr.
func ServeTLS(pathTree handler.PathPart, addr, certFile, keyFile string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: Handler(pathTree),
	}
	return server.ListenAndServeTLS(certFile, keyFile)
}

func handleError(e interface{}, writer http.ResponseWriter, request *http.Request) {
	var (
		statusString string
		statusCode   int
	)
	switch err := e.(type) {
	case httperr.Redirect:
		http.Redirect(writer, request, err.URL(), err.StatusCode())
		return
	case httperr.Error:
		statusString = err.Error()
		statusCode = err.StatusCode()
	case error:
		statusString = err.Error()
		statusCode = http.StatusInternalServerError
	default:
		statusString = fmt.Sprint(e)
		statusCode = http.StatusInternalServerError
	}

	log.Print(statusString)
	if !ShowInternalErrors {
		statusString = http.StatusText(statusCode)
	}
	http.Error(writer, statusString, statusCode)

}

// Handler returns the http.HandlerFunc for pathTree
func Handler(pathTree handler.PathPart) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Handler recovered panic:", r)
				handleError(r, writer, request)
			}
		}()

		log.Printf("%s: %s", request.Method, request.URL.Path)

		path := strings.TrimSuffix(request.URL.Path, "/")
		parts := strings.Split(path, "/")
		parts[0] = "/"
		handled := handlePartsRecursive(pathTree, pathTree, parts, 0, writer, request)
		if !handled {
			http.NotFound(writer, request)
		}
	}
}

func handlePartsRecursive(pathTree, currentPart handler.PathPart, parts []string, currentDepth int, writer http.ResponseWriter, request *http.Request) bool {
	// fmt.Printf("handlePartsRecursive(pathTree: %#v, parts: %#v, currentDepth: %v)\n", pathTree, parts, currentDepth)
	if currentDepth < len(parts) {
		if currentPart.Match(parts[currentDepth]) {
			for _, child := range currentPart.Children() {
				handled := handlePartsRecursive(pathTree, child, parts, currentDepth+1, writer, request)
				if handled {
					return handled
				}
			}
		}
	} else {
		// currentDepth is at index one higher than last element of parts,
		// so this must be a handler with the HTTP method as name that can be matched
		handlerFunc := currentPart.HandlerFunc()
		if handlerFunc != nil && currentPart.Match(request.Method) {
			handlerFunc(handler.NewContext(currentPart, parts, writer, request))
			return true
		}
	}
	return false
}

// PrintEndpoints prints the handler endpoints of pathTree to stdout
func PrintEndpoints(pathTree handler.PathPart) {
	printEndpointsRecursive(pathTree, nil, 0)
}

func printEndpointsRecursive(currentPart handler.PathPart, parts []string, currentDepth int) {
	if currentPart.HandlerFunc() != nil {
		method := currentPart.Name()
		path := parts[0] + strings.Join(parts[1:], "/")
		fmt.Println(method, path)
	} else {
		parts = append(parts, currentPart.String())
		for _, child := range currentPart.Children() {
			printEndpointsRecursive(child, parts, currentDepth+1)
		}
	}
}
