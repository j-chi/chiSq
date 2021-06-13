package chisq

import "math"

func Test(sample []float64, prob []float64) (xSquared float64) {
	err := ChiSqValidate(sample, prob)
	if err != nil {
		panic("data err")
	}

	sampleSize := Sum(sample)
	for i := 0; i < len(sample); i++ {
		o := sampleSize * prob[i]
		xSquared += math.Pow(float64(sample[i]) - o, 2)   / o
	}
	return
}
