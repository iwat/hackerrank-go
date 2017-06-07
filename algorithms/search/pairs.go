package search

import (
	"fmt"
	"os"
)

func PairsMain() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	count := 0
	wantedMore := make(map[int]bool)
	wantedLess := make(map[int]bool)

	for i := 0; i < n; i++ {
		var d int
		fmt.Scanf("%d", &d)
		inMore := false
		inLess := false
		if _, ok := wantedMore[d]; ok {
			count++
			fmt.Fprintln(os.Stderr, "found", d, "in more list")
			delete(wantedMore, d)
			inMore = true
		}
		if _, ok := wantedLess[d]; ok {
			count++
			fmt.Fprintln(os.Stderr, "found", d, "in less list")
			delete(wantedLess, d)
			inLess = true
		}
		if !inMore && d-k >= 0 {
			fmt.Fprintln(os.Stderr, "wanted", d-k)
			wantedLess[d-k] = true
		}
		if !inLess {
			fmt.Fprintln(os.Stderr, "wanted", d+k)
			wantedMore[d+k] = true
		}

	}
	fmt.Println(count)
}
