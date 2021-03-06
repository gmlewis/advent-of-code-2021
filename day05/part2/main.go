// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/maps"
	"github.com/gmlewis/advent-of-code-2021/must"
)

// That's the right answer!
// You are one gold star closer to finding the sleigh keys.
// You got rank 843 on this star's leaderboard.

var logf = log.Printf
var printf = fmt.Printf

func main() {
	flag.Parse()

	Each(flag.Args(), process)
}

func process(filename string) {
	logf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)
	b := Reduce(lines, map[string]int{}, parseLines)
	crossings := maps.Count(b, func(k string, v int) bool { return v >= 2 })

	printf("Solution: %v\n", crossings)
}

var lineRE = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func parseLines(line string, b map[string]int) map[string]int {
	m := lineRE.FindStringSubmatch(line)
	if len(m) < 5 {
		log.Fatalf("bad line: %q", line)
	}
	x1, y1, x2, y2 := must.Atoi(m[1]), must.Atoi(m[2]), must.Atoi(m[3]), must.Atoi(m[4])

	pts := Ranges([]int{x1, y1}, []int{x2, y2})
	Each(pts, func(pt []int) { b[fmt.Sprintf("%v,%v", pt[1], pt[0])]++ })

	return b
}
