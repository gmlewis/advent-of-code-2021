// -*- compile-command: "go run main.go ../example1.txt ../example3.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
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
		logf("cmd=%v", cmd)
		logf("on=%v: xi=%v..%v, yi=%v..%v, zi=%v..%v", cmd.on, xi1, xi2, yi1, yi2, zi1, zi2)
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
						if x1 >= c.x1 && x2 <= c.x2 && y1 >= c.y1 && y2 <= c.y2 && z1 >= c.z1 && z2 <= c.z2 {
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
						logf("deleting space%+v: %v, size=%v", k, c, c.size())
						delete(space, k)
						continue
					}
					space[k].subtract(newCuboid(x1, x2, xm, y1, y2, ym, z1, z2, zm))
				}
			}
		}

		return space
	}

	space := Reduce(cmds, spaceT{}, f)
	// var debug []string
	cubesOn := maps.Reduce(space, int64(0), func(k keyT, c *cuboidT, acc int64) int64 {
		// debug = append(debug, fmt.Sprintf("SUM: space%+v x=%v..%v,y=%v..%v,z=%v..%v = %v", k, c.x1, c.x2, c.y1, c.y2, c.z1, c.z2, c.size()))
		return acc + c.size()
	})
	// sort.Strings(debug)
	// logf("\n\n%v", strings.Join(debug, "\n"))

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

func (t featuresT) String() string {
	if t&totallyFilled == totallyFilled {
		return "totallyFilled"
	}
	var features []string
	if t&singleDot == singleDot {
		features = append(features, "singleDot")
	}
	if t&xAxis == xAxis {
		features = append(features, "xAxis")
	}
	if t&yAxis == yAxis {
		features = append(features, "yAxis")
	}
	if t&zAxis == zAxis {
		features = append(features, "zAxis")
	}
	if t&xyPlane == xyPlane {
		features = append(features, "xyPlane")
	}
	if t&yzPlane == yzPlane {
		features = append(features, "yzPlane")
	}
	if t&xzPlane == xzPlane {
		features = append(features, "xzPlane")
	}
	if t&cubeBody == cubeBody {
		features = append(features, "cubeBody")
	}
	return strings.Join(features, "|")
}

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
		c.features = xAxis
		return c
	case x1 == x2 && y1 < y2 && z1 == z2:
		c.features = yAxis
		return c
	case x1 == x2 && y1 == y2 && z1 < z2:
		c.features = zAxis
		return c
	case x1 < x2 && y1 < y2 && z1 == z2:
		c.features = xyPlane
		return c
	case x1 == x2 && y1 < y2 && z1 < z2:
		c.features = yzPlane
		return c
	case x1 < x2 && y1 == y2 && z1 < z2:
		c.features = xzPlane
		return c
	default:
		log.Fatalf("bad cuboidT: %#v", c)
	}
	return nil
}

func (c *cmdT) String() string {
	state := "off"
	if c.on {
		state = "on"
	}
	return fmt.Sprintf("%v: x=%v..%v,y=%v..%v,z=%v..%v", state, c.x1, c.x2, c.y1, c.y2, c.z1, c.z2)
}

func (c *cuboidT) String() string {
	return fmt.Sprintf("x=%v..%v,y=%v..%v,z=%v..%v; features=%v", c.x1, c.x2, c.y1, c.y2, c.z1, c.z2, c.features)
}

func (c *cuboidT) size() int64 { // inclusive
	return int64(c.x2-c.x1+1) * int64(c.y2-c.y1+1) * int64(c.z2-c.z1+1)
}

func (c *cuboidT) add(o *cuboidT) *cuboidT {
	logf(`start: "on: %v",  // %v`, c, c.size())
	logf(`add: "on: %v",  // %v`, o, o.size())
	before := c.size()
	c.features |= o.features
	logf("want: %q, // %v", c, c.size())
	if c.size() <= before {
		log.Fatalf("add: before=%v, after=%v", before, c.size())
	}
	return c
}

func (c *cuboidT) subtract(o *cuboidT) *cuboidT {
	logf(`start: "on: %v",  // %v`, c, c.size())
	logf(`sub: "off: %v",  // %v`, o, o.size())
	before := c.size()
	c.features ^= o.features
	logf("want: %q, // %v", c, c.size())
	if c.size() >= before {
		log.Fatalf("subtract: before=%v, after=%v", before, c.size())
	}
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
