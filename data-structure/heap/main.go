package main

import (
	"fmt"
	"log"
)

type BinaryHeap interface {
	ExtractMax() (int, bool)
	GetItems() []int
	Insert(int)
	IsEmpty() bool
	Peek() int
	Size() int
}

type MaxHeap struct {
	items []int
}

func NewMaxHeap(nums []int) BinaryHeap {
	heap := &MaxHeap{
		items: nums,
	}
	heap.heapify()
	return heap
}

func isValidMaxHeap(heap BinaryHeap) (bool, int) {
	if heap.IsEmpty() {
		return true, -1
	}
	l := heap.Size()
	items := heap.GetItems()
	for i := range l {
		left := 2*i + 1
		right := 2*i + 2
		if left < l && items[i] < items[left] {
			return false, i
		}
		if right < l && items[i] < items[right] {
			return false, i
		}
	}
	return true, -1
}

func (m *MaxHeap) heapify() {
	l := m.Size()
	for i := (l - 1) / 2; i >= 0; i-- {
		m.siftDown(i, l)
	}
}

func (m *MaxHeap) siftDown(index, size int) {
	i := index
	for i < size {
		left := 2*i + 1
		right := 2*i + 2
		child := -1
		if left < size {
			child = left
			if right < size && m.items[right] > m.items[left] {
				child = right
			}
		}
		if child == -1 || m.items[i] > m.items[child] {
			break
		}
		m.items[child], m.items[i] = m.items[i], m.items[child]
		i = child
	}
}

// Shift and restore.
//func (m *MaxHeap) siftUp(index int) {
//	newItem := m.items[index]
//	parent := (index - 1) / 2
//	cur := index
//	for cur > 0 && newItem > m.items[parent] {
//		m.items[cur] = m.items[parent]
//		cur = parent
//		parent = (parent - 1) / 2
//	}
//	m.items[parent] = newItem
//}

func (m *MaxHeap) siftUp(index int) {
	cur := index
	for cur > 0 {
		parent := (cur - 1) / 2
		if m.items[cur] < m.items[parent] {
			break
		}
		m.items[cur], m.items[parent] = m.items[parent], m.items[cur]
		cur = parent
	}
}

func (m *MaxHeap) ExtractMax() (int, bool) {
	if m.IsEmpty() {
		return 0, false
	}
	v := m.items[0]
	m.items[0] = m.items[m.Size()-1]
	m.items = m.items[:m.Size()-1]
	if !m.IsEmpty() {
		m.siftDown(0, m.Size())
	}
	return v, true
}

func (m *MaxHeap) GetItems() []int {
	v := make([]int, m.Size())
	copy(v, m.items)
	return v
}

func (m *MaxHeap) Insert(v int) {
	m.items = append(m.items, v)
	m.siftUp(m.Size() - 1)
}

func (m *MaxHeap) IsEmpty() bool {
	return m.Size() == 0
}

func (m *MaxHeap) Peek() int {
	if !m.IsEmpty() {
		return m.items[0]
	}
	return -1
}

func (m *MaxHeap) Size() int {
	return len(m.items)
}

func heapSort(nums []int) {
	heap := &MaxHeap{
		items: nums,
	}
	heap.heapify()
	fmt.Printf("nums=%#v\n", nums)
	heapSize := heap.Size() - 1
	for heapSize > 0 {
		heap.items[heapSize], heap.items[0] = heap.items[0], heap.items[heapSize]
		heap.siftDown(0, heapSize)
		heapSize--
	}
	fmt.Printf("nums=%#v\n", nums)
}

func main() {
	tests := []struct {
		input  []int
		insert int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			insert: 10,
		},
	}
	for _, test := range tests {
		heap := NewMaxHeap(test.input)
		heap.Insert(test.insert)
		valid, badIndex := isValidMaxHeap(heap)
		if !valid {
			log.Printf("Heap property violated at index %d: %#v", badIndex, heap.GetItems())
			continue
		}
		log.Printf("Heap is valid %#v", heap.GetItems())
		heap.ExtractMax()
		valid, badIndex = isValidMaxHeap(heap)
		if !valid {
			log.Printf("Heap property violated at index %d: %#v", badIndex, heap.GetItems())
			continue
		}
		log.Printf("Heap is valid %#v", heap.GetItems())
	}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	heapSort(nums)
}
