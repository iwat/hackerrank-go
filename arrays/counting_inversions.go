package arrays

// Counting inversion implements challenge from HackerRank.com
//
// https://www.hackerrank.com/challenges/ctci-merge-sort
func CountingInversions(data []int) int {
	aux := make([]int, len(data))
	copy(aux, data)
	return mergeSort(data, aux)
}

func mergeSort(dest, src []int) int {
	if len(dest) <= 1 {
		return 0
	}
	swaps := mergeSort(src[:len(dest)/2], dest[:len(src)/2])
	swaps += mergeSort(src[len(dest)/2:], dest[len(src)/2:])
	swaps += merge(dest, src)
	return swaps
}

func merge(dest, src []int) int {
	swaps := 0
	i := 0
	left := src[:len(src)/2]
	right := src[len(src)/2:]
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			dest[i] = right[0]
			right = right[1:]
		} else if len(right) == 0 {
			dest[i] = left[0]
			left = left[1:]
		} else if left[0] <= right[0] {
			dest[i] = left[0]
			left = left[1:]
		} else {
			dest[i] = right[0]
			right = right[1:]
			swaps += len(left)
		}
		i++
	}
	return swaps
}
