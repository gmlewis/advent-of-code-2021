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

	p1wins, p2wins := player1(p1Start, 0, 1, p2Start, 0, 1, fromTo)
	logf("p1wins=%v, p2wins=%v", p1wins, p2wins)

	printf("Solution: %v %v\n", p1wins, p2wins)
}

func roll(r1, r2, r3, oldPos int) (pos int) {
	pos = (oldPos+r1+r2+r3-1)%10 + 1
	return pos
}

func player1(p1pos, p1score int, p1ways int64,
	p2pos, p2score int, p2ways int64,
	fromTo map[int]map[int]int) (p1wins, p2wins int64) {

	for k, v := range fromTo[p1pos] {
		if p1score+k >= 21 {
			return p1ways * 27, p2ways
		}

		p1, p2 := player2(k, p1score+k, p1ways*int64(v), p2pos, p2score, p2ways, fromTo)
		p1wins += p1
		p2wins += p2
	}
	return p1wins, p2wins
}

func player2(p1pos, p1score int, p1ways int64,
	p2pos, p2score int, p2ways int64,
	fromTo map[int]map[int]int) (p1wins, p2wins int64) {

	for k, v := range fromTo[p2pos] {
		if p2score+k >= 21 {
			return p1ways, p2ways * 27
		}

		p1, p2 := player1(p1pos, p1score, p1ways, k, p2score+k, p2ways*int64(v), fromTo)
		p1wins += p1
		p2wins += p2
	}
	return p1wins, p2wins
}
