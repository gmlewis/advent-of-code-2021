// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

	. "github.com/gmlewis/advent-of-code-2021/v1/enum"
	"github.com/gmlewis/advent-of-code-2021/v1/must"
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

	ls := Map(lines, identify)
	incomplete := Filter(ls, func(line *lineT) bool { return line.lineType == incomplete })
	logf("%v incomplete lines", len(incomplete))

	scores := Map(incomplete, func(line *lineT) int { return line.score })
	sort.Ints(scores)
	score := scores[len(scores)/2]

	printf("Solution: %v\n", score)
}

type lineTypeT int

const (
	valid lineTypeT = iota
	incomplete
	corrupt
)

type lineT struct {
	line     string
	lineType lineTypeT
	illegal  rune
	score    int
}

func identify(line string) *lineT {
	lt := &lineT{line: line}
	var stack []rune
	for _, r := range line {
		if c, ok := open2close[r]; ok {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 || r != stack[len(stack)-1] {
			lt.lineType = corrupt
			lt.illegal = r
			lt.score = score[r]
			break
		}
		stack = stack[:len(stack)-1]
	}

	if lt.lineType == valid && len(stack) != 0 {
		lt.lineType = incomplete
		for i := len(stack) - 1; i >= 0; i-- {
			r := stack[i]
			lt.score = 5*lt.score + score[r]
		}
	}

	return lt
}

var open2close = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
	'<': '>',
}

var score = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}
