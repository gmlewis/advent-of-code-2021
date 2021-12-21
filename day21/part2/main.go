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
	p1Start := must.Atoi(strings.TrimSpace(lines[0][len(lines[0])-2:]))
	p2Start := must.Atoi(strings.TrimSpace(lines[1][len(lines[1])-2:]))
	// logf("p1Start: %v, p2Start: %v", p1Start, p2Start)

	fromTo := map[int]map[int]int{}
	for from := 1; from <= 10; from++ {
		fromTo[from] = map[int]int{}
		for r1 := 1; r1 <= 3; r1++ {
			for r2 := 1; r2 <= 3; r2++ {
				for r3 := 1; r3 <= 3; r3++ {
					to := roll(r1, r2, r3, from)
					fromTo[from][to]++
				}
			}
		}
	}
	logf("fromTo=%+v", fromTo)

	p1 := reach21In4(p1Start, nil, nil, fromTo)
	p2 := reach21In4(p2Start, nil, nil, fromTo)
	logf("p1=%v, p2=%v", p1, p2)

	printf("Solution: %v\n", 0)
}

func roll(r1, r2, r3, oldPos int) (pos int) {
	pos = (oldPos+r1+r2+r3-1)%10 + 1
	return pos
}

func reach21In4(pos int, ways, positions []int64, fromTo map[int]map[int]int) int64 {
	if len(positions) == 4 {
		score := positions[0] + positions[1] + positions[2] + positions[3]
		if score >= 21 {
			return ways[0] * ways[1] * ways[2] * ways[3]
		}
		return 0
	}

	var total int64
	for k, v := range fromTo[pos] {
		w := append([]int64{int64(v)}, ways...)
		p := append([]int64{int64(k)}, positions...)
		total += reach21In4(k, w, p, fromTo)
	}
	return total
}
