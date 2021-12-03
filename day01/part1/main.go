// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

func main() {
	flag.Parse()

	must.Process(process)
}

func process(filename string) {
	lines := must.ReadFileLines(filename)
	vals := enum.Map(lines, must.Atoi)
	pairs := enum.ChunkEvery(vals, 2, 1)
	count := enum.Reduce(pairs, 0, func(pair []int, acc int) int {
		if pair[1] > pair[0] {
			return acc + 1
		}
		return acc
	})
	fmt.Printf("Solution: %v\n", count)
}
