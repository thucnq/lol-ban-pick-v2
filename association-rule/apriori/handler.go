package apriori

type handler struct {
	minimumSupport float64
	dataset        [][]string
}

type Option func(*handler)
