// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"sort"

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
	// risk := lowestRisk(keyT{0, 0}, keyT{xmax, ymax}, b, visitedT{}, 0, math.MaxInt)

	risk := dijkstra(b, keyT{0, 0}, keyT{xsize, ysize}, keyT{5*xsize - 1, 5*ysize - 1})

	printf("Solution: %v\n", risk)
}

func dijkstra(b gridT, source, stepSize, target keyT) int {
	var q []keyT
	inQ := map[keyT]bool{}
	dist := gridT{}
	prev := map[keyT]keyT{}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			for oldK := range b {
				k := keyT{oldK[0] + x*stepSize[0], oldK[1] + y*stepSize[1]}
				dist[k] = math.MaxInt
				q = append(q, k)
				inQ[k] = true
			}
		}
	}
	dist[source] = 0

	valueOf := func(v keyT) int {
		x := v[0] % stepSize[0]
		y := v[1] % stepSize[1]
		d := v[0]/stepSize[0] + v[1]/stepSize[1]
		nv := b[keyT{x, y}] + d
		if nv > 9 {
			// logf("B:valueOf(%v) = %v", v, nv%10+1)
			return nv%10 + 1
		}
		// logf("A:valueOf(%v) = %v", v, nv)
		return nv
	}

	f := func(u, v keyT) {
		alt := dist[u] + valueOf(v)
		if alt < dist[v] {
			dist[v] = alt
			prev[v] = u
		}
		// logf("f(%v,%v): alt=%v, dist[%v]=%v, prev[%v]=%v", u, v, alt, v, alt, v, u)
	}

	for len(q) > 0 {
		sort.Slice(q, func(a, b int) bool { return dist[q[a]] < dist[q[b]] })
		u := q[0]
		q = q[1:]
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

	// logf("dist=%+v, prev=%+v", dist, prev)
	return dist[target]
}

// func lowestRisk(pos, goal keyT, b gridT, visited visitedT, risk, bestRisk int) int {
// 	if pos == goal {
// 		return risk
// 	}
// 	// visited[pos] = true
//
// 	if p, ok := moveR(pos, goal); ok && !visited[p] && risk+b[p] < bestRisk {
// 		r := lowestRisk(p, goal, b, visited, risk+b[p], bestRisk)
// 		if r < bestRisk {
// 			bestRisk = r
// 		}
// 	}
//
// 	if p, ok := moveD(pos, goal); ok && !visited[p] && risk+b[p] < bestRisk {
// 		r := lowestRisk(p, goal, b, visited, risk+b[p], bestRisk)
// 		if r < bestRisk {
// 			bestRisk = r
// 		}
// 	}
//
// 	return bestRisk
// }
//
// func copyV(v visitedT) visitedT {
// 	ret := visitedT{}
// 	for k, v := range v {
// 		ret[k] = v
// 	}
// 	return ret
// }

type gridT map[keyT]int

type keyT [2]int

type visitedT map[keyT]bool

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
