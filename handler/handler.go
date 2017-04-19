package handler

type HandlerFunc func(ctx *Context)

func NewHandler(method string, handlerFunc HandlerFunc) *Handler {
	return &Handler{method: method, handlerFunc: handlerFunc}
}

type Handler struct {
	method      string
	parent      PathPart
	handlerFunc HandlerFunc
}

func (h *Handler) Name() string {
	return h.method
}

func (h *Handler) String() string {
	return h.method
}

func (h *Handler) Parent() PathPart {
	return h.parent
}

func (h *Handler) setParent(parent PathPart) {
	h.parent = parent
}

func (*Handler) Children() []PathPart {
	return nil
}

func (*Handler) IsArg() bool {
	return true
}

func (h *Handler) Match(part string) bool {
	return part == h.method
}

func (h *Handler) HandlerFunc() HandlerFunc {
	return h.handlerFunc
}
