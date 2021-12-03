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
	position := enum.Reduce(lines, []int{0, 0}, processLine)
	fmt.Printf("Solution: %v - product: %v\n", position, position[0]*position[1])
}

func processLine(line string, acc []int) []int {
	parts := strings.Split(line, " ")
	v := must.Atoi(parts[1])
	switch parts[0] {
	case "forward":
		return []int{acc[0] + v, acc[1]}
	case "down":
		return []int{acc[0], acc[1] + v}
	case "up":
		return []int{acc[0], acc[1] - v}
	}
	return acc
}
