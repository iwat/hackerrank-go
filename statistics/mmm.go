package main

import (
	"sort"
)

type IntSortable []int

func (a IntSortable) Len() int           { return len(a) }
func (a IntSortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntSortable) Less(i, j int) bool { return a[i] < a[j] }

func MeanMedianMode(data []int) (float32, float32, int) {
	sort.Sort(IntSortable(data))

	sum := 0
	hitMap := make(map[int]int, len(data)/2)
	for _, item := range data {
		sum += item
		if _, ok := hitMap[item]; ok {
			hitMap[item]++
		} else {
			hitMap[item] = 1
		}
	}

	mean := float32(sum) / float32(len(data))
	median := float32(data[len(data)/2])
	if len(data)%2 == 0 {
		median = (median + float32(data[len(data)/2-1])) / float32(2)
	}

	mode := data[0]
	modeHits := 0
	for _, item := range data {
		if hitMap[item] > modeHits {
			modeHits = hitMap[item]
			mode = item
		}
	}

	return mean, median, mode
}
