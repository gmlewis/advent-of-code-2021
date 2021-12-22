// -*- compile-command: "go run main.go ../example1.txt ../example3.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"sort"

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

	extents := func(vals []int, i1, i2, limit int) (int, int) {
		v1 := vals[i1]
		if i1 == i2 {
			return v1, limit
		}
		return v1, vals[i1+1] - 1
	}

	f := func(cmd *cmdT, space spaceT) spaceT {
		xi1 := xIndices[cmd.x1]
		yi1 := yIndices[cmd.y1]
		zi1 := zIndices[cmd.z1]
		xi2 := xIndices[cmd.x2]
		yi2 := yIndices[cmd.y2]
		zi2 := zIndices[cmd.z2]
		logf("cmd=%+v", *cmd)
		logf("on=%v: xi=%v..%v, yi=%v..%v, zi=%v..%v", cmd.on, xi1, xi2, yi1, yi2, zi1, zi2)
		for zi := zi1; zi <= zi2; zi++ {
			for yi := yi1; yi <= yi2; yi++ {
				for xi := xi1; xi <= xi2; xi++ {
					k := keyT{xi, yi, zi}
					x1, x2 := extents(xVals, xi, xi2, cmd.x2)
					y1, y2 := extents(yVals, yi, yi2, cmd.y2)
					z1, z2 := extents(zVals, zi, zi2, cmd.z2)
					if cmd.on {
						c, ok := space[k] // DEBUG ONLY
						if !ok {
							space[k] = &cmdT{x1: x1, x2: x2, y1: y1, y2: y2, z1: z1, z2: z2}
							logf("ADD: space%+v=%+v, size=%v...", k, space[k], space[k].size())
							continue
						}

						var debug bool
						if x1 < c.x1 || x2 > c.x2 ||
							y1 < c.y1 || y2 > c.y2 ||
							z1 < c.z1 || z2 > c.z2 {
							logf("REPLACE: space%+v=%+v, size=%v...", k, c, c.size())
							debug = true
						}
						if x1 < c.x1 {
							c.x1 = x1
						}
						if x2 > c.x2 {
							c.x2 = x2
						}
						if y1 < c.y1 {
							c.y1 = y1
						}
						if y2 > c.y2 {
							c.y2 = y2
						}
						if z1 < c.z1 {
							c.z1 = z1
						}
						if z2 > c.z2 {
							c.z2 = z2
						}
						if debug {
							logf("WITH: space%+v=%+v, size=%v", k, c, c.size())
						}
						continue
					}

					// logf("REMOVE: space%+v", k)
					c, ok := space[k]
					if !ok {
						// logf("nothing to delete at space%+v", k)
						continue
					}
					if (cmd.x1 <= c.x1 && cmd.x2 >= c.x2) ||
						(cmd.y1 <= c.y1 && cmd.y2 >= c.y2) ||
						(cmd.z1 <= c.z1 && cmd.z2 >= c.z2) {
						// logf("deleting space%+v: %+v, size=%v", k, *c, c.size())
						delete(space, k)
						continue
					}
					if c.x1 != c.x2 && c.x1 <= cmd.x2 {
						// logf("trimming X line space%+v BEFORE: %+v", k, *c)
						c.x1 = cmd.x2 + 1
						// logf("trimming X line space%+v AFTER: %+v", k, *c)
						if c.x1 > c.x2 {
							log.Fatalf("c.x1 > c.x2: %+v", *c)
						}
					}
					if c.y1 != c.y2 && c.y1 <= cmd.y2 {
						// logf("trimming Y line space%+v BEFORE: %+v", k, *c)
						c.y1 = cmd.y2 + 1
						// logf("trimming Y line space%+v AFTER: %+v", k, *c)
						if c.y1 > c.y2 {
							log.Fatalf("c.y1 > c.y2: %+v", *c)
						}
					}
					if c.z1 != c.z2 && c.z1 <= cmd.z2 {
						// logf("trimming Z line space%+v BEFORE: %+v", k, *c)
						c.z1 = cmd.z2 + 1
						// logf("trimming Z line space%+v AFTER: %+v", k, *c)
						if c.z1 > c.z2 {
							log.Fatalf("c.z1 > c.z2: %+v", *c)
						}
					}
				}
			}
		}

		return space
	}

	space := Reduce(cmds, spaceT{}, f)
	// var debug []string
	cubesOn := maps.Reduce(space, int64(0), func(k keyT, c *cmdT, acc int64) int64 {
		// debug = append(debug, fmt.Sprintf("SUM: space%+v x=%v..%v,y=%v..%v,z=%v..%v = %v", k, c.x1, c.x2, c.y1, c.y2, c.z1, c.z2, c.size()))
		return acc + c.size()
	})
	// sort.Strings(debug)
	// logf("\n\n%v", strings.Join(debug, "\n"))

	printf("Solution: %v\n", cubesOn)
}

type lookupT map[int]int
type keyT [3]int
type spaceT map[keyT]*cmdT
type cmdT struct {
	on bool
	x1 int
	x2 int
	y1 int
	y2 int
	z1 int
	z2 int
}

func (c *cmdT) size() int64 { // inclusive
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
