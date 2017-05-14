package main

import (
	"sort"
)

func Quartiles(data []int) (int, int, int) {
	sort.Sort(IntSortable(data))

	q2 := median(data)

	q1UpperBound := len(data) / 2
	for data[q1UpperBound] >= q2 {
		q1UpperBound--
	}
	q1UpperBound++

	q3LowerBound := len(data) / 2
	for data[q3LowerBound] <= q2 {
		q3LowerBound++
	}

	q1 := median(data[:q1UpperBound])
	q3 := median(data[q3LowerBound:])

	return q1, q2, q3
}

func median(data []int) int {
	q2 := data[len(data)/2]
	if len(data)%2 == 0 {
		q2 = (q2 + data[len(data)/2-1]) / 2
	}
	return q2
}
