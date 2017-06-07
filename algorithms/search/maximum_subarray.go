package search

import (
	"sort"
)

func MaximumSubarraySum(mod int, data ...int) int {
	var values []int

	for i := 1; i <= len(data); i++ {
		values = append(values, subarraySum(mod, i, data)...)
	}

	sort.Ints(values)
	return values[len(values)-1]
}

func subarraySum(mod, size int, data []int) []int {
	sums := make([]int, len(data)-size+1)
	for i := size - 1; i < len(data); i++ {
		sum := 0
		for j := 0; j <= i; j++ {
			sum += data[j]
		}
		sums[i-size+1] = sum % mod
	}
	return sums
}
