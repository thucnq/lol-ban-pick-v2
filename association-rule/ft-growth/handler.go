package ftgrowth

type handler struct {
	minimumSupport float64
	minConfidence  float64
	minLift        float64
	dataset        [][]string
}

type Option func(*handler)
