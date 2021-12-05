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
	pairs := enum.ChunkEvery(vals, 2, 1)
	count := enum.Count(pairs, func(pair []int) bool { return pair[1] > pair[0] })
	printf("Solution: %v\n", count)
}
