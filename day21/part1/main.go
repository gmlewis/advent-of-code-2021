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

	ch := make(chan *gameT, 1000)
	go func() {
		g := &gameT{pos1: p1Start, pos2: p2Start}
		for {
			r, p := roll(g.rolls, g.pos1)
			g = &gameT{rolls: r, pos1: p, pos2: g.pos2, score1: g.score1 + p, score2: g.score2}
			ch <- g

			r, p = roll(g.rolls, g.pos2)
			g = &gameT{rolls: r, pos1: g.pos1, pos2: p, score1: g.score1, score2: g.score2 + p}
			ch <- g
		}
	}()

	for game := range ch {
		if game.score1 >= 1000 {
			printf("Solution: %v\n", game.score2*game.rolls)
			break
		}
		if game.score2 >= 1000 {
			printf("Solution: %v\n", game.score1*game.rolls)
			break
		}
	}
}

type gameT struct {
	rolls  int
	pos1   int
	pos2   int
	score1 int
	score2 int
}

func roll(oldRolls, oldPos int) (rolls, pos int) {
	r1 := (oldRolls)%100 + 1
	r2 := (oldRolls+1)%100 + 1
	r3 := (oldRolls+2)%100 + 1
	rolls = oldRolls + 3
	pos = (oldPos+r1+r2+r3-1)%10 + 1
	return rolls, pos
}
