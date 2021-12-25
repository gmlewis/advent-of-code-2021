// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

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
	buf := must.ReadFile(filename)
	pos := Map(strings.Split(buf, ","), must.Atoi)
	max := Max(pos)
	min := Min(pos)
	logf("min=%v, max=%v", min, max)

	f := func(i int) int {
		return Reduce(pos, 0, func(x, acc int) int {
			return acc + mathfn.Abs(i-x)
		})
	}

	bestSum := Min(Map(Range(min, max), f))

	printf("Solution: %v\n", bestSum)
}
