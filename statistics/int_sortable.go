package main

type IntSortable []int

func (a IntSortable) Len() int           { return len(a) }
func (a IntSortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntSortable) Less(i, j int) bool { return a[i] < a[j] }
