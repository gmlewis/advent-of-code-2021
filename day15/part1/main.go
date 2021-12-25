// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"container/heap"
	"flag"
	"fmt"
	"log"
	"math"

	. "github.com/gmlewis/advent-of-code-2021/v1/enum"
	"github.com/gmlewis/advent-of-code-2021/v1/must"
	"github.com/gmlewis/advent-of-code-2021/v1/strfn"
)

var logf = log.Printf
var printf = fmt.Printf

func main() {
	flag.Parse()

	Each(flag.Args(), process)
}

func process(filename string) {
	logf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)

	b := ReduceWithIndex(lines, gridT{}, func(y int, line string, acc gridT) gridT {
		strfn.RunesWithIndex(line, func(x int, r rune) { acc[keyT{x, y}] = int(r - '0') })
		return acc
	})

	xmax := len(lines[0]) - 1
	ymax := len(lines) - 1
	risk := dijkstra(b, keyT{0, 0}, keyT{xmax, ymax})

	printf("Solution: %v\n", risk)
}

type priorityQueue struct {
	dist  gridT
	index gridT
	items []keyT
}

func (pq *priorityQueue) Len() int { return len(pq.items) }

func (pq *priorityQueue) Less(a, b int) bool {
	va, okA := pq.dist[pq.items[a]]
	vb, okB := pq.dist[pq.items[b]]
	switch {
	case okA && okB:
		return va < vb
	case okA:
		return true
	default:
		return false
	}
}

func (pq *priorityQueue) Swap(a, b int) {
	pq.items[a], pq.items[b] = pq.items[b], pq.items[a]
	pq.index[pq.items[a]] = a
	pq.index[pq.items[b]] = b
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(keyT)
	pq.index[item] = n
	pq.items = append(pq.items, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	delete(pq.index, item)
	pq.items = old[0 : n-1]
	return item
}

func dijkstra(b gridT, source, target keyT) int {
	inQ := map[keyT]bool{}
	dist := gridT{}
	q := &priorityQueue{dist: dist, index: gridT{}}
	prev := map[keyT]keyT{}

	for k := range b {
		dist[k] = math.MaxInt
		heap.Push(q, k)
		inQ[k] = true
	}
	dist[source] = 0
	heap.Fix(q, q.index[source])

	f := func(u, v keyT) {
		alt := dist[u] + b[v]
		if alt < dist[v] {
			dist[v] = alt
			prev[v] = u
			heap.Fix(q, q.index[v])
		}
	}

	for q.Len() > 0 {
		u := heap.Pop(q).(keyT)
		delete(inQ, u)

		if u == target {
			break
		}

		if v, ok := moveR(u, target); ok && inQ[v] {
			f(u, v)
		}
		if v, ok := moveD(u, target); ok && inQ[v] {
			f(u, v)
		}
		if v, ok := moveL(u, target); ok && inQ[v] {
			f(u, v)
		}
		if v, ok := moveU(u, target); ok && inQ[v] {
			f(u, v)
		}
	}

	return dist[target]
}

type gridT map[keyT]int

type keyT [2]int

func moveR(key, goal keyT) (keyT, bool) {
	return keyT{key[0] + 1, key[1]}, key[0]+1 <= goal[0]
}

func moveL(key, goal keyT) (keyT, bool) {
	return keyT{key[0] - 1, key[1]}, key[0]-1 >= 0
}

func moveD(key, goal keyT) (keyT, bool) {
	return keyT{key[0], key[1] + 1}, key[1]+1 <= goal[1]
}

func moveU(key, goal keyT) (keyT, bool) {
	return keyT{key[0], key[1] - 1}, key[1]-1 >= 0
}
