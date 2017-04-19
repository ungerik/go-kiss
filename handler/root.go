package handler

type Root struct {
	children []PathPart
}

func NewRoot(children ...PathPart) *Root {
	return &Root{children}
}

func (*Root) Name() string {
	return "/"
}

func (*Root) String() string {
	return "/"
}

func (*Root) Parent() PathPart {
	return nil
}

func (*Root) setParent(parent PathPart) {
	if parent != nil {
		panic("only nil parent allowed for pathpart.Root")
	}
}

func (c *Root) Children() []PathPart {
	return c.children
}

func (*Root) IsArg() bool {
	return false
}

func (*Root) Match(part string) bool {
	return part == "/"
}

func (*Root) HandlerFunc() HandlerFunc {
	return nil
}
