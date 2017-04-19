package handler

func NewString(name string, children ...PathPart) *String {
	name = name[1 : len(name)-1]
	return &String{name: name, children: children}
}

type String struct {
	name     string
	parent   PathPart
	children []PathPart
}

func (s *String) Name() string {
	return s.name
}

func (s *String) String() string {
	return "_" + s.name + "_"
}

func (s *String) Parent() PathPart {
	return s.parent
}

func (s *String) setParent(parent PathPart) {
	s.parent = parent
}

func (s *String) Children() []PathPart {
	return s.children
}

func (*String) IsArg() bool {
	return true
}

func (*String) Match(part string) bool {
	return true
}

func (*String) HandlerFunc() HandlerFunc {
	return nil
}
