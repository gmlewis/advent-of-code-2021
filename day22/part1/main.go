// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/mathfn"
	"github.com/gmlewis/advent-of-code-2021/must"
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
	grid := Reduce(lines, gridT{}, processLine)

	if len(grid) == 248314 {
		logf("%+v", grid)
	}

	printf("Solution: %v\n", len(grid))
}

type keyT [3]int
type gridT map[keyT]bool

var lineRE = regexp.MustCompile(`^(\S+) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$`)

const bounds = 50

func processLine(line string, acc gridT) gridT {
	m := lineRE.FindStringSubmatch(line)
	if len(m) != 8 {
		log.Fatalf("unable to parse line: %v", line)
	}
	f := func(k keyT) { acc[k] = true }
	if m[1] == "off" {
		f = func(k keyT) { delete(acc, k) }
	}
	x1 := must.Atoi(m[2])
	x2 := must.Atoi(m[3])
	y1 := must.Atoi(m[4])
	y2 := must.Atoi(m[5])
	z1 := must.Atoi(m[6])
	z2 := must.Atoi(m[7])
	if mathfn.Abs(x1) > bounds || mathfn.Abs(x2) > bounds ||
		mathfn.Abs(y1) > bounds || mathfn.Abs(y2) > bounds ||
		mathfn.Abs(z1) > bounds || mathfn.Abs(z2) > bounds {
		return acc
	}

	for z := z1; z <= z2; z++ {
		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				f(keyT{x, y, z})
			}
		}
	}

	return acc
}
