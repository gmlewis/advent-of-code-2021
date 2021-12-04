// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"log"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/gmlewis/advent-of-code-2021/strfn"
)

func main() {
	flag.Parse()

	enum.Each(flag.Args(), process)
}

func process(filename string) {
	log.Printf("Processing %v ...", filename)
	lines := must.ReadFileLines(filename)

	oxygen := filterLines(lines, 0, enum.Longer[string])
	co2 := filterLines(lines, 0, enum.Shorter[string])

	oxygenRating := must.ParseInt(oxygen, 2, 64)
	co2Rating := must.ParseInt(co2, 2, 64)

	log.Printf("oxygen=%v, co2=%v, product=%v", oxygenRating, co2Rating, oxygenRating*co2Rating)
}

func filterLines(lines []string, bit int, f func(a, b []string) []string) string {
	if len(lines) == 1 {
		return lines[0]
	}

	ones := enum.FilterMap(lines, strfn.Substr(bit, bit+1), strfn.Equals("1"))
	zeros := enum.FilterMap(lines, strfn.Substr(bit, bit+1), strfn.Equals("0"))

	return filterLines(f(ones, zeros), bit+1, f)
}
