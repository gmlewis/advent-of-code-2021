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

	sum := Sum(Map(lines, solveLine))

	printf("Solution: %v\n", sum)
}

func solveLine(line string) int {
	parts := strings.Split(line, " | ")
	if len(parts) != 2 {
		log.Fatalf("bad line: %v", line)
	}
	lhs := Map(strings.Split(parts[0], " "), sortWord)
	rhs := Map(strings.Split(parts[1], " "), sortWord)
	all := append(append([]string{}, lhs...), rhs...)

	digits := Reduce(all, str2digitT{}, identify147)
	str2bits := Reduce(all, str2bitsT{}, strToBits)
	if All(rhs, maps.HasKey(digits)) {
		return ReduceWithIndex(rhs, 0, calcNum(digits))
	}

	digit2str := maps.Swap(digits)
	bits2str := maps.Swap(str2bits)

	mask := byte(0x7f)
	one := str2bits[digit2str[1]]
	four := str2bits[digit2str[4]]
	seven := str2bits[digit2str[7]]
	a := one ^ seven
	ll2 := mask &^ (four | seven)
	e := byte(1 << bits.TrailingZeros(uint(ll2)))
	g := ll2 ^ e
	lm2 := one ^ four
	b := byte(1 << bits.TrailingZeros(uint(lm2)))
	d := lm2 ^ b
	c := byte(1 << bits.TrailingZeros(uint(one)))
	f := one ^ c

	if _, ok := bits2str[four|seven|g]; !ok {
		e, g = g, e
	}
	nine := four | seven | g

	save := func(bits byte, digit int) { digits[bits2str[bits]] = digit }
	save(nine, 9)

	if _, ok := bits2str[nine^b]; !ok {
		b, d = d, b
	}
	three := nine ^ b
	save(three, 3)

	if _, ok := bits2str[nine^c]; !ok {
		c, f = f, c
	}
	five := nine ^ c
	save(five, 5)

	zero := a | b | c | e | f | g
	save(zero, 0)

	two := a | c | d | e | g
	save(two, 2)

	six := a | b | d | e | f | g
	save(six, 6)

	return ReduceWithIndex(rhs, 0, calcNum(digits))
}

type str2digitT map[string]int
type str2bitsT map[string]byte

func calcNum(digits map[string]int) func(index int, w string, acc int) int {
	return func(index int, w string, acc int) int {
		return acc + digits[w]*int(math.Pow10(3-index))
	}
}

func strToBits(w string, acc str2bitsT) str2bitsT {
	var sum byte
	for _, r := range w {
		sum |= (1 << (r - 'a'))
	}
	acc[w] = sum
	return acc
}

func identify147(w string, acc str2digitT) str2digitT {
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

var digitSize = map[int]int{2: 1, 4: 4, 3: 7, 7: 8}
