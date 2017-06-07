package arrays

func InsertionSort(input []int) {
	for i := 1; i < len(input); i++ {
		j := i
		for j > 0 && input[j] < input[j-1] {
			tmp := input[j]
			input[j] = input[j-1]
			input[j-1] = tmp
			j--
		}
	}
}
