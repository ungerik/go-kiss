package handler

import (
	"strconv"
	"strings"
)

func NewInt(name string, children ...PathPart) *Int {
	name = strings.TrimSuffix(strings.TrimPrefix(name, "_"), "_int_")
	return &Int{name: name, children: children}
}

type Int struct {
	name     string
	parent   PathPart
	children []PathPart
}

func (i *Int) Name() string {
	return i.name
}

func (i *Int) String() string {
	return "_" + i.name + "_int_"
}

func (i *Int) Parent() PathPart {
	return i.parent
}

func (i *Int) setParent(parent PathPart) {
	i.parent = parent
}

func (i *Int) Children() []PathPart {
	return i.children
}

func (*Int) IsArg() bool {
	return true
}

func (*Int) Match(part string) bool {
	_, err := strconv.ParseInt(part, 10, 64)
	return err == nil
}

func (*Int) HandlerFunc() HandlerFunc {
	return nil
}
