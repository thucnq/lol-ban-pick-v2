package apriori

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
