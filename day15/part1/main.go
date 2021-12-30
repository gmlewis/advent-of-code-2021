// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	algo "github.com/gmlewis/advent-of-code-2021/algorithm"
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

	xmax := len(lines[0]) - 1
	ymax := len(lines) - 1
	risk := dijkstra(b, keyT{0, 0}, keyT{xmax, ymax})

	printf("Solution: %v\n", risk)
}

func dijkstra(b gridT, source, target keyT) int {
	inQ := map[keyT]bool{}
	dist := gridT{}
	less := func(a, b keyT) bool {
		va, okA := dist[a]
		vb, okB := dist[b]
		switch {
		case okA && okB:
			return va < vb
		case okA:
			return true
		default:
			return false
		}
	}
	q := algo.NewPriorityQueue(less)
	prev := map[keyT]keyT{}

	for k := range b {
		dist[k] = math.MaxInt
		if k == source {
			dist[k] = 0
		}
		q.Push(k)
		inQ[k] = true
	}

	f := func(u, v keyT) {
		alt := dist[u] + b[v]
		if alt < dist[v] {
			dist[v] = alt
			prev[v] = u
			q.Fix(v)
		}
	}

	for q.Len() > 0 {
		u := q.Pop()
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
