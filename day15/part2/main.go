// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"container/heap"
	"flag"
	"fmt"
	"log"
	"math"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/gmlewis/advent-of-code-2021/strfn"
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

	xsize := len(lines[0])
	ysize := len(lines)
	logf("xsize=%v, ysize=%v", xsize, ysize)

	risk := dijkstra(b, keyT{0, 0}, keyT{xsize, ysize}, keyT{5*xsize - 1, 5*ysize - 1})

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

func dijkstra(b gridT, source, stepSize, target keyT) int {
	inQ := map[keyT]bool{}
	dist := gridT{}
	q := &priorityQueue{dist: dist, index: gridT{}}
	prev := map[keyT]keyT{}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			for oldK := range b {
				k := keyT{oldK[0] + x*stepSize[0], oldK[1] + y*stepSize[1]}
				dist[k] = math.MaxInt
				heap.Push(q, k)
				inQ[k] = true
			}
		}
	}
	dist[source] = 0
	heap.Fix(q, q.index[source])

	valueOf := func(v keyT) int {
		x := v[0] % stepSize[0]
		y := v[1] % stepSize[1]
		d := v[0]/stepSize[0] + v[1]/stepSize[1]
		nv := b[keyT{x, y}] + d
		if nv > 9 {
			return nv%10 + 1
		}
		return nv
	}

	f := func(u, v keyT) {
		alt := dist[u] + valueOf(v)
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
