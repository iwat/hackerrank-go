package algorithms

func IsUnique(input string) bool {
	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[i] == input[j] {
				println(input[:i])
				return false
			}
		}
	}
	return true
}
