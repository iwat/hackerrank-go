package recursions

func DavisStaircase(cases int) int {
	switch cases {
	case 0, 1:
		return cases
	case 2:
		return 2
	case 3:
		return 4
	}

	ways := DavisStaircase(cases - 1)
	ways += DavisStaircase(cases - 2)
	ways += DavisStaircase(cases - 3)
	return ways
}

func DavisStaircaseIterative(cases int) int {
	switch cases {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	}

	ways := make([]int, cases)
	ways[0] = 1
	ways[1] = 2
	ways[2] = 4
	for i := 3; i < cases; i++ {
		ways[i] = ways[i-1] + ways[i-2] + ways[i-3]
	}
	return ways[len(ways)-1]
}

func DavisStaircaseIterativeSlow(cases int) int {
	ways := 0
	queue := []int{cases}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		switch n {
		case 0:
			break
		case 1:
			ways += 1
			break
		case 2:
			ways += 2
			break
		case 3:
			ways += 4
			break
		default:
			queue = append(queue, n-1)
			queue = append(queue, n-2)
			queue = append(queue, n-3)
		}
	}
	return ways
}
