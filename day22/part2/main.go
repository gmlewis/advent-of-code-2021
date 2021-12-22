// -*- compile-command: "go run main.go ../example1.txt ../example3.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"sort"

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
	lines := must.ReadFileLines(filename)
	cmds := Reduce(lines, []*cmdT{}, processLine)
	xVals := Reduce(cmds, []int{}, func(c *cmdT, acc []int) []int { return append(acc, c.x1, c.x2) })
	yVals := Reduce(cmds, []int{}, func(c *cmdT, acc []int) []int { return append(acc, c.y1, c.y2) })
	zVals := Reduce(cmds, []int{}, func(c *cmdT, acc []int) []int { return append(acc, c.z1, c.z2) })
	sort.Ints(xVals)
	xVals = Dedup(xVals)
	sort.Ints(yVals)
	yVals = Dedup(yVals)
	sort.Ints(zVals)
	zVals = Dedup(zVals)
	xIndices := ReduceWithIndex(xVals, lookupT{}, func(i, x int, acc lookupT) lookupT { acc[x] = i; return acc })
	yIndices := ReduceWithIndex(yVals, lookupT{}, func(i, y int, acc lookupT) lookupT { acc[y] = i; return acc })
	zIndices := ReduceWithIndex(zVals, lookupT{}, func(i, z int, acc lookupT) lookupT { acc[z] = i; return acc })
	logf("xVals=%+v", xVals)
	logf("xIndices=%+v", xIndices)
	logf("yVals=%+v", yVals)
	logf("yIndices=%+v", yIndices)
	logf("zVals=%+v", zVals)
	logf("zIndices=%+v", zIndices)

	printf("Solution: %v\n", len(cmds))
}

type lookupT map[int]int
type keyT [3]int
type cmdT struct {
	on bool
	x1 int
	x2 int
	y1 int
	y2 int
	z1 int
	z2 int
}

func (c *cmdT) size() int64 {
	return int64(c.x2-c.x1+1) * int64(c.y2-c.y1+1) * int64(c.z2-c.z1+1)
}

var lineRE = regexp.MustCompile(`^(\S+) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$`)

func processLine(line string, acc []*cmdT) []*cmdT {
	m := lineRE.FindStringSubmatch(line)
	if len(m) != 8 {
		log.Fatalf("unable to parse line: %v", line)
	}
	cmd := &cmdT{
		on: m[1] == "on",
		x1: must.Atoi(m[2]),
		x2: must.Atoi(m[3]),
		y1: must.Atoi(m[4]),
		y2: must.Atoi(m[5]),
		z1: must.Atoi(m[6]),
		z2: must.Atoi(m[7]),
	}
	return append(acc, cmd)
}
