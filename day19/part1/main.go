// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/enum"
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
	scanners[0].identified = true // scanners[0] is defined to be the origin.

	allBeacons := map[keyT]bool{}
	for k := range scanners[0].beacons {
		allBeacons[k] = true
	}

	for i, base := range scanners {
		if !base.identified {
			continue
		}
		for j, other := range scanners {
			if i == j || other.identified {
				continue
			}
			logf("Comparing %v against %v", base.name, other.name)
			fromBase, fromOther := findCommonBeacons(base, other)
			if len(fromBase) == 0 {
				continue
			}
			// logf("\nfromBase=%+v,\nfromOther=%+v", fromBase, fromOther)
			other.calcPosition(base, fromBase, fromOther)

			for k := range other.beacons {
				nk := other.xform.multKeyT(k)
				ws := keyT{nk[0] + other.pos[0], nk[1] + other.pos[1], nk[2] + other.pos[2]}
				// logf("transformed beacon %+v to worldspace %+v", k, ws)
				allBeacons[ws] = true
			}
		}
	}

	printf("Solution: %v\n", len(allBeacons))
}

type keyT [3]int
type beaconMapT map[keyT]map[int][]keyT
type scannerT struct {
	name       string
	identified bool
	pos        keyT
	beacons    beaconMapT

	xform M3
}

func (s *scannerT) calcPosition(base *scannerT, fromBase, fromOther []keyT) {
	for _, xform := range allXForms {
		delta := MapWithIndex(fromBase, func(i int, base keyT) keyT {
			k := xform.multKeyT(fromOther[i])
			return base.sub(k)
		})
		if All(delta[1:], func(k keyT) bool { return k == delta[0] }) {
			logf("%v has delta: %+v", s.name, delta)
			logf("base.pos=%+v", base.pos)
			s.pos = delta[0].sub(base.pos)
			logf("s.pos=%+v", s.pos)
			s.identified = true
			s.xform = xform
			logf("%v is located at %+v with xform: %+v", s.name, s.pos, s.xform)
			return
		}
	}
	log.Fatalf("unable to calculate xform!")
}

func findCommonBeacons(base, other *scannerT) (fromBase, fromOther []keyT) {
	identified := map[keyT]bool{}
	for kb, vb := range base.beacons {
		for ko, vo := range other.beacons {
			if identified[ko] {
				continue
			}
			var common int
			for k := range vb {
				if _, ok := vo[k]; ok {
					common++
				}
			}
			if common >= 11 { // the beacon itself is the 12th commonality
				// logf("found a match between beacon %v and %v: common=%v", kb, ko, common)
				fromBase = append(fromBase, kb)
				fromOther = append(fromOther, ko)
				identified[ko] = true
				break
			}
		}
	}
	return fromBase, fromOther
}

func parseScanner(buf string) *scannerT {
	lines := strings.Split(buf, "\n")
	beacons := Reduce(lines[1:], beaconMapT{}, func(line string, acc beaconMapT) beaconMapT {
		p := strings.Split(line, ",")
		x := must.Atoi(p[0])
		y := must.Atoi(p[1])
		z := must.Atoi(p[2])
		acc[keyT{x, y, z}] = map[int][]keyT{}
		return acc
	})
	// Find the manhattan distances from each beacon to every other beacon.
	for k := range beacons {
		for j := range beacons {
			if j == k {
				continue
			}
			dist := mathfn.Abs(j[0]-k[0]) + mathfn.Abs(j[1]-k[1]) + mathfn.Abs(j[2]-k[2])
			beacons[k][dist] = append(beacons[k][dist], j)
		}
		// logf("\n\n%v: beacon%v: %+v", lines[0], k, beacons[k])
	}
	return &scannerT{name: lines[0], beacons: beacons}
}

type M3 [3]keyT

var xformsPass1 = []M3{
	{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},   // identity, +X
	{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}},  // (Z rot 90) = +Y
	{{-1, 0, 0}, {0, -1, 0}, {0, 0, 1}}, // (Z rot 180) = -X
	{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}},  // (Z rot 270) = -Y
	{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}},  // (Y rot 90) = +Z
	{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}},  // (Y rot 2700) = -Z
}

var xformsPass2 = []M3{
	{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},   // identity, (X rot 0)
	{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}},  // identity, (X rot 90)
	{{1, 0, 0}, {0, -1, 0}, {0, 0, -1}}, // identity, (X rot 180)
	{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}},  // identity, (X rot 270)
}

var allXForms []M3

func init() {
	for _, m1 := range xformsPass1 {
		for _, m2 := range xformsPass2 {
			m := m2.mult(m1)
			allXForms = append(allXForms, m)
		}
	}
}

// mult multiplies two M3 matrices. Order is important.
func (m M3) mult(other M3) M3 {
	oc := M3{other.column(0), other.column(1), other.column(2)}
	return M3{
		{m[0].dot(oc[0]), m[0].dot(oc[1]), m[0].dot(oc[2])},
		{m[1].dot(oc[0]), m[1].dot(oc[1]), m[1].dot(oc[2])},
		{m[2].dot(oc[0]), m[2].dot(oc[1]), m[2].dot(oc[2])},
	}
}

// column returns a column of the matrix.
func (m M3) column(col int) keyT { return keyT{m[0][col], m[1][col], m[2][col]} }

// dot computes the dot product (aka "scalar product" or "inner product")
// of two vectors (keyTs). The dot product is the cosine of the angle
// between two unit vectors.
func (t keyT) dot(other keyT) int { return t[0]*other[0] + t[1]*other[1] + t[2]*other[2] }

// multKeyT multiples a M3 matrix by a keyT.
func (m M3) multKeyT(other keyT) keyT {
	return keyT{
		m[0].dot(other),
		m[1].dot(other),
		m[2].dot(other),
	}
}

// add adds one keyT to another.
func (k keyT) add(other keyT) keyT {
	return keyT{k[0] + other[0], k[1] + other[1], k[2] + other[2]}
}

// sub subtracts one keyT from another.
func (k keyT) sub(other keyT) keyT {
	return keyT{k[0] - other[0], k[1] - other[1], k[2] - other[2]}
}
