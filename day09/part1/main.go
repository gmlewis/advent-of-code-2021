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

	m := ReduceWithIndex(lines, gridT{}, func(y int, line string, acc gridT) gridT {
		strfn.RunesWithIndex(line, func(x int, r rune) { acc[key(x, y)] = int(r - '0') })
		return acc
	})
	log.Printf("m=%+v", m)

	risk := maps.Reduce(m, 0, riskLevel(m))

	printf("Solution: %v\n", risk)
}

type gridT map[string]int

func key(x, y int) string { return fmt.Sprintf("%v,%v", y, x) }

func riskLevel(m gridT) func(k, v) int {
	return func(k, v) int {
		const d = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
		if !All(d, isLowPoint(m, k, v)) { return 0 }
		return v + 1
	}
}

func isLowPoint(m gridT, k string, v int) bool {
	p := strings.
}
