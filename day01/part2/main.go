// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

var printf = fmt.Printf

func main() {
	flag.Parse()

	enum.Each(flag.Args(), process)
}

func process(filename string) {
	log.Printf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)
	vals := enum.Map(lines, must.Atoi)
	groups := enum.ChunkEvery(vals, 4, 1)
	count := enum.Count(groups, func(group []int) bool {
		return enum.Sum(group[1:4]) > enum.Sum(group[0:3])
	})
	printf("Solution: %v\n", count)
}
