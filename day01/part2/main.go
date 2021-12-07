// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

var printf = fmt.Printf

func main() {
	flag.Parse()

	Each(flag.Args(), process)
}

func process(filename string) {
	log.Printf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)
	vals := Map(lines, must.Atoi)
	groups := ChunkEvery(vals, 4, 1)
	count := Count(groups, func(group []int) bool {
		return Sum(group[1:4]) > Sum(group[0:3])
	})
	printf("Solution: %v\n", count)
}
