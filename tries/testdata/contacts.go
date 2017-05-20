package main

import (
	"fmt"

	"github.com/iwat/algorithms/tries"
)

func main() {
	trie := tries.NewTrie("")

	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var cmd string
		var txt string
		fmt.Scan(&cmd)
		fmt.Scan(&txt)
		switch cmd {
		case "add":
			trie.Add(txt)
		case "find":
			fmt.Println(trie.Find(txt))
		}
	}
}
