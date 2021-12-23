// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"sort"

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
	p = p.solve(math.MaxInt)
	logf("\n%v", p)

	printf("Solution: %v\n", p.energy)
}

type keyT [2]int
type puzT struct {
	energy   int
	landings map[keyT]rune
	inMotion map[keyT]rune
}

func (p *puzT) solve(bestEnergy int) *puzT {
	if p.isSolved() {
		return p
	}

	moves := p.allPossibleMoves()
	if len(moves) == 0 {
		return nil
	}

	// logf("solve(%v): %v allPossibleMoves: %+v", bestEnergy, p, moves)
	var best *puzT
	for _, move := range moves {
		f := move.from
		t := move.to
		e := move.energy
		if e+p.energy >= bestEnergy {
			continue
		}

		// logf("Moving '%c' from %+v to %+v using %v energy", p.inMotion[f], f, t, e)
		np := &puzT{energy: e + p.energy, landings: dup(p.landings), inMotion: dup(p.inMotion)}
		if t[1] == 0 || arrivedX[p.inMotion[f]] != t[0] {
			np.inMotion[t] = np.inMotion[f]
		}
		delete(np.inMotion, f)
		np.landings[t] = np.landings[f]
		delete(np.landings, f)

		np = np.solve(bestEnergy)
		if np != nil && np.energy < bestEnergy {
			best = np
			bestEnergy = np.energy
			// logf("NEW BEST ENERGY: %v", bestEnergy)
		}
	}
	return best
}

type moveT struct {
	from   keyT
	to     keyT
	energy int
}

func (p *puzT) allPossibleMoves() (moves []moveT) {
	for k := range p.inMotion {
		moves = append(moves, p.possibleMoves(k)...)
	}
	sort.Slice(moves, func(a, b int) bool { return moves[a].energy < moves[b].energy })
	return moves
}

func (p *puzT) possibleMoves(from keyT) (moves []moveT) {
	r := p.inMotion[from]
	if r == 0 {
		log.Fatalf("possibleMoves(%+v) not in p.inMotion=%#v", from, p.inMotion)
	}
	roomX := 2*int(r-'A') + 2
	if from[1] == 0 {
		// logf("%c at %+v must move from hallway into its own roomX=%v", r, from, roomX)
		to := keyT{roomX, 2}
		// logf("Can %c move from %+v to %+v?", r, from, to)
		if p.landings[to] == 0 && p.clearPath(from, to) {
			return []moveT{{from: from, to: to, energy: energy(r, from, to)}}
		}
		to = keyT{roomX, 1}
		// logf("No. Can %c move from %+v to %+v?", r, from, to)
		if p.landings[to] == 0 && p.clearPath(from, to) {
			return []moveT{{from: from, to: to, energy: energy(r, from, to)}}
		}
		// logf("No.")
		return nil
	}

	// Moving from room into hallway.
	if from[1] == 2 && p.landings[keyT{from[0], 1}] != 0 {
		// logf("%c at %+v is blocked from moving into hallway", r, from)
		return nil // blocked
	}

	// Can this be a final move into place?
	column := roomX - 1
	if from[0] > roomX {
		column = roomX + 1
	}
	// logf("Can %c move from %+v to %+v?", r, from, keyT{column, 0})
	if p.clearPath(from, keyT{column, 0}) {
		// logf("Yes! Can %c move from %+v to %+v?", r, keyT{column, 0}, keyT{roomX, 2})
		if p.clearPath(keyT{column, 0}, keyT{roomX, 2}) {
			// logf("Yes!")
			to := keyT{roomX, 2}
			return []moveT{{from: from, to: to, energy: energy(r, from, to)}}
		}
		// logf("No. Can %c move instead from %+v to %+v?", r, keyT{column, 0}, keyT{roomX, 1})
		if p.clearPath(keyT{column, 0}, keyT{roomX, 1}) {
			// logf("Yes!")
			to := keyT{roomX, 1}
			return []moveT{{from: from, to: to, energy: energy(r, from, to)}}
		}
	}
	// logf("No.")

	f := func(x int) {
		to := keyT{x, 0}
		if p.clearPath(from, to) {
			moves = append(moves, moveT{from: from, to: to, energy: energy(r, from, to)})
		}
	}

	for _, x := range orderX[from[0]] {
		f(x)
	}

	return moves
}

var orderX = map[int][]int{
	2: []int{1, 3, 0, 5, 7, 9, 10},
	4: []int{3, 5, 1, 7, 0, 9, 10},
	6: []int{5, 7, 3, 9, 10, 1, 0},
	8: []int{7, 9, 10, 5, 3, 1, 0},
}

var arrivedX = map[rune]int{'A': 2, 'B': 4, 'C': 6, 'D': 8}

func (p *puzT) clearPath(from, to keyT) bool {
	if p.landings[to] != 0 {
		return false
	}

	if to[1] != 0 {
		if p.landings[keyT{to[0], 1}] != 0 || p.landings[to] != 0 {
			return false
		}
		for x := from[0] - 1; x > to[0]; x-- {
			if p.landings[keyT{x, 0}] != 0 {
				return false
			}
		}
		for x := from[0] + 1; x < to[0]; x++ {
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

	for x := from[0] - 1; x > to[0]; x-- {
		if p.landings[keyT{x, 0}] != 0 {
			return false
		}
	}
	for x := from[0] + 1; x < to[0]; x++ {
		if p.landings[keyT{x, 0}] != 0 {
			return false
		}
	}
	return true
}

func energy(r rune, from, to keyT) int {
	if from[1] > 0 && to[1] > 0 {
		dist := mathfn.Abs(from[0]-to[0]) + from[1] + to[1]
		return energyPerStep[r] * dist
	}

	dist := mathfn.Abs(from[0]-to[0]) + mathfn.Abs(from[1]-to[1])
	return energyPerStep[r] * dist
}

var energyPerStep = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

func dup(src map[keyT]rune) map[keyT]rune {
	dst := map[keyT]rune{}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func (p *puzT) isSolved() bool {
	if len(p.landings) != 8 {
		log.Fatalf("lost an amphipod! %v", p)
	}
	if len(p.inMotion) != 0 {
		return false
	}
	return p.landings[keyT{2, 2}] == 'A' && p.landings[keyT{2, 1}] == 'A' &&
		p.landings[keyT{4, 2}] == 'B' && p.landings[keyT{4, 1}] == 'B' &&
		p.landings[keyT{6, 2}] == 'C' && p.landings[keyT{6, 1}] == 'C' &&
		p.landings[keyT{8, 2}] == 'D' && p.landings[keyT{8, 1}] == 'D'
}

func parse(lines []string) *puzT {
	p := &puzT{landings: map[keyT]rune{}, inMotion: map[keyT]rune{}}

	f := func(x, y int) {
		k := keyT{x, y}
		r := rune(lines[y+1][x+1])
		if r == '.' {
			return
		}
		p.landings[k] = r
		if y != 2 || arrivedX[r] != x {
			p.inMotion[k] = r
		}
	}

	for y := 1; y <= 2; y++ {
		for x := 2; x <= 8; x += 2 {
			f(x, y)
		}
	}
	for x := 0; x <= 10; x++ {
		f(x, 0)
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

	return fmt.Sprintf(`energy=%v
#%v#
###%c#%c#%c#%c###
  #%c#%c#%c#%c#
`, p.energy, landings, f(2, 1), f(4, 1), f(6, 1), f(8, 1), f(2, 2), f(4, 2), f(6, 2), f(8, 2))
}
