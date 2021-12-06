// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

var printf = fmt.Printf

func main() {
	flag.Parse()

	enum.Each(flag.Args(), process)
}

func process(filename string) {
	log.Printf("Processing %v ...", filename)
	buf := must.ReadFile(filename)
	fish := enum.Map(strings.Split(strings.TrimSpace(buf), ","), must.Atoi)
	fish = enum.Reduce(enum.Range(1, 80), fish, simFish)

	printf("Solution: %v\n", len(fish))
}

func simFish(days int, fish []int) []int {
	n := len(fish)
	for i := 0; i < n; i++ {
		if fish[i] == 0 {
			fish = append(fish, 8)
			fish[i] = 6
			continue
		}
		fish[i]--
	}

	// log.Printf("After %v days: %v fish: %+v", days, len(fish), fish)
	return fish
}
