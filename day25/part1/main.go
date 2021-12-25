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

	ReduceWithIndex(lines, puz, parseLine)

	for steps := 1; true; steps++ {
		if puz.step() == 0 {
			printf("Solution: %v\n", steps)
			break
		}
	}
}

type keyFunc func(keyT) keyT
type keyT [2]int
type puzT struct {
	b     map[keyT]rune
	xsize int
	ysize int
}

func (p *puzT) step() (count int) {
	nb := map[keyT]rune{}
	// east-facing
	for k, v := range p.b {
		if v != '>' {
			nb[k] = v
			continue
		}
		if rk := p.right(k); p.b[rk] == 0 {
			nb[rk] = v
			count++
		} else {
			nb[k] = v
		}
	}
	// south-facing
	nb2 := map[keyT]rune{}
	for k, v := range nb {
		if v != 'v' {
			nb2[k] = v
			continue
		}
		if dk := p.down(k); nb[dk] == 0 {
			nb2[dk] = v
			count++
		} else {
			nb2[k] = v
		}
	}
	p.b = nb2

	return count
}

func (p *puzT) right(k keyT) keyT {
	return keyT{(k[0] + 1) % p.xsize, k[1]}
}

func (p *puzT) down(k keyT) keyT {
	return keyT{k[0], (k[1] + 1) % p.ysize}
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
