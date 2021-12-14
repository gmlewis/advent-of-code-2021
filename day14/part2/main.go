// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
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
	start := parts[0]
	rules := Reduce(strings.Split(parts[1], "\n"), mapT{}, func(rule string, acc mapT) mapT {
		p := strings.Split(rule, " -> ")
		acc[p[0]] = p[1]
		return acc
	})

	h := Reduce(ChunkEvery([]rune(start), 2, 1), countT{}, func(in []rune, acc countT) countT {
		acc[string(in)]++
		return acc
	})

	for i := 1; i <= 40; i++ {
		nh := countT{}
		for k, v := range h {
			c := rules[k]
			el1 := k[0:1] + c
			el2 := c + k[1:2]
			nh[el1] += v
			nh[el2] += v
		}
		h = nh
	}

	histo := countT{}
	for k, v := range h {
		histo[k[0:1]] += v
	}
	histo[start[len(start)-1:]]++

	values := make([]int, 0, len(histo))
	for _, v := range histo {
		values = append(values, v)
	}
	sort.Ints(values)

	printf("Solution: %v\n", values[len(values)-1]-values[0])
}

type mapT map[string]string
type countT map[string]int
