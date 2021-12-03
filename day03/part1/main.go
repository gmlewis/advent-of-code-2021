// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"log"

	"github.com/gmlewis/advent-of-code-2021/must"
)

func main() {
	flag.Parse()

	must.Process(process)
}

func process(filename string) {
	lines := must.ReadFileLines(filename)
	numBits := len(lines[0])

	var oxygenRating, co2Rating int

	moreInput := lines[:]
	fewerInput := lines[:]
	for i := 0; i < numBits; i++ {
		moreFilter, fewerFilter := genFilters(i, moreInput, fewerInput)
		moreInput = must.FilterStrings(moreInput, moreFilter)
		fewerInput = must.FilterStrings(fewerInput, fewerFilter)
		if len(moreInput) == 1 {
			oxygenRating = must.ParseInt(moreInput[0], 2, 64)
		}
		if len(fewerInput) == 1 {
			co2Rating = must.ParseInt(fewerInput[0], 2, 64)
		}
	}

	log.Printf("oxygen=%v, co2=%v, product=%v", oxygenRating, co2Rating, oxygenRating*co2Rating)
}

func genFilters(bit int, moreInput, fewerInput []string) (moreFilter, fewerFilter must.FilterFunc) {
	var moreCount int
	for _, line := range moreInput {
		if line[bit] == '1' {
			moreCount++
		}
	}

	if 2*moreCount >= len(moreInput) {
		moreFilter = func(line string) bool { return line[bit] == '1' }
	} else {
		moreFilter = func(line string) bool { return line[bit] == '0' }
	}

	var fewerCount int
	for _, line := range fewerInput {
		if line[bit] == '1' {
			fewerCount++
		}
	}

	if 2*fewerCount >= len(fewerInput) {
		fewerFilter = func(line string) bool { return line[bit] == '0' }
	} else {
		fewerFilter = func(line string) bool { return line[bit] == '1' }
	}

	return moreFilter, fewerFilter
}
