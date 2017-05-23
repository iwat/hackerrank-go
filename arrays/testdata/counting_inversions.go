package main

import (
	"fmt"

	"github.com/iwat/algorithms/arrays"
)

func main() {
	var d int
	fmt.Scanf("%d", &d)
	for i := 0; i < d; i++ {
		var n int
		fmt.Scanf("%d", &n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &arr[j])
		}
		fmt.Println(arrays.CountingInversions(arr))
	}
}
