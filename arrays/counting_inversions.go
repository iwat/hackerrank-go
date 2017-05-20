package arrays

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
	result := make([]int, 0, len(data))
	left := data[:len(data)/2]
	right := data[len(data)/2:]
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			result = append(result, right[0])
			right = right[1:]
		} else if len(right) == 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
			swaps += len(left)
		}
	}
	copy(data, result)
	return swaps
}
