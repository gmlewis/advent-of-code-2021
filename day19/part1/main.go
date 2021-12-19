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
	"github.com/gmlewis/advent-of-code-2021/mathfn"
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
	scanners := Map(strings.Split(buf, "\n\n"), parseScanner)
	values := Map(scanners, func(s scannerT) []int { v := maps.Values(s); sort.Ints(v); return v })
	values = Map(values, func(arr []int) []int {
		return MapWithIndex(arr, func(i, v int) int { return v - arr[0] })
	})

	for i := range scanners {
		// logf("scanners[%v]=%+v", i, scanner)
		logf("values[%v] (%v) =%+v", i, len(values[i]), values[i])
		// if i > 0 {
		// 	last := values[i-1]
		// 	diff := MapWithIndex(values[i], func(index, v int) int { return mathfn.Abs(v - last[index]) })
		// 	sort.Ints(diff)
		// 	logf("diff[%v-%v]=%+v", i, i-1, diff)
		// }
	}

	printf("Solution: %v\n", len(scanners))
}

type scannerT map[[3]int]int

func parseScanner(buf string) scannerT {
	return Reduce(strings.Split(buf, "\n")[1:], scannerT{}, func(line string, acc scannerT) scannerT {
		p := strings.Split(line, ",")
		x := must.Atoi(p[0])
		y := must.Atoi(p[1])
		z := must.Atoi(p[2])
		acc[[3]int{x, y, z}] = mathfn.Abs(x) + mathfn.Abs(y) + mathfn.Abs(z)
		return acc
	})
}
