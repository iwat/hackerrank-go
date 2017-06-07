package main

import (
	"fmt"

	"github.com/iwat/algorithms/heaps"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	median := heaps.NewRunningMedian()
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		fmt.Printf("%.1f\n", median.Add(a))
	}
}
