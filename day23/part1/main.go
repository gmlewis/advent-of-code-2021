// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

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
	p := parse(lines)
	logf("\n%v", p)

	printf("Solution: %v\n", p.energy)
}

type keyT [2]int
type puzT struct {
	energy   int
	landings map[keyT]rune
	inMotion map[keyT]rune
}

func parse(lines []string) *puzT {
	p := &puzT{landings: map[keyT]rune{}, inMotion: map[keyT]rune{}}

	f := func(x, y int) {
		k := keyT{x, y}
		r := rune(lines[y+1][x+1])
		p.landings[k] = r
		p.inMotion[k] = r
	}

	for y := 1; y <= 2; y++ {
		for x := 2; x <= 8; x += 2 {
			f(x, y)
		}
	}

	return p
}

func (p *puzT) String() string {
	var landings string
	for x := 0; x <= 10; x++ {
		r, ok := p.landings[keyT{x, 0}]
		if ok {
			landings += string(r)
			continue
		}
		landings += "."
	}

	f := func(x, y int) rune {
		r, ok := p.landings[keyT{x, y}]
		if !ok {
			return '.'
		}
		return r
	}

	return fmt.Sprintf(`#############
#%v#
###%c#%c#%c#%c###
  #%c#%c#%c#%c#
  #########
`, landings, f(2, 1), f(4, 1), f(6, 1), f(8, 1), f(2, 2), f(4, 2), f(6, 2), f(8, 2))
}
