package handler

type Const struct {
	name     string
	parent   PathPart
	children []PathPart
}

func NewConst(name string, children ...PathPart) *Const {
	return &Const{name: name, children: children}
}

func (c *Const) Name() string {
	return c.name
}

func (c *Const) String() string {
	return c.name
}

func (c *Const) Parent() PathPart {
	return c.parent
}

func (c *Const) setParent(parent PathPart) {
	c.parent = parent
}

func (c *Const) Children() []PathPart {
	return c.children
}

func (*Const) IsArg() bool {
	return false
}

func (c *Const) Match(part string) bool {
	return part == c.name
}

func (*Const) HandlerFunc() HandlerFunc {
	return nil
}
