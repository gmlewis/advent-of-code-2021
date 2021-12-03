// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

func main() {
	flag.Parse()

	must.Process(process)
}

func process(filename string) {
	lines := must.ReadFileLines(filename)
	position := enum.Reduce(lines, []int{0, 0, 0}, func(line string, acc []int) []int {
		switch {
		case strings.HasPrefix(line, "forward "):
			v := must.Atoi(line[8:])
			return []int{acc[0] + v, acc[1] + v*acc[2], acc[2]}
		case strings.HasPrefix(line, "down "):
			v := must.Atoi(line[5:])
			return []int{acc[0], acc[1], acc[2] + v}
		case strings.HasPrefix(line, "up "):
			v := must.Atoi(line[3:])
			return []int{acc[0], acc[1], acc[2] - v}
		}
		return acc
	})
	fmt.Printf("Solution: %v - product: %v\n", position, position[0]*position[1])
}
