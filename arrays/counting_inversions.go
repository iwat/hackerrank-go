package arrays

import (
	"fmt"
)

// Counting inversion implements challenge from HackerRank.com
//
// https://www.hackerrank.com/challenges/ctci-merge-sort
func CountingInversions(data []int) int {
	return mergeSort(data)
}

func mergeSort(data []int) int {
	if len(data) <= 1 {
		return 0
	}
	swaps := mergeSort(data[:len(data)/2])
	swaps += mergeSort(data[len(data)/2:])
	swaps += merge(data)
	return swaps
}

func merge(data []int) int {
	swaps := 0
	for {
		swapped := false
		for i := 1; i < len(data); i++ {
			if data[i] < data[i-1] {
				tmp := data[i]
				data[i] = data[i-1]
				data[i-1] = tmp
				swaps++
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return swaps
}
