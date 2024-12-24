package util

import "container/heap"

// Item represents an item in the heap.
type Item[T any] struct {
	Value    T
	Priority int
}

type minHeap[T any] []Item[T]

func (pq minHeap[T]) Len() int { return len(pq) }

func (pq minHeap[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq minHeap[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *minHeap[T]) Push(x any) {
	item := x.(Item[T])
	*pq = append(*pq, item)
}

func (pq *minHeap[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func NewMaxHeap[T any](items ...Item[T]) *MaxHeap[T] {
	h := make(minHeap[T], 0, len(items))
	m := &MaxHeap[T]{heap: &h}
	for _, item := range items {
		m.Push(item.Value, item.Priority)
	}
	return m
}

type MaxHeap[T any] struct {
	heap *minHeap[T]
}

func (m *MaxHeap[T]) Init(items ...Item[T]) {
	h := make(minHeap[T], len(items))
	heap.Init(&h)
	m.heap = &h
}

func (m *MaxHeap[T]) Push(thing T, priority int) {
	heap.Push(m.heap, Item[T]{Priority: -priority, Value: thing})
}

func (m *MaxHeap[T]) Pop() T {
	i := heap.Pop(m.heap).(Item[T])
	return i.Value
}

func (m *MaxHeap[T]) PopItem() Item[T] {
	return heap.Pop(m.heap).(Item[T])
}

func (m *MaxHeap[T]) Len() int { return len(*m.heap) }

func NewMinHeap[T any](items ...Item[T]) *MinHeap[T] {
	h := make(minHeap[T], len(items))
	for idx, item := range items {
		h[idx] = item
	}
	heap.Init(&h)
	return &MinHeap[T]{heap: &h}
}

type MinHeap[T any] struct {
	heap *minHeap[T]
}

func (m *MinHeap[T]) Push(thing T, priority int) {
	heap.Push(m.heap, Item[T]{Priority: priority, Value: thing})
}

func (m *MinHeap[T]) Pop() T {
	i := heap.Pop(m.heap).(Item[T])
	return i.Value
}

func (m *MinHeap[T]) PopItem() Item[T] {
	return heap.Pop(m.heap).(Item[T])
}

func (m *MinHeap[T]) Len() int { return len(*m.heap) }
