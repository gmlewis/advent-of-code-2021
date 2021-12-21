// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/gmlewis/advent-of-code-2021/stream"
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
		}
		return acc
	})
	img := parse(lines[2:])

	after2 := img.enhance(filter).enhance(filter)

	printf("Solution: %v\n", len(after2.p))
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

func (i *imageT) enhance(filter filterT) *imageT {
	ch := make(chan keyT, 1000)

	go func() {
		for y := i.ymin - 1; y <= i.ymax+1; y++ {
			for x := i.xmin - 1; x <= i.xmax+1; x++ {
				ch <- keyT{x, y}
			}
		}
		close(ch)
	}()

	result := &imageT{p: pixelsT{}, xmin: i.xmin - 1, xmax: i.xmax + 1, ymin: i.ymin - 1, ymax: i.ymax + 1}
	pixelCh := stream.MapStream(ch, func(p keyT) (pixelT, bool) {
		bits := i.p[keyT{p[0] - 1, p[1] - 1}]<<8 |
			i.p[keyT{p[0], p[1] - 1}]<<7 |
			i.p[keyT{p[0] + 1, p[1] - 1}]<<6 |
			i.p[keyT{p[0] - 1, p[1]}]<<5 |
			i.p[keyT{p[0], p[1]}]<<4 |
			i.p[keyT{p[0] + 1, p[1]}]<<3 |
			i.p[keyT{p[0] - 1, p[1] + 1}]<<2 |
			i.p[keyT{p[0], p[1] + 1}]<<1 |
			i.p[keyT{p[0] + 1, p[1] + 1}]
		v := filter[bits]
		return pixelT{p[0], p[1], v}, v == 1
	})

	for p := range pixelCh {
		result.p[keyT{p[0], p[1]}] = 1
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
	img := &imageT{p: pixelsT{}}
	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				img.p[keyT{x, y}] = 1
				if x > img.xmax {
					img.xmax = x
				}
			}
		}
		if y > img.ymax {
			img.ymax = y
		}
	}
	return img
}
