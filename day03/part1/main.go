// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

func main() {
	flag.Parse()

	enum.Each(flag.Args(), process)
}

func process(filename string) {
	log.Printf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)
	numBits := len(lines[0])
	half := len(lines) / 2

	sums := make([]int, numBits)
	sums = enum.Reduce(lines, sums, func(line string, acc []int) []int {
		enum.RunesWithIndex(line, func(i int, r rune) { acc[i] += int(r - '0') })
		return acc
	})

	var gamma int
	var toggle int
	enum.EachWithIndex(sums, func(i int, sum int) {
		toggle += (1 << (numBits - i - 1))
		if sum >= half {
			gamma += (1 << (numBits - i - 1))
		}
	})
	epsilon := gamma ^ toggle

	fmt.Printf("Sums: %+v, gamma=%v, toggle=%v, epsilon=%v, product: %v\n", sums, gamma, toggle, epsilon, gamma*epsilon)
}
