// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/maps"
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

	count := Sum(Map(lines, solveLine))

	printf("Solution: %v\n", count)
}

func solveLine(line string) int {
	parts := strings.Split(line, " | ")
	if len(parts) != 2 {
		log.Fatalf("bad line: %v", line)
	}
	lhs := Map(strings.Split(parts[0], " "), sortWord)
	rhs := Map(strings.Split(parts[1], " "), sortWord)
	all := append(append([]string{}, lhs...), rhs...)
	logf("lhs=%v, rhs=%v", lhs, rhs)

	digits := Reduce(all, str2digit{}, identify147)
	digBits := Reduce(all, str2bits{}, strToBits)
	if All(rhs, maps.HasKey(digits)) {
		v := ReduceWithIndex(rhs, 0, calcNum(digits))
		logf("digits=%+v, value=%v", digits, v)
		return v
	}
	logf("digits=%+v, digBits=%+v", digits, digBits)

	return 0
}

type str2digit map[string]int
type str2bits map[string]byte

func calcNum(digits map[string]int) func(index int, w string, acc int) int {
	return func(index int, w string, acc int) int {
		return acc + digits[w]*int(math.Pow10(3-index))
	}
}

func strToBits(w string, acc str2bits) str2bits {
	var sum byte
	for _, r := range w {
		sum |= (1 << (r - 'a'))
	}
	acc[w] = sum
	return acc
}

func identify147(w string, acc str2digit) str2digit {
	if v, ok := digitSize[len(w)]; ok {
		acc[w] = v
	}
	return acc
}

func sortWord(w string) string {
	r := []rune(w)
	sort.Slice(r, func(a, b int) bool { return r[a] < r[b] })
	return string(r)
}

var digitSize = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

var patterns = str2digit{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}
