// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

	. "github.com/gmlewis/advent-of-code-2021/v1/enum"
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
	numBits := len(lines[0])
	half := len(lines) / 2

	sums := make([]int, numBits)
	sums = Reduce(lines, sums, func(line string, acc []int) []int {
		strfn.RunesWithIndex(line, func(i int, r rune) { acc[i] += int(r - '0') })
		return acc
	})

	var gamma int
	var toggle int
	EachWithIndex(sums, func(i int, sum int) {
		toggle += (1 << (numBits - i - 1))
		if sum >= half {
			gamma += (1 << (numBits - i - 1))
		}
	})
	epsilon := gamma ^ toggle

	printf("Sums: %+v, gamma=%v, toggle=%v, epsilon=%v, product: %v\n", sums, gamma, toggle, epsilon, gamma*epsilon)
}
