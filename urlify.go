package algorithms

// URLify : Write a method to replace all spaces in a string with '%20  You may
// assume that the string has suf cient space at the end to hold the additional
// characters,and that you are given the "true" length of the string. (Note: If
// implementing in Java,please use a character array so that you can perform
// this operation in place.)
//
// EXAMPLE
// Input:  "Mr John Smith    ", 13
// Output: "Mr%20John%20Smith"
func URLify(buffer []byte, originalLen int) {
	spaces := 0
	for i := 0; i < originalLen; i++ {
		if buffer[i] == ' ' {
			spaces++
		}
	}
	if originalLen+spaces*2 > len(buffer) {
		return
	}
	i := originalLen - 1
	j := originalLen + spaces*2 - 1
	for i >= 0 && j >= 2 {
		if buffer[i] == ' ' {
			buffer[j-2] = '%'
			buffer[j-1] = '2'
			buffer[j] = '0'
			j -= 3
		} else {
			buffer[j] = buffer[i]
			j--
		}
		i--
	}
}
