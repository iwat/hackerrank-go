package heaps

import (
	"fmt"
)

const debug = false

type RunningMedian struct {
	minHeap *Heap
	maxHeap *Heap
}

func NewRunningMedian() *RunningMedian {
	return &RunningMedian{NewHeap(1), NewHeap(-1)}
}

func (m *RunningMedian) Add(value int) float64 {
	if m.maxHeap.Peek() != -1 && value < m.maxHeap.Peek() {
		m.maxHeap.Add(value)
	} else if m.minHeap.Peek() != -1 && value > m.minHeap.Peek() {
		m.minHeap.Add(value)
	} else {
		m.minHeap.Add(value)
	}

	for m.minHeap.Size()-m.maxHeap.Size() > 1 {
		m.maxHeap.Add(m.minHeap.Poll())
	}
	for m.maxHeap.Size()-m.minHeap.Size() > 0 {
		m.minHeap.Add(m.maxHeap.Poll())
	}

	if (m.minHeap.Size()+m.maxHeap.Size())%2 == 0 {
		return (float64(m.minHeap.Peek()) + float64(m.maxHeap.Peek())) / 2.0
	}
	return float64(m.minHeap.Peek())
}

type Heap struct {
	data []int
	mode int
}

func NewHeap(mode int) *Heap {
	return &Heap{nil, mode}
}

func (h *Heap) Add(value int) {
	if debug {
		fmt.Println(" +", value, h.data)
	}
	h.data = append(h.data, value)
	if debug {
		fmt.Println(" ?", value, h.data)
	}
	h.heapifyUp()
	if debug {
		fmt.Println("++", value, h.data)
	}
}

func (h *Heap) Peek() int {
	if len(h.data) == 0 {
		return -1
	}
	return h.data[0]
}

func (h *Heap) Poll() int {
	head := h.data[0]
	if debug {
		fmt.Println(" -", head, h.data)
	}
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if debug {
		fmt.Println(" ?", head, h.data)
	}
	h.heapifyDown()
	if debug {
		fmt.Println("--", head, h.data)
	}
	return head
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) heapifyUp() {
	index := len(h.data) - 1
	parent := (index - 1) / 2
	for index > 0 {
		if h.mode*(h.data[parent]-h.data[index]) <= 0 {
			break
		}

		tmp := h.data[parent]
		h.data[parent] = h.data[index]
		h.data[index] = tmp
		if debug {
			fmt.Println("xx", parent, h.data[parent], "<->", index, h.data[index])
		}

		index = parent
		parent = (index - 1) / 2
	}
}

func (h *Heap) heapifyDown() {
	index := 0
	left := index*2 + 1
	right := index*2 + 2
	for left < len(h.data) {
		diffLeft := h.mode * (h.data[left] - h.data[index])
		diffRight := 0
		if right < len(h.data) {
			diffRight = h.mode * (h.data[right] - h.data[index])
		}
		next := -1
		if diffLeft < 0 && diffLeft <= diffRight {
			next = left
		} else if diffRight < 0 && diffRight <= diffLeft {
			next = right
		} else {
			break
		}

		tmp := h.data[next]
		h.data[next] = h.data[index]
		h.data[index] = tmp
		if debug {
			fmt.Println("xx", next, h.data[next], "<->", index, h.data[index])
		}

		index = next

		left = index*2 + 1
		right = index*2 + 2
	}
}
