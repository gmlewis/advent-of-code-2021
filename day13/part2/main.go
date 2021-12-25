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
	buf := must.ReadFile(filename)
	parts := strings.Split(buf, "\n\n")
	folds := strings.Split(parts[1], "\n")
	paper := Reduce(strings.Split(parts[0], "\n"), paperT{}, func(line string, acc paperT) paperT {
		p := strings.Split(line, ",")
		x := must.Atoi(p[0])
		y := must.Atoi(p[1])
		acc[keyT{x, y}] = true
		return acc
	})

	for _, line := range folds {
		paper = fold(line, paper)
	}
	prettyPrint(paper)

	printf("Solution: %v\n", len(paper))
}

type keyT [2]int
type paperT map[keyT]bool

func fold(cmd string, paper paperT) paperT {
	parts := strings.Split(cmd, "=")
	v := must.Atoi(parts[1])
	if strings.HasSuffix(parts[0], "x") {
		return doFoldX(paper, v)
	}
	return doFoldY(paper, v)
}

func doFoldY(paper paperT, y int) paperT {
	maxy := 2 * y
	ret := paperT{}
	for k := range paper {
		if k[1] > y {
			ret[keyT{k[0], maxy - k[1]}] = true
			continue
		}
		ret[k] = true
	}
	return ret
}

func doFoldX(paper paperT, x int) paperT {
	maxx := 2 * x
	ret := paperT{}
	for k := range paper {
		if k[0] > x {
			ret[keyT{maxx - k[0], k[1]}] = true
			continue
		}
		ret[k] = true
	}
	return ret
}

func prettyPrint(paper paperT) {
	var maxx, maxy int
	for k := range paper {
		if k[0] > maxx {
			maxx = k[0]
		}
		if k[1] > maxy {
			maxy = k[1]
		}
	}

	var lines []string
	for y := 0; y <= maxy; y++ {
		var line string
		for x := 0; x <= maxx; x++ {
			if paper[keyT{x, y}] {
				line += "█"
			} else {
				line += " "
			}
		}
		lines = append(lines, line)
	}
	logf("\n\npaper:\n%v", strings.Join(lines, "\n"))
}
