// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

	. "github.com/gmlewis/advent-of-code-2021/v1/enum"
	"github.com/gmlewis/advent-of-code-2021/v1/maps"
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

	m := ReduceWithIndex(lines, gridT{}, func(y int, line string, acc gridT) gridT {
		strfn.RunesWithIndex(line, func(x int, r rune) { acc[keyT{x, y}] = int(r - '0') })
		return acc
	})

	lowPts := maps.Reduce(m, []keyT{}, findLowPoints(m))
	if len(lowPts) > 4 { // optimization - only return '0' low points.
		lowPts = Filter(lowPts, func(k keyT) bool { return m[k] == 0 })
	}

	allBasins := Map(lowPts, calcBasinSize(m))
	sort.Sort(sort.Reverse(sort.IntSlice(allBasins)))
	answer := Product(allBasins[0:3])

	printf("Solution: %v\n", answer)
}

type gridT map[keyT]int
type keyT [2]int

func calcBasinSize(m gridT) func(k keyT) int {
	return func(k keyT) int {
		visited := map[keyT]bool{}
		return checkNeighbors(m, k, visited)
	}
}

func checkNeighbors(m gridT, k keyT, visited map[keyT]bool) int {
	visited[k] = true
	if m[k] == 9 {
		return 0
	}
	acc := 1
	if p := right(k); !visited[p] && m[p] > m[k] {
		acc += checkNeighbors(m, p, visited)
	}
	if p := left(k); !visited[p] && m[p] > m[k] {
		acc += checkNeighbors(m, p, visited)
	}
	if p := up(k); !visited[p] && m[p] > m[k] {
		acc += checkNeighbors(m, p, visited)
	}
	if p := down(k); !visited[p] && m[p] > m[k] {
		acc += checkNeighbors(m, p, visited)
	}
	return acc
}

func right(k keyT) keyT { return keyT{k[0] + 1, k[1]} }
func left(k keyT) keyT  { return keyT{k[0] - 1, k[1]} }
func up(k keyT) keyT    { return keyT{k[0], k[1] - 1} }
func down(k keyT) keyT  { return keyT{k[0], k[1] + 1} }

func findLowPoints(m gridT) func(k keyT, v int, acc []keyT) []keyT {
	return func(k keyT, v int, acc []keyT) []keyT {
		p := []keyT{up(k), down(k), left(k), right(k)}
		if !All(p, isLowPoint(m, v)) {
			return acc
		}
		return append(acc, k)
	}
}

func isLowPoint(m gridT, v int) func(p keyT) bool {
	return func(p keyT) bool {
		dv, ok := m[p]
		return !ok || dv > v
	}
}
