package main

func Fibonacci(a int) int {
	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	}
	return Fibonacci(a-1) + Fibonacci(a-2)
}
