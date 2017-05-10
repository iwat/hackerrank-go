package algorithms

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
