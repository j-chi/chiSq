package chisq


func Sum(prob []float64) (sum float64) {
	for _, v := range prob {
		sum += v
	}
	return
}