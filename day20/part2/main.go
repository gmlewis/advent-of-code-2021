// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/v1/enum"
	"github.com/gmlewis/advent-of-code-2021/v1/maps"
	"github.com/gmlewis/advent-of-code-2021/v1/must"
	"github.com/gmlewis/advent-of-code-2021/v1/stream"
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
	filter := ReduceWithIndex([]rune(lines[0]), filterT{}, func(i int, r rune, acc filterT) filterT {
		if r == '#' {
			acc[i] = 1
		} else {
			acc[i] = 0
		}
		return acc
	})
	img := parse(lines[2:])

	for i := 0; i < 50; i++ {
		background := 0
		if i%2 == 1 {
			background = filter[0]
		}
		img = img.enhance(filter, background)
	}

	printf("Solution: %v\n", img.lit())
}

type filterT map[int]int
type keyT [2]int
type pixelsT map[keyT]int
type imageT struct {
	p    pixelsT
	xmin int
	xmax int
	ymin int
	ymax int
}
type pixelT [3]int // x, y, value

func (i *imageT) lit() int {
	return maps.Count(i.p, func(k keyT, v int) bool { return v == 1 })
}

func (i *imageT) enhance(filter filterT, background int) *imageT {
	ch := make(chan keyT, 1000)

	go func() {
		for y := i.ymin - 1; y <= i.ymax+1; y++ {
			for x := i.xmin - 1; x <= i.xmax+1; x++ {
				ch <- keyT{x, y}
			}
		}
		close(ch)
	}()

	f := func(k keyT) int {
		if v, ok := i.p[k]; ok {
			return v
		}
		return background
	}

	result := &imageT{p: pixelsT{}, xmin: i.xmin - 1, xmax: i.xmax + 1, ymin: i.ymin - 1, ymax: i.ymax + 1}
	pixelCh := stream.MapStream(ch, func(p keyT) pixelT {
		bits := f(keyT{p[0] - 1, p[1] - 1})<<8 |
			f(keyT{p[0], p[1] - 1})<<7 |
			f(keyT{p[0] + 1, p[1] - 1})<<6 |
			f(keyT{p[0] - 1, p[1]})<<5 |
			f(keyT{p[0], p[1]})<<4 |
			f(keyT{p[0] + 1, p[1]})<<3 |
			f(keyT{p[0] - 1, p[1] + 1})<<2 |
			f(keyT{p[0], p[1] + 1})<<1 |
			f(keyT{p[0] + 1, p[1] + 1})
		return pixelT{p[0], p[1], filter[bits]}
	})

	for p := range pixelCh {
		result.p[keyT{p[0], p[1]}] = p[2]
	}

	return result
}

func (i *imageT) String() string {
	lines := make([]string, 0, 1+i.ymax-i.ymin)
	for y := i.ymin; y <= i.ymax; y++ {
		var line string
		for x := i.xmin; x <= i.xmax; x++ {
			if i.p[keyT{x, y}] == 1 {
				line += "#"
				continue
			}
			line += "."
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func parse(lines []string) *imageT {
	img := &imageT{p: pixelsT{}, ymax: len(lines) - 1, xmax: len(lines[0])}
	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				img.p[keyT{x, y}] = 1
			}
		}
	}
	return img
}
