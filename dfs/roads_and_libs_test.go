package dfs

import (
	"testing"
)

func TestRoadsAndLibs1(t *testing.T) {
	ral := NewRoadsAndLibs(3, 3, 2, 1)
	ral.Link(1, 2)
	ral.Link(3, 1)
	ral.Link(2, 3)
	if 4 != ral.Cost() {
		t.Error("cost != 4")
	}
}

func TestRoadsAndLibs2(t *testing.T) {
	ral := NewRoadsAndLibs(6, 6, 2, 5)
	ral.Link(1, 3)
	ral.Link(3, 4)
	ral.Link(2, 4)
	ral.Link(1, 2)
	ral.Link(2, 3)
	ral.Link(5, 6)
	if 12 != ral.Cost() {
		t.Error("cost != 12")
	}
}

func TestRoadsAndLibs3(t *testing.T) {
	ral := NewRoadsAndLibs(9, 2, 91, 84)
	ral.Link(8, 3)
	ral.Link(2, 9)
	if 805 != ral.Cost() {
		t.Error("cost != 805")
	}
}

func TestRoadsAndLibs4(t *testing.T) {
	ral := NewRoadsAndLibs(5, 9, 92, 23)
	ral.Link(2, 1)
	ral.Link(5, 3)
	ral.Link(5, 1)
	ral.Link(3, 4)
	ral.Link(3, 1)
	ral.Link(5, 4)
	ral.Link(4, 1)
	ral.Link(5, 2)
	ral.Link(4, 2)
	if 184 != ral.Cost() {
		t.Error("cost != 184")
	}
}

func TestRoadsAndLibs5(t *testing.T) {
	ral := NewRoadsAndLibs(8, 3, 10, 55)
	ral.Link(6, 4)
	ral.Link(3, 2)
	ral.Link(7, 1)
	if 80 != ral.Cost() {
		t.Error("cost != 80")
	}
}

func TestRoadsAndLibs6(t *testing.T) {
	ral := NewRoadsAndLibs(1, 0, 5, 3)
	if 5 != ral.Cost() {
		t.Error("cost != 5")
	}
}

func TestRoadsAndLibs7(t *testing.T) {
	ral := NewRoadsAndLibs(2, 0, 102, 1)
	if 204 != ral.Cost() {
		t.Error("cost != 204")
	}
}
