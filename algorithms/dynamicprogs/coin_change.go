package dynamicprogs

import (
	"bytes"
	"fmt"
	"sort"
)

const debug = false

type ByValue []int

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i] > a[j] }

func CoinChange(value int, coins ...int) int {
	ways := make(map[string]int, len(coins))
	sort.Sort(ByValue(coins))

	return calculateWays(value, coins, ways, "")
}

func calculateWays(value int, coins []int, ways map[string]int, stack string) int {
	if value < 0 {
		return 0
	}
	if value == 0 {
		if debug {
			fmt.Println(stack, "Good")
		}
		return 1
	}
	key := makeKey(value, coins)
	if w, ok := ways[key]; ok {
		return w
	}
	for i, coin := range coins {
		ways[key] += calculateWays(value-coin, coins[i:], ways, fmt.Sprint(stack, coin))
	}
	if debug {
		fmt.Println("There are", ways[key], "way(s) to change", value, "into", coins)
	}
	return ways[key]
}

func makeKey(value int, coins []int) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprint(value))
	buffer.WriteString("/")
	for i, coin := range coins {
		buffer.WriteString(fmt.Sprint(coin))
		if i < len(coins)-1 {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}
