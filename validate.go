package chisq

func ChiSqValidate(sample []float64, prob []float64) (err error) {
	if len(sample) != len(prob) {
		panic("sample and prob length err")
	}

	if Sum(prob) != 1 {
		panic("sum of prob not equal 1")
	}

	return
}
