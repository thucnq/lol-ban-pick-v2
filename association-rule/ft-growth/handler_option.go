package ftgrowth

func WithDataSet(dataSet [][]string) Option {
	return func(s *handler) {
		s.dataset = dataSet
	}
}

func WithMinimumSupport(minimumSupport float64) Option {
	return func(s *handler) {
		s.minimumSupport = minimumSupport
	}
}

func WithMinConfidence(minConfidence float64) Option {
	return func(s *handler) {
		s.minConfidence = minConfidence
	}
}

func WithMinLift(minLift float64) Option {
	return func(s *handler) {
		s.minLift = minLift
	}
}
