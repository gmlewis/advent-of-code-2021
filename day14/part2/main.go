// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/v1/enum"
	"github.com/gmlewis/advent-of-code-2021/v1/maps"
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

	h = Reduce(Range(1, 40), h, func(step int, acc countT) countT {
		return maps.Reduce(acc, countT{}, func(k string, v int, acc2 countT) countT {
			c := rules[k]
			acc2[k[0:1]+c] += v
			acc2[c+k[1:2]] += v
			return acc2
		})
	})

	histo := countT{start[len(start)-1:]: 1}
	for k, v := range h {
		histo[k[0:1]] += v
	}

	values := maps.Values(histo)
	sort.Ints(values)

	printf("Solution: %v\n", values[len(values)-1]-values[0])
}

type mapT map[string]string
type countT map[string]int
