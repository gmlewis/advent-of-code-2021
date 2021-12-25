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

	puz := &puzT{
		b:     map[keyT]rune{},
		xsize: len(lines[0]),
		ysize: len(lines),
	}
	puz.right = func(k keyT) keyT {
		return keyT{(k[0] + 1) % puz.xsize, k[1]}
	}
	puz.down = func(k keyT) keyT {
		return keyT{k[0], (k[1] + 1) % puz.ysize}
	}
	logf("xsize=%v, ysize=%v", puz.xsize, puz.ysize)

	ReduceWithIndex(lines, puz, parseLine)
	logf("puz:\n%v", puz)

	for steps := 1; true; steps++ {
		np, count := puz.step()
		logf("step=%v, count=%v", steps, count)
		if count == 0 {
			printf("Solution: %v\n", steps)
			break
		}
		puz = np
	}
}

type keyFunc func(keyT) keyT
type keyT [2]int
type puzT struct {
	b     map[keyT]rune
	xsize int
	ysize int
	right keyFunc
	down  keyFunc
}

func (p *puzT) step() (np *puzT, count int) {
	np = &puzT{b: map[keyT]rune{}, xsize: p.xsize, ysize: p.ysize, right: p.right, down: p.down}
	// east-facing
	for k, v := range p.b {
		if v != '>' {
			np.b[k] = v
			continue
		}
		if rk := p.right(k); p.b[rk] == 0 {
			np.b[rk] = v
			count++
		} else {
			np.b[k] = v
		}
	}
	// logf("after east-facing puz:\n%v", np)
	// south-facing
	np2 := &puzT{b: map[keyT]rune{}, xsize: p.xsize, ysize: p.ysize, right: p.right, down: p.down}
	for k, v := range np.b {
		if v != 'v' {
			np2.b[k] = v
			continue
		}
		if dk := np.down(k); np.b[dk] == 0 {
			np2.b[dk] = v
			count++
		} else {
			np2.b[k] = v
		}
	}
	// logf("after south-facing puz:\n%v", np2)

	return np2, count
}

func (p *puzT) String() string {
	var lines []string
	for y := 0; y < p.ysize; y++ {
		var line string
		for x := 0; x < p.xsize; x++ {
			k := keyT{x, y}
			r := p.b[k]
			if r == 0 {
				line += "."
				continue
			}
			line += string(r)
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func parseLine(y int, line string, acc *puzT) *puzT {
	for x, r := range line {
		if r == '.' {
			continue
		}
		acc.b[keyT{x, y}] = r
	}
	return acc
}
