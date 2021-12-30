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

	xsize := len(lines[0])
	ysize := len(lines)
	logf("xsize=%v, ysize=%v", xsize, ysize)

	g := &graphT{b: b, stepSize: keyT{xsize, ysize}, target: keyT{5*xsize - 1, 5*ysize - 1}}
	risks := algo.Dijkstra[keyT, int](g, keyT{0, 0}, &g.target, math.MaxInt)

	printf("Solution: %v\n", risks[g.target])
}

type gridT map[keyT]int
type keyT [2]int

// graphT implements the algorithm.Graph[keyT, int] interface.
type graphT struct {
	b        gridT
	stepSize keyT
	target   keyT
}

func (g *graphT) Distance(from, to keyT) int {
	x := to[0] % g.stepSize[0]
	y := to[1] % g.stepSize[1]
	d := to[0]/g.stepSize[0] + to[1]/g.stepSize[1]
	nv := g.b[keyT{x, y}] + d
	if nv > 9 {
		return nv%10 + 1
	}
	return nv
}

func (g *graphT) Each(f func(keyT)) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			for oldK := range g.b {
				k := keyT{oldK[0] + x*g.stepSize[0], oldK[1] + y*g.stepSize[1]}
				f(k)
			}
		}
	}
}

func (g *graphT) EachNeighbor(u keyT, f func(from, to keyT)) {
	if v, ok := moveR(u, g.target); ok {
		f(u, v)
	}
	if v, ok := moveD(u, g.target); ok {
		f(u, v)
	}
	if v, ok := moveL(u, g.target); ok {
		f(u, v)
	}
	if v, ok := moveU(u, g.target); ok {
		f(u, v)
	}
}

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
