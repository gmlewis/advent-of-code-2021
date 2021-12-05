// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/maps"
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
	b := enum.Reduce(lines, map[string]int{}, parseLines)
	crossings := maps.Reduce(b, 0, func(k string, v, acc int) int {
		if v >= 2 {
			return acc + 1
		}
		return acc
	})

	printf("Solution: %v\n", crossings)
}

var lineRE = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func parseLines(line string, b map[string]int) map[string]int {
	m := lineRE.FindStringSubmatch(line)
	if len(m) < 5 {
		log.Fatalf("bad line: %q", line)
	}
	x1 := must.Atoi(m[1])
	y1 := must.Atoi(m[2])
	x2 := must.Atoi(m[3])
	y2 := must.Atoi(m[4])

	if x1 == x2 {
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			k := fmt.Sprintf("%v,%v", y, x1)
			b[k]++
		}
	} else if y1 == y2 {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			k := fmt.Sprintf("%v,%v", y1, x)
			b[k]++
		}
	}

	return b
}