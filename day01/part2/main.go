// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"

	"github.com/gmlewis/advent-of-code-2021/must"
)

func main() {
	flag.Parse()

	must.Process(process)
}

func process(filename string) {
	lines := must.ReadFileLines(filename)
	vals := must.StringToInt(lines)
	groups := must.ChunkEveryInt(vals, 4, 1)
	count := must.ReduceIntSlicesToInt(groups, 0, func(group []int, acc int) int {
		if group[1]+group[2]+group[3] > group[0]+group[1]+group[2] {
			return acc + 1
		}
		return acc
	})
	fmt.Printf("Solution: %v\n", count)
}
