// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/bits"
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

	digits := Reduce(all, str2digit{}, identify147)
	digBits := Reduce(all, str2bits{}, strToBits)
	if All(rhs, maps.HasKey(digits)) {
		return ReduceWithIndex(rhs, 0, calcNum(digits))
	}

	swapDigits := maps.Swap(digits)
	swapDigBits := maps.Swap(digBits)

	mask := byte(0x7f)
	one := digBits[swapDigits[1]]
	four := digBits[swapDigits[4]]
	seven := digBits[swapDigits[7]]
	a := one ^ seven
	ll2 := mask &^ (four | seven)
	e := byte(1 << bits.TrailingZeros(uint(ll2)))
	g := ll2 ^ e
	lm2 := one ^ four
	b := byte(1 << bits.TrailingZeros(uint(lm2)))
	d := lm2 ^ b
	c := byte(1 << bits.TrailingZeros(uint(one)))
	f := one ^ c

	var nine byte
	if _, ok := swapDigBits[four|seven|g]; !ok {
		e, g = g, e
	}
	nine = four | seven | g
	digits[swapDigBits[nine]] = 9
	swapDigits[9] = swapDigBits[nine]

	var three byte
	if _, ok := swapDigBits[nine^b]; !ok {
		b, d = d, b
	}
	three = nine ^ b
	digits[swapDigBits[three]] = 3
	swapDigits[3] = swapDigBits[three]

	var five byte
	if _, ok := swapDigBits[nine^c]; !ok {
		c, f = f, c
	}
	five = nine ^ c
	digits[swapDigBits[five]] = 5
	swapDigits[5] = swapDigBits[five]

	zero := a | b | c | e | f | g
	digits[swapDigBits[zero]] = 0
	swapDigits[0] = swapDigBits[zero]

	two := a | c | d | e | g
	digits[swapDigBits[two]] = 2
	swapDigits[2] = swapDigBits[two]

	six := a | b | d | e | f | g
	digits[swapDigBits[six]] = 6
	swapDigits[6] = swapDigBits[six]

	return ReduceWithIndex(rhs, 0, calcNum(digits))
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
