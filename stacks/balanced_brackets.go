package stacks

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
