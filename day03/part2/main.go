// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/gmlewis/advent-of-code-2021/strfn"
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

	oxygen := filterLines(lines, 0, Longer[string])
	co2 := filterLines(lines, 0, Shorter[string])

	oxygenRating := must.ParseInt(oxygen, 2, 64)
	co2Rating := must.ParseInt(co2, 2, 64)

	printf("Solution: oxygen=%v, co2=%v, product=%v\n", oxygenRating, co2Rating, oxygenRating*co2Rating)
}

func filterLines(lines []string, bit int, f func(a, b []string) []string) string {
	if len(lines) == 1 {
		return lines[0]
	}

	ones := FilterMap(lines, strfn.Substr(bit, bit+1), Equals("1"))
	zeros := FilterMap(lines, strfn.Substr(bit, bit+1), Equals("0"))

	return filterLines(f(ones, zeros), bit+1, f)
}
