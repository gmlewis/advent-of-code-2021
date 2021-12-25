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
	fish := Map(strings.Split(strings.TrimSpace(buf), ","), must.Atoi)
	m := Reduce(Range(1, 80), Frequencies(fish), simFish)
	sum := maps.SumValues(m)

	printf("Solution: %v\n", sum)
}

func simFish(days int, fishMap map[int]int) map[int]int {
	newFish := map[int]int{}
	for k, v := range fishMap {
		if k == 0 {
			newFish[8] = v
			newFish[6] += v
			continue
		}
		newFish[k-1] += v
	}

	return newFish
}
