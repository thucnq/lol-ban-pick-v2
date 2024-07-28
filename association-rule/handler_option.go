package associationrule

func WithAlgorithm(name string, algorithm IAlgorithm) Option {
	return func(s *handler) {
		s.algorithms[name] = algorithm
	}
}
