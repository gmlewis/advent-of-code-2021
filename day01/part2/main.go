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
	vals := enum.Map(lines, must.Atoi)
	groups := enum.ChunkEvery(vals, 4, 1)
	count := enum.Reduce(groups, 0, func(group []int, acc int) int {
		if enum.Sum(group[1:4]) > enum.Sum(group[0:3]) {
			return acc + 1
		}
		return acc
	})
	fmt.Printf("Solution: %v\n", count)
}
