package arrays

// Counting inversion implements challenge from HackerRank.com
//
// https://www.hackerrank.com/challenges/ctci-merge-sort
func CountingInversions(data []int) int {
	aux := make([]int, len(data))
	return mergeSort(data, aux)
}

func mergeSort(data, aux []int) int {
	if len(data) <= 1 {
		return 0
	}
	swaps := mergeSort(data[:len(data)/2], aux)
	swaps += mergeSort(data[len(data)/2:], aux)
	swaps += merge(data, aux)
	return swaps
}

func merge(data, aux []int) int {
	swaps := 0
	i := 0
	left := data[:len(data)/2]
	right := data[len(data)/2:]
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			aux[i] = right[0]
			right = right[1:]
		} else if len(right) == 0 {
			aux[i] = left[0]
			left = left[1:]
		} else if left[0] <= right[0] {
			aux[i] = left[0]
			left = left[1:]
		} else {
			aux[i] = right[0]
			right = right[1:]
			swaps += len(left)
		}
		i++
	}
	copy(data, aux)
	return swaps
}
