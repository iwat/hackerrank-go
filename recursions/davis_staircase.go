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
