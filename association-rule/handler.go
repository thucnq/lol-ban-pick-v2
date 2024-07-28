package associationrule

type IHandler interface {
	Process()
}

type IAlgorithm interface {
	Process()
}

type handler struct {
	algorithms map[string]IAlgorithm
}

type Option func(*handler)

func New(opts ...Option) handler {
	h := handler{
		algorithms: map[string]IAlgorithm{},
	}

	for i := range opts {
		opts[i](&h)
	}

	return h
}

func (h handler) ProcessBy(name string) {
	h.algorithms[name].Process()
	return
}
