// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/mathfn"
	"github.com/gmlewis/advent-of-code-2021/must"
)

var printf = fmt.Printf

func main() {
	flag.Parse()

	enum.Each(flag.Args(), process)
}

func process(filename string) {
	log.Printf("Processing %v ...", filename)
	buf := must.ReadFile(filename)
	pos := enum.Map(strings.Split(buf, ","), must.Atoi)
	max := enum.Max(pos)
	min := enum.Min(pos)
	log.Printf("min=%v, max=%v", min, max)

	var bestSum int
	for i := min; i <= max; i++ {
		sum := enum.Reduce(pos, 0, func(x, acc int) int {
			d := mathfn.Abs(i - x)
			return acc + (d * (d + 1) / 2)
		})
		// log.Printf("i=%v, sum=%v", i, sum)
		if i == min || sum < bestSum {
			bestSum = sum
		}
	}

	printf("Solution: %v\n", bestSum)
}
