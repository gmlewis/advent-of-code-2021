// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/maps"
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

	m := ReduceWithIndex(lines, gridT{}, func(y int, line string, acc gridT) gridT {
		strfn.RunesWithIndex(line, func(x int, r rune) { acc[keyT{x, y}] = int(r - '0') })
		return acc
	})

	lowPts := maps.Reduce(m, []keyT{}, findLowPoints(m))
	logf("lowPts=%+v", lowPts)
	visited := map[keyT]bool{}
	allBasins := Map(lowPts, calcBasinSize(m, visited))
	logf("allBasins=%+v", allBasins)
	sort.Sort(sort.Reverse(sort.IntSlice(allBasins)))
	answer := Product(allBasins[0:3])

	// 791200 is wrong. = 100 * 92 * 86

	printf("Solution: %v\n", answer)
}

type gridT map[keyT]int

type keyT struct{ x, y int }

func calcBasinSize(m gridT, visited map[keyT]bool) func(k keyT) int {
	return func(k keyT) int {
		v := checkNeighbors(m, k, visited)
		if v == 100 || v == 92 || v == 86 {
			logf("v=%v, lowPt=%v", v, k)
		}
		return v
	}
}

func checkNeighbors(m gridT, k keyT, visited map[keyT]bool) int {
	if visited[k] {
		return 0
	}
	visited[k] = true
	if m[k] == 9 {
		return 0
	}
	acc := 1
	if p := right(k); m[p] == 1+m[k] {
		acc += checkNeighbors(m, p, visited)
	}
	if p := left(k); m[p] == 1+m[k] {
		acc += checkNeighbors(m, p, visited)
	}
	if p := up(k); m[p] == 1+m[k] {
		acc += checkNeighbors(m, p, visited)
	}
	if p := down(k); m[p] == 1+m[k] {
		acc += checkNeighbors(m, p, visited)
	}
	return acc
}

func right(k keyT) keyT { return keyT{x: k.x + 1, y: k.y} }
func left(k keyT) keyT  { return keyT{x: k.x - 1, y: k.y} }
func up(k keyT) keyT    { return keyT{x: k.x, y: k.y - 1} }
func down(k keyT) keyT  { return keyT{x: k.x, y: k.y + 1} }

func findLowPoints(m gridT) func(k keyT, v int, acc []keyT) []keyT {
	d := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	return func(k keyT, v int, acc []keyT) []keyT {
		if !All(d, isLowPoint(m, k, v)) {
			return acc
		}
		return append(acc, k)
	}
}

func isLowPoint(m gridT, k keyT, v int) func(d []int) bool {
	return func(d []int) bool {
		dv, ok := m[keyT{x: k.x + d[0], y: k.y + d[1]}]
		return !ok || dv > v
	}
}
