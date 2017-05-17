package xu

// Average calculates the average of a slice of float64 values.
func Average(g []float64) float64 {
	var s float64
	for _, x := range g {
		s += x
	}
	return s / float64(len(g))
}
