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

var printf = fmt.Printf

func main() {
	flag.Parse()

	Each(flag.Args(), process)
}

func process(filename string) {
	log.Printf("Processing %v ...", filename)
	buf := must.ReadFile(filename)
	pos := Map(strings.Split(buf, ","), must.Atoi)
	max := Max(pos)
	min := Min(pos)
	log.Printf("min=%v, max=%v", min, max)

	f := func(i int) int {
		return Reduce(pos, 0, func(x, acc int) int {
			d := mathfn.Abs(i - x)
			return acc + (d * (d + 1) / 2)
		})
	}

	bestSum := Min(Map(Range(min, max), f))

	printf("Solution: %v\n", bestSum)
}
