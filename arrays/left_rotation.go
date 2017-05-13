package arrays

func RotateLeft(arr []int, d int) {
	// Golang cheat
	// result := append(arr[d:], arr[0:d]...)
	// for i, value := range result {
	//     arr[i] = value
	// }

	gcd := calculateGCD(len(arr), d)
	loops := len(arr) / gcd
	for i := 0; i < d; i++ {
		tmp := arr[i]
		for j := 0; j < loops; j++ {
			ndx := (i + j*gcd) % len(arr)
			if j <= loops-2 {
				arr[ndx] = arr[(i+(j+1)*gcd)%len(arr)]
			} else {
				arr[ndx] = tmp
			}
		}
	}
}

func calculateGCD(n, d int) int {
	if d == 0 {
		return n
	}
	return calculateGCD(d, n%d)
}
