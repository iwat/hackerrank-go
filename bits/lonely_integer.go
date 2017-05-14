package main

func LonelyInteger(arr ...uint) int {
	low := uint64(0)
	high := uint64(0)

	for _, i := range arr {
		if i < 64 {
			low ^= 1 << i
		} else {
			high ^= 1 << (i - 64)
		}
	}
	for i := 0; i < 64; i++ {
		if low&1 == 1 {
			return i
		}
		low = low >> 1
	}
	for i := 0; i < 64; i++ {
		if high&1 == 1 {
			return i + 64
		}
		high = high >> 1
	}
	return -1
}
