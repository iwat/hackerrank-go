package stacks

func BalancedBracketsStringBacked(input string) bool {
	pair := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := make([]rune, 0)
	for _, c := range input {
		if open, ok := pair[c]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

func BalancedBrackets(input string) bool {
	pair := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := NewStack()
	for _, c := range input {
		if open, ok := pair[c]; ok {
			if stack.Peek() != open {
				return false
			} else {
				stack.Pop()
			}
		} else {
			stack.Push(c)
		}
	}
	return stack.Peek() == nil
}
