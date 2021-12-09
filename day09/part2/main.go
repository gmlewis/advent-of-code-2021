// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"

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
	logf("%v lowPts", len(lowPts))
	allBasins := Map(lowPts, calcBasinSize(m))
	sort.Sort(sort.Reverse(sort.IntSlice(allBasins)))
	logf("allBasins=%+v", allBasins)
	answer := Product(allBasins[0:3])

	// 791200 is wrong. = 100 * 92 * 86
	// 640872 is wrong. = 92 * 86 * 81
	// 564246 is wrong. = 86 * 81 * 81
	// 531441 = 81 * 81 * 81

	printf("Solution: %v\n", answer)
}

type gridT map[keyT]int

type keyT struct{ x, y int }

func less(k1, k2 keyT) bool {
	if k1.y == k2.y {
		return k1.x < k2.x
	}
	return k1.y < k2.y
}

func calcBasinSize(m gridT) func(k keyT) int {
	return func(k keyT) int {
		visited := map[keyT]bool{}
		debug := gridT{}
		v := checkNeighbors(m, k, visited, debug)
		if len(debug) > 81 {
			logf("v=%v, lowPt=%v, len(debug)=%v", v, k, len(debug))
			prettyPrint(debug)
		}
		return v
	}
}

func prettyPrint(debug gridT) {
	keys := maps.Keys(debug)
	minX := ReduceWithIndex(keys, 0, func(i int, k keyT, acc int) int {
		if i == 0 || k.x < acc {
			return k.x
		}
		return acc
	})
	maxX := ReduceWithIndex(keys, 0, func(i int, k keyT, acc int) int {
		if i == 0 || k.x > acc {
			return k.x
		}
		return acc
	})
	minKey := MinFunc(keys, less)
	minY := minKey.y
	maxKey := MaxFunc(keys, less)
	maxY := maxKey.y

	sort.Slice(keys, func(a, b int) bool { return less(keys[a], keys[b]) })
	printf("keys=%v, min=(%v,%v), max=(%v,%v)\n", len(keys), minX, minY, maxX, maxY)

	var ret string
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if v, ok := debug[keyT{x, y}]; ok {
				ret += strconv.Itoa(v)
			} else {
				ret += " "
			}
		}
		ret += "\n"
	}
	printf("%v\n", ret)
}

func checkNeighbors(m gridT, k keyT, visited map[keyT]bool, debug gridT) int {
	visited[k] = true
	if m[k] == 9 {
		return 0
	}
	debug[k] = m[k]
	acc := 1
	if p := right(k); !visited[p] && m[p] == 1+m[k] {
		acc += checkNeighbors(m, p, visited, debug)
	}
	if p := left(k); !visited[p] && m[p] == 1+m[k] {
		acc += checkNeighbors(m, p, visited, debug)
	}
	if p := up(k); !visited[p] && m[p] == 1+m[k] {
		acc += checkNeighbors(m, p, visited, debug)
	}
	if p := down(k); !visited[p] && m[p] == 1+m[k] {
		acc += checkNeighbors(m, p, visited, debug)
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
