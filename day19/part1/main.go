// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/v1/enum"
	"github.com/gmlewis/advent-of-code-2021/v1/mathfn"
	"github.com/gmlewis/advent-of-code-2021/v1/must"
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
	for k := range scanners[0].wsBeacons {
		allBeacons[k] = true
	}

	foundScanners := map[int]bool{0: true}
	for len(foundScanners) != len(scanners) {
		for i, base := range scanners {
			if !base.identified {
				continue
			}
			for j, other := range scanners {
				if i == j || other.identified {
					continue
				}
				fromBase, fromOther := findCommonBeacons(base, other)
				if len(fromBase) == 0 {
					continue
				}
				foundScanners[j] = true
				other.calcPosition(base, fromBase, fromOther)

				for k := range other.lsBeacons {
					nk := other.xform.multKeyT(k)
					ws := keyT{nk[0] + other.pos[0], nk[1] + other.pos[1], nk[2] + other.pos[2]}
					allBeacons[ws] = true
				}
			}
		}
	}

	printf("Solution: %v\n", len(allBeacons))
}

type keyT [3]int
type fingerPrintT map[int]struct{}
type beaconMapT map[keyT]fingerPrintT
type scannerT struct {
	name       string
	identified bool
	pos        keyT
	lsBeacons  beaconMapT
	wsBeacons  beaconMapT

	xform M3
}

func (s *scannerT) calcPosition(base *scannerT, fromBase, fromOther []keyT) {
	for _, xform := range allXForms {
		delta := MapWithIndex(fromBase, func(i int, base keyT) keyT {
			k := xform.multKeyT(fromOther[i])
			return base.sub(k)
		})
		if All(delta[1:], func(k keyT) bool { return k == delta[0] }) {
			s.wsBeacons = beaconMapT{}
			s.pos = delta[0]
			s.identified = true
			s.xform = xform
			for k, v := range s.lsBeacons {
				wsk := xform.multKeyT(k).add(s.pos)
				s.wsBeacons[wsk] = v
			}
			return
		}
	}
	log.Fatalf("unable to calculate xform!")
}

func findCommonBeacons(base, other *scannerT) (fromBase, fromOther []keyT) {
	identified := map[keyT]bool{}
	for kb, vb := range base.wsBeacons {
		for ko, vo := range other.lsBeacons {
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
	wsBeacons, lsBeacons := beaconMapT{}, beaconMapT{}
	Each(lines[1:], func(line string) {
		p := strings.Split(line, ",")
		k := keyT{must.Atoi(p[0]), must.Atoi(p[1]), must.Atoi(p[2])}
		lsBeacons[k] = fingerPrintT{}
		wsBeacons[k] = fingerPrintT{}
	})
	// Find the manhattan distances from each beacon to every other beacon.
	for k := range lsBeacons {
		for j := range lsBeacons {
			if j == k {
				continue
			}
			dist := mathfn.Abs(j[0]-k[0]) + mathfn.Abs(j[1]-k[1]) + mathfn.Abs(j[2]-k[2])
			lsBeacons[k][dist] = struct{}{}
			wsBeacons[k][dist] = struct{}{}
		}
	}
	return &scannerT{name: lines[0], lsBeacons: lsBeacons, wsBeacons: wsBeacons}
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
