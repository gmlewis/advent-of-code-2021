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

	extents := func(vals []int, i1, i2, limit int) (int, int, int) {
		v1 := vals[i1]
		if i1 == i2 {
			if i1+1 < len(vals) {
				return v1, limit, vals[i1+1] - 1
			}
			return v1, limit, limit
		}
		return v1, vals[i1+1] - 1, vals[i1+1] - 1
	}

	f := func(cmd *cmdT, space spaceT) spaceT {
		xi1 := xIndices[cmd.x1]
		yi1 := yIndices[cmd.y1]
		zi1 := zIndices[cmd.z1]
		xi2 := xIndices[cmd.x2]
		yi2 := yIndices[cmd.y2]
		zi2 := zIndices[cmd.z2]
		for zi := zi1; zi <= zi2; zi++ {
			for yi := yi1; yi <= yi2; yi++ {
				for xi := xi1; xi <= xi2; xi++ {
					k := keyT{xi, yi, zi}
					x1, x2, xm := extents(xVals, xi, xi2, cmd.x2)
					y1, y2, ym := extents(yVals, yi, yi2, cmd.y2)
					z1, z2, zm := extents(zVals, zi, zi2, cmd.z2)
					c, ok := space[k]
					if cmd.on {
						if !ok {
							space[k] = newCuboid(x1, x2, xm, y1, y2, ym, z1, z2, zm)
							continue
						}
						space[k] = c.add(newCuboid(x1, x2, xm, y1, y2, ym, z1, z2, zm))
						continue
					}

					if !ok {
						continue
					}
					if x1 <= c.x1 && x2 >= c.x2 &&
						y1 <= c.y1 && y2 >= c.y2 &&
						z1 <= c.z1 && z2 >= c.z2 {
						delete(space, k)
						continue
					}
					space[k].subtract(newCuboid(x1, x2, xm, y1, y2, ym, z1, z2, zm))
					if space[k] == nil {
						delete(space, k)
					}
				}
			}
		}

		return space
	}

	space := Reduce(cmds, spaceT{}, f)
	cubesOn := maps.Reduce(space, int64(0), func(k keyT, c *cuboidT, acc int64) int64 {
		return acc + c.size()
	})

	printf("Solution: %v\n", cubesOn)
}

type lookupT map[int]int
type keyT [3]int
type spaceT map[keyT]*cuboidT
type cmdT struct {
	on bool
	x1 int
	x2 int
	y1 int
	y2 int
	z1 int
	z2 int
}
type cuboidT struct {
	x1       int
	x2       int
	y1       int
	y2       int
	z1       int
	z2       int
	features featuresT
}
type featuresT int

const (
	singleDot featuresT = 1 << iota
	xAxis
	yAxis
	zAxis
	xyPlane
	yzPlane
	xzPlane
	cubeBody
)

const totallyFilled = singleDot | xAxis | yAxis | zAxis | xyPlane | yzPlane | xzPlane | cubeBody

func newCuboid(x1, x2, xm, y1, y2, ym, z1, z2, zm int) *cuboidT {
	c := &cuboidT{x1: x1, x2: xm, y1: y1, y2: ym, z1: z1, z2: zm}
	switch {
	case x1 == x2 && y1 == y2 && z1 == z2:
		c.features = singleDot
		return c
	case x1 < x2 && y1 < y2 && z1 < z2:
		c.features = totallyFilled
		return c
	case x1 < x2 && y1 == y2 && z1 == z2:
		c.features = singleDot | xAxis
		return c
	case x1 == x2 && y1 < y2 && z1 == z2:
		c.features = singleDot | yAxis
		return c
	case x1 == x2 && y1 == y2 && z1 < z2:
		c.features = singleDot | zAxis
		return c
	case x1 < x2 && y1 < y2 && z1 == z2:
		c.features = singleDot | xAxis | yAxis | xyPlane
		return c
	case x1 == x2 && y1 < y2 && z1 < z2:
		c.features = singleDot | yAxis | zAxis | yzPlane
		return c
	case x1 < x2 && y1 == y2 && z1 < z2:
		c.features = singleDot | xAxis | zAxis | xzPlane
		return c
	default:
		log.Fatalf("bad cuboidT: %#v", c)
	}
	return nil
}

func (c *cuboidT) size() int64 { // inclusive
	if c.features&totallyFilled == totallyFilled {
		return int64(c.x2-c.x1+1) * int64(c.y2-c.y1+1) * int64(c.z2-c.z1+1)
	}
	var sum int64
	if c.features&singleDot == singleDot {
		sum += 1
	}
	if c.features&xAxis == xAxis && c.x2 > c.x1 {
		sum += int64(c.x2 - c.x1)
	}
	if c.features&yAxis == yAxis && c.y2 > c.y1 {
		sum += int64(c.y2 - c.y1)
	}
	if c.features&zAxis == zAxis && c.z2 > c.z1 {
		sum += int64(c.z2 - c.z1)
	}
	if c.features&xyPlane == xyPlane && c.x2 > c.x1 && c.y2 > c.y1 {
		sum += int64(c.x2-c.x1) * int64(c.y2-c.y1)
	}
	if c.features&yzPlane == yzPlane && c.y2 > c.y1 && c.z2 > c.z1 {
		sum += int64(c.y2-c.y1) * int64(c.z2-c.z1)
	}
	if c.features&xzPlane == xzPlane && c.x2 > c.x1 && c.z2 > c.z1 {
		sum += int64(c.x2-c.x1) * int64(c.z2-c.z1)
	}
	if c.features&cubeBody == cubeBody && c.x2 > c.x1 && c.y2 > c.y1 && c.z2 > c.z1 {
		sum += int64(c.x2-c.x1) * int64(c.y2-c.y1) * int64(c.z2-c.z1)
	}
	return sum
}

func (c *cuboidT) add(o *cuboidT) *cuboidT {
	c.features |= o.features
	return c
}

func (c *cuboidT) subtract(o *cuboidT) *cuboidT {
	c.features &^= o.features
	return c
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
