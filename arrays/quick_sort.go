package arrays

import (
	"fmt"
)

func QuickSort(arr []int, from, to int) {
	if to-from <= 1 {
		return
	}
	p := lumoto(arr, from, to)
	QuickSort(arr, from, p-1)
	QuickSort(arr, p, to)
}

func printArr(arr []int) {
	for i, d := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(d)
	}
	fmt.Println()
}

func lumoto(A []int, lo, hi int) int {
	pivot := A[hi-1]
	i := lo
	for j := lo; j < hi; j++ {
		if A[j] <= pivot {
			if i != j {
				A[i], A[j] = A[j], A[i]
			}
			i += 1
		}
	}
	printArr(A)
	return i
}
