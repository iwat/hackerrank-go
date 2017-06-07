package statistics

func WeightedMean(data []int, weight []int) float32 {
	productSum := 0
	weightSum := 0
	for i, item := range data {
		productSum += item * weight[i]
		weightSum += weight[i]
	}

	return float32(productSum) / float32(weightSum)
}
