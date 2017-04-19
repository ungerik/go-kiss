// go run go-kiss-generate.go ../../ephesoft-results-server

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ungerik/go-fs"
)

func main() {
	var (
		webRootDirName   string
		generatedDirName string
	)

	flag.StringVar(&webRootDirName, "root", "root", "name of the root directory from where code for handling URLs should be generated, relative to the project directory")
	flag.StringVar(&generatedDirName, "generated", "generatednew", "name of the directory where the generated code will be placed, relative to the project directory")
	flag.Parse()

	if len(flag.Args()) == 0 {
		panic("Need project path as command line argument!")
	}

	projectPath, err := filepath.Abs(filepath.Clean(flag.Arg(0)))
	if err != nil {
		panic(err)
	}
	projectImportPath := strings.TrimPrefix(filepath.ToSlash(projectPath), goPathSrc())

	// fmt.Println("go-kiss-generate", projectPath, projectImportPath)

	webRootImportPath := path.Join(projectImportPath, webRootDirName)
	webRootPath := filepath.Join(projectPath, webRootDirName)
	generatedPath := filepath.Join(projectPath, generatedDirName)

	parseResults := parseDir(webRootPath, nil)
	packages := parseResults.onlyPackages()

	var sourceBuf bytes.Buffer
	sourceBuf.WriteString("package generated\n\n")
	sourceBuf.WriteString("import (\n")
	sourceBuf.WriteString("\t\"github.com/ungerik/go-kiss/handler\"\n\n")

	for _, p := range packages {
		importPath := path.Join(webRootImportPath, path.Join(p.pathParts...))
		fmt.Fprintf(&sourceBuf, "\t%s \"%s\"\n", p.pkg.Name, importPath)
	}
	sourceBuf.WriteString(")\n\n")

	sourceBuf.WriteString("// PathTree is the root path tree\n")
	sourceBuf.WriteString("var PathTree = handler.NewRoot(\n")

	sourceBuf.WriteString(")\n\n")

	sourceBuf.WriteString("func init() {\n")
	sourceBuf.WriteString("\thandler.InitTree(PathTree)\n")
	sourceBuf.WriteString("}\n")

	generatedDir, err := fs.MakeDir(generatedPath)
	if err != nil {
		panic(err)
	}
	generatedFile := generatedDir.Relative("pathtree.go")

	err = generatedFile.WriteAll(sourceBuf.Bytes())
	if err != nil {
		panic(err)
	}
	// q.Q(parseResults)
}

func goFileFilter(info os.FileInfo) bool {
	return !strings.HasSuffix(info.Name(), "_test.go")
}

type dirData struct {
	rootPath  string
	pathParts []string
	children  map[string]*dirData
	pkg       *ast.Package
}

func (d *dirData) onlyPackages() (result []*dirData) {
	result = make([]*dirData, 0)
	if d.pkg != nil {
		// fmt.Println("package", d.pkg.Name)
		result = append(result, d)
	}
	for _, child := range d.children {
		result = append(result, child.onlyPackages()...)
	}
	return result
}

func parseDir(rootPath string, pathParts []string) (data *dirData) {
	dir := filepath.Join(rootPath, filepath.Join(pathParts...))
	dirFile, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer dirFile.Close()

	info, err := dirFile.Stat()
	if err != nil {
		panic(err)
	}
	if !info.IsDir() {
		panic("Path is not a directory: " + dir)
	}

	data = &dirData{
		rootPath:  rootPath,
		pathParts: pathParts,
		children:  make(map[string]*dirData),
	}
	// indent := strings.Repeat("  ", len(pathParts))

	// fmt.Println("Parsing dir:", dir)
	pkgs, err := parser.ParseDir(token.NewFileSet(), dir, goFileFilter, 0)
	if err != nil {
		panic(err)
	}
	for _, pkg := range pkgs {
		// fmt.Println("Found package:", indent+pkg.Name)
		data.pkg = pkg
		// We only expect one package per directory,
		// just take the first one
		break
	}

	files, err := dirFile.Readdir(-1)
	if err != nil {
		panic(err)
	}
	for _, info := range files {
		name := info.Name()
		if info.IsDir() && !(len(pathParts) == 0 && name == "generated") {
			subData := parseDir(rootPath, append(pathParts, name))
			data.children[name] = subData
		}
	}

	return data
}

func goPathSrc() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = "~/go"
	}
	gopath, err := filepath.Abs(filepath.Clean(gopath))
	if err != nil {
		panic(err)
	}
	return filepath.ToSlash(gopath) + "/src/"
}
