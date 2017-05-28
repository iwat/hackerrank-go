package search

import "fmt"

func SherlockAndArray() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	if n == 1 {
		fmt.Println("YES")
		return
	} else if n == 2 {
		fmt.Println("NO")
		return
	}

	leftNdx := 1
	leftSum := arr[0]

	rightNdx := n - 2
	rightSum := arr[n-1]

	for leftNdx != rightNdx {
		if leftSum <= rightSum {
			leftSum += arr[leftNdx]
			leftNdx++
		} else {
			rightSum += arr[rightNdx]
			rightNdx--
		}
	}

	if leftSum == rightSum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func SherlockAndArrayMain() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		do()
	}
}
