// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"sort"
	"sync"

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

var mu sync.RWMutex // protects bestEnergy
var bestEnergy int

func process(filename string) {
	logf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)
	p := parse(lines)
	logf("START:\n%v", p)
	bestEnergy = math.MaxInt
	p.solve()

	printf("Solution: %v\n", bestEnergy)
}

type keyT [2]int
type puzT struct {
	energy   int
	landings map[keyT]rune
	inMotion map[keyT]rune
}

func (p *puzT) solve() *puzT {
	if p.isSolved() {
		return p
	}

	moves := p.allPossibleMoves()
	if len(moves) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	throttleCh := make(chan struct{}, 3)
	ch := make(chan *puzT, 10)

	for _, move := range moves {
		f := move.from
		t := move.to
		e := move.energy
		mu.RLock()
		if e+p.energy >= bestEnergy {
			mu.RUnlock()
			continue
		}
		mu.RUnlock()

		wg.Add(1)
		throttleCh <- struct{}{}
		go func(f, t keyT, e int) {
			np := &puzT{energy: e + p.energy, landings: dup(p.landings), inMotion: dup(p.inMotion)}
			if t[1] == 0 || arrivedX[p.inMotion[f]] != t[0] {
				np.inMotion[t] = np.inMotion[f]
			}
			delete(np.inMotion, f)
			np.landings[t] = np.landings[f]
			delete(np.landings, f)

			np = np.solve()
			if np != nil {
				ch <- np
			}
			<-throttleCh
			wg.Done()
		}(f, t, e)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var best *puzT
	for np := range ch {
		mu.Lock()
		if np.energy < bestEnergy {
			best = &puzT{energy: np.energy}
			bestEnergy = np.energy
			logf("NEW BEST ENERGY: %v", bestEnergy)
		}
		mu.Unlock()
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
	sort.Slice(moves, func(a, b int) bool { return moves[a].from[1] > moves[b].from[1] })
	return moves
}

func (p *puzT) possibleMoves(from keyT) (moves []moveT) {
	r := p.inMotion[from]
	if r == 0 {
		log.Fatalf("possibleMoves(%+v) not in p.inMotion=%#v", from, p.inMotion)
	}
	roomX := 2*int(r-'A') + 2
	if from[1] == 0 {
		if p.landings[keyT{roomX, 4}] != r && p.landings[keyT{roomX, 4}] != 0 {
			return nil // column is stil blocked
		}

		for y := 4; y >= 1; y-- {
			to := keyT{roomX, y}
			if p.clearPath(from, to) {
				return []moveT{{from: from, to: to, energy: energy(r, from, to)}}
			}
		}

		return nil
	}

	if !p.clearPath(from, keyT{from[0], 0}) {
		return nil // blocked
	}

	column := roomX - 1
	if from[0] > roomX {
		column = roomX + 1
	}
	if p.clearPath(from, keyT{column, 0}) && (p.landings[keyT{roomX, 4}] == r || p.landings[keyT{roomX, 4}] == 0) {
		for y := 4; y >= 1; y-- {
			if p.clearPath(keyT{column, 0}, keyT{roomX, y}) {
				to := keyT{roomX, y}
				moves = append(moves, moveT{from: from, to: to, energy: energy(r, from, to)})
				break
			}
		}
	}

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
var xToRune = map[int]rune{2: 'A', 4: 'B', 6: 'C', 8: 'D'}

func (p *puzT) clearPath(from, to keyT) bool {
	if from == to {
		return true
	}
	if p.landings[to] != 0 {
		return false
	}

	switch {
	case from[0] < to[0] && to[1] == 0:
		return p.clearPath(from, keyT{to[0] - 1, to[1]})
	case from[0] > to[0] && to[1] == 0:
		return p.clearPath(from, keyT{to[0] + 1, to[1]})
	case from[0] != to[0] && to[1] > 0:
		return p.clearPath(from, keyT{to[0], to[1] - 1})
	case from[0] == to[0]:
		return p.clearPath(from, keyT{to[0], to[1] + 1})
	default:
		log.Fatalf("unhandled clearPath: from=%+v, to=%+v", from, to)
	}
	return false
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
	if len(p.landings) != 16 {
		log.Fatalf("lost an amphipod! %v", p)
	}
	if len(p.inMotion) != 0 {
		return false
	}
	for y := 1; y <= 4; y++ {
		for x := 2; x <= 8; x += 2 {
			if p.landings[keyT{x, y}] != xToRune[x] {
				return false
			}
		}
	}
	return true
}

func parse(lines []string) *puzT {
	lines = append(lines[:4], lines[2], lines[3])
	lines[3] = "  #D#C#B#A#"
	lines[4] = "  #D#B#A#C#"
	return parseLiteral(lines)
}

func parseLiteral(lines []string) *puzT {
	p := &puzT{landings: map[keyT]rune{}, inMotion: map[keyT]rune{}}

	f := func(x, y int) {
		k := keyT{x, y}
		r := rune(lines[y+1][x+1])
		if r == '.' {
			return
		}
		p.landings[k] = r
		if y < 4 || arrivedX[r] != x {
			p.inMotion[k] = r
		}
	}

	for y := 1; y <= 4; y++ {
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
#############
#%v#
###%c#%c#%c#%c###
  #%c#%c#%c#%c#
  #%c#%c#%c#%c#
  #%c#%c#%c#%c#
  #########
`, p.energy, landings,
		f(2, 1), f(4, 1), f(6, 1), f(8, 1),
		f(2, 2), f(4, 2), f(6, 2), f(8, 2),
		f(2, 3), f(4, 3), f(6, 3), f(8, 3),
		f(2, 4), f(4, 4), f(6, 4), f(8, 4),
	)
}
