// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/mathfn"
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

func (p *puzT) possibleMoves(from keyT) (moves []keyT, energies []int) {
	r := p.inMotion[from]
	if from[1] == 0 {
		// Must move from hallway into its own room.
		roomX := 2*int(r-'A') + 2
		to := keyT{roomX, 2}
		if p.landings[to] == 0 && p.clearPath(from, to) {
			return []keyT{to}, []int{energy(r, from, to)}
		}
		to = keyT{roomX, 1}
		if p.landings[to] == 0 && p.clearPath(from, to) {
			return []keyT{to}, []int{energy(r, from, to)}
		}
		return nil, nil
	}

	// Moving from room into hallway.
	if from[1] == 2 && p.landings[keyT{from[0], 1}] != 0 {
		return nil, nil // blocked
	}

	f := func(x int) {
		to := keyT{x, 0}
		if p.clearPath(from, to) {
			moves = append(moves, to)
			energies = append(energies, energy(r, from, to))
		}
	}

	for _, x := range orderX[from[0]] {
		f(x)
	}

	return moves, energies
}

var orderX = map[int][]int{
	2: []int{1, 3, 0, 5, 7, 9, 10},
	4: []int{3, 5, 1, 7, 0, 9, 10},
	6: []int{5, 7, 3, 9, 10, 1, 0},
	8: []int{7, 9, 10, 5, 3, 1, 0},
}

func (p *puzT) clearPath(from, to keyT) bool {
	if to[1] != 0 {
		// Moving from hallway to room.
		if p.landings[keyT{to[0], 1}] != 0 || p.landings[to] != 0 {
			return false
		}
		for x := from[0]; x > to[0]; x-- {
			if p.landings[keyT{x, 0}] != 0 {
				return false
			}
		}
		for x := from[0]; x < to[0]; x++ {
			if p.landings[keyT{x, 0}] != 0 {
				return false
			}
		}
		return true
	}

	// Moving from room into hallway.
	if from[1] == 2 && p.landings[keyT{from[0], 1}] != 0 {
		return false
	}
	for _, x := range orderX[from[0]] {
		if x == to[0] && p.landings[to] == 0 {
			return true
		}
		if p.landings[keyT{x, 0}] != 0 {
			return false
		}
	}
	return false
}

func energy(r rune, from, to keyT) int {
	dist := mathfn.Abs(from[0]-to[0]) + mathfn.Abs(from[1]-to[1])
	return energyPerStep[r] * dist
}

var energyPerStep = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
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
