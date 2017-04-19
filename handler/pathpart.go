package handler

type PathPart interface {
	Name() string
	String() string
	Parent() PathPart
	setParent(parent PathPart)
	Children() []PathPart
	IsArg() bool
	Match(part string) bool
	HandlerFunc() HandlerFunc
}
