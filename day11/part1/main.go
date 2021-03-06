// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

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

	sim := Reduce(Range(1, 100), &accT{b: b}, simStep)

	printf("Solution: %v\n", sim.flashes)
}

func flash(k keyT, acc *accT) {
	d := [][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	acc.flashes++
	Each(d, func(v [2]int) { incr(keyT{k[0] + v[0], k[1] + v[1]}, acc) })
}

func incr(k keyT, acc *accT) {
	if k[0] < 0 || k[0] > 9 || k[1] < 0 || k[1] > 9 {
		return
	}
	acc.b[k]++
	if acc.b[k] == 10 {
		flash(k, acc)
	}
}

func simStep(step int, acc *accT) *accT {
	for k := range acc.b {
		incr(k, acc)
	}
	for k, v := range acc.b {
		if v > 9 {
			acc.b[k] = 0
		}
	}

	return acc
}

type gridT map[keyT]int

type keyT [2]int

type accT struct {
	b       gridT
	flashes int
}
