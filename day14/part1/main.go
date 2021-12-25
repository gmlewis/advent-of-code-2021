// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
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
	buf := must.ReadFile(filename)
	parts := strings.Split(buf, "\n\n")
	start := parts[0]
	rules := Reduce(strings.Split(parts[1], "\n"), mapT{}, func(rule string, acc mapT) mapT {
		p := strings.Split(rule, " -> ")
		acc[p[0]] = p[1]
		return acc
	})

	final := Reduce(Range(1, 10), start, func(step int, acc string) string {
		next := Map(ChunkEvery([]rune(acc), 2, 1), func(in []rune) string {
			v := string(in)
			return v[0:1] + rules[v]
		})
		return strings.Join(next, "") + acc[len(acc)-1:]
	})

	histo := Frequencies([]rune(final))
	values := maps.Values(histo)
	sort.Ints(values)

	printf("Solution: %v\n", values[len(values)-1]-values[0])
}

type mapT map[string]string
