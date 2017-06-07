package arrays

func BubbleSort(input []int) int {
	totalSwaps := 0
	for {
		swaps := 0
		for i := 1; i < len(input); i++ {
			if input[i] < input[i-1] {
				tmp := input[i]
				input[i] = input[i-1]
				input[i-1] = tmp
				swaps++
			}
		}
		totalSwaps += swaps
		if swaps == 0 {
			break
		}
	}
	return totalSwaps
}
