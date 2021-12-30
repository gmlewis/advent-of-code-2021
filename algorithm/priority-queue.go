package algorithm

import (
	"container/heap"
)

// PriorityQueue represents a queue that maintains its priority order.
type PriorityQueue[K comparable] struct {
	q *queueT[K]
}

// queue is the internal representation of the queue to be used by container/heap.
type queueT[K comparable] struct {
	less  func(K, K) bool
	index map[K]int
	items []K
}

// NewPriorityQueue returns a new priority queue.
func NewPriorityQueue[K comparable](less func(K, K) bool) *PriorityQueue[K] {
	return &PriorityQueue[K]{q: &queueT[K]{less: less, index: map[K]int{}}}
}

// Len returns the current number of items in the priority queue.
func (pq *PriorityQueue[K]) Len() int { return len(pq.q.items) }

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (pq *PriorityQueue[K]) Push(key K) {
	heap.Push(pq.q, key)
}

// Pop removes and returns the minimum element (according to Less)
// from the heap. The complexity is O(log n) where n = h.Len(). Pop is
// equivalent to Remove(h, 0).
func (pq *PriorityQueue[K]) Pop() K {
	return heap.Pop(pq.q).(K)
}

// Fix re-establishes the heap ordering after the element at key
// has changed its value. Changing the value of the element at key
// and then calling Fix is equivalent to, but less expensive than,
// calling Remove(key) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (pq *PriorityQueue[K]) Fix(key K) {
	heap.Fix(pq.q, pq.q.index[key])
}

// Init establishes the heap invariants required by the other routines
// in this package. Init is idempotent with respect to the heap
// invariants and may be called whenever the heap invariants may have
// been invalidated.
// The complexity is O(n) where n = h.Len().
func (pq *PriorityQueue[K]) Init() {
	heap.Init(pq.q)
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (pq *PriorityQueue[K]) Remove(key K) {
	if n, ok := pq.q.index[key]; ok {
		heap.Remove(pq.q, n)
	}
}

// Len returns the current number of items in the priority queue.
// It is used by the heap interface.
func (pq *queueT[K]) Len() int { return len(pq.items) }

// Less returns true if item a is less than item b.
// It is used by the heap interface.
func (pq *queueT[K]) Less(a, b int) bool { return pq.less(pq.items[a], pq.items[b]) }

// Swap swaps two items in the priority queue and is used by the heap interface.
func (pq *queueT[K]) Swap(a, b int) {
	pq.items[a], pq.items[b] = pq.items[b], pq.items[a]
	pq.index[pq.items[a]] = a
	pq.index[pq.items[b]] = b
}

// Push is used only by the heap interface.
func (pq *queueT[K]) Push(x interface{}) {
	n := len(pq.items)
	item := x.(K)
	pq.index[item] = n
	pq.items = append(pq.items, item)
}

// Pop is used only by the heap interface.
func (pq *queueT[K]) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	delete(pq.index, item)
	pq.items = old[0 : n-1]
	return item
}
