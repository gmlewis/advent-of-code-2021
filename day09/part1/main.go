// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

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

	risk := maps.Reduce(m, 0, riskLevel(m))

	printf("Solution: %v\n", risk)
}

type gridT map[keyT]int

type keyT [2]int

func riskLevel(m gridT) func(k keyT, v, acc int) int {
	d := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	return func(k keyT, v, acc int) int {
		if !All(d, isLowPoint(m, k, v)) {
			return acc
		}
		return acc + v + 1
	}
}

func isLowPoint(m gridT, k keyT, v int) func(d []int) bool {
	return func(d []int) bool {
		dv, ok := m[keyT{k[0] + d[0], k[1] + d[1]}]
		return !ok || dv > v
	}
}
