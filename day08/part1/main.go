// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
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

	sum := Sum(Map(lines, count1478))

	printf("Solution: %v\n", sum)
}

func count1478(line string) int {
	parts := strings.Split(line, " | ")
	if len(parts) != 2 {
		log.Fatalf("bad line: %v", line)
	}

	return Count(strings.Split(parts[1], " "), func(w string) bool {
		_, ok := digits[len(w)]
		return ok
	})
}

var digits = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}
