package search

import (
	"fmt"
	"io"
)

func MissingNumber(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscanf(in, "%d", &n)
	a := make([]int, 10001)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(in, "%d", &x)
		if x == 0 {
			i--
		}
		a[x]++
	}

	var m int
	fmt.Fscanf(in, "%d", &m)
	if m == 0 {
		fmt.Fscanf(in, "%d", &m)
	}

	for i := 0; i < m; i++ {
		var x int
		fmt.Fscanf(in, "%d", &x)
		if x == 0 {
			i--
		}
		a[x]--
	}

	for i, x := range a {
		if i > 0 && x < 0 {
			fmt.Fprint(out, i)
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out)
}
