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
			logf("\nfromBase=%+v,\nfromOther=%+v", fromBase, fromOther)
			// other.identified = true
			other.calcPosition(fromBase, fromOther)
		}
	}

	printf("Solution: %v\n", len(scanners))
}

type keyT [3]int
type beaconMapT map[keyT]map[int][]keyT
type scannerT struct {
	name       string
	identified bool
	pos        keyT
	beacons    beaconMapT
}

func (s *scannerT) calcPosition(fromBase, fromOther []keyT) {
	delta := MapWithIndex(fromBase, func(i int, base keyT) keyT {
		other := fromOther[i]
		return keyT{base[0] - other[0], base[1] - other[1], base[2] - other[2]}
	})
	sameCount := Reduce(delta, keyT{}, func(k, acc keyT) keyT {
		if k[0] == delta[0][0] {
			acc[0]++
		}
		if k[1] == delta[0][1] {
			acc[1]++
		}
		if k[2] == delta[0][2] {
			acc[2]++
		}
		return acc
	})
	logf("delta: %+v", delta)
	if sameCount[0] == len(delta) {
		s.pos[0] = delta[0][0]
		logf("calcPosition: same X delta value: %v", s.pos[0])
	}
	if sameCount[1] == len(delta) {
		s.pos[1] = delta[0][1]
		logf("calcPosition: same Y delta value: %v", s.pos[1])
	}
	if sameCount[2] == len(delta) {
		s.pos[2] = delta[0][2]
		logf("calcPosition: same Z delta value: %v", s.pos[2])
	}
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
				logf("found a match between beacon %v and %v: common=%v", kb, ko, common)
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

// first beacon:
// --- scanner 0 ---: beacon[-618 -824 -621]: map[97:[[-661 -816 -575]] 245:[[-537 -823 -458]] 1329:[[390 -675 -793]] 1538:[[404 -588 -901]] 1568:[[-485 -357 347]] 1605:[[-447 -329 318]] 1628:[[544 -627 -890]] 1788:[[-345 -311 381]] 1790:[[-584 868 -557]] 1831:[[-689 845 -530]] 1965:[[-789 900 -551]] 1966:[[7 -33 -71]] 2216:[[459 -707 401]] 2219:[[423 -701 434]] 2357:[[528 -643 409]] 2394:[[553 345 -567]] 2542:[[564 392 -477]] 2633:[[630 319 -379]] 2927:[[-892 524 684]] 2990:[[-838 591 734]] 3115:[[-876 649 763]] 3748:[[443 580 662]] 3784:[[474 580 667]] 3975:[[455 729 728]]]
// matches second beacon:
// --- scanner 1 ---: beacon[686 422 578]: map[97:[[729 430 532]] 245:[[605 423 415]] 863:[[669 -402 600]] 978:[[586 -435 557]] 1051:[[567 -361 727]] 1329:[[-322 571 750]] 1431:[[95 138 22]] 1538:[[-336 658 858]] 1568:[[553 889 -390]] 1605:[[515 917 -361]] 1628:[[-476 619 847]] 1788:[[413 935 -424]] 2037:[[703 -491 -529]] 2042:[[755 -354 -619]] 2133:[[-429 -592 574]] 2179:[[-328 -685 520]] 2216:[[-391 539 -444]] 2219:[[-355 545 -477]] 2331:[[807 -499 -711]] 2357:[[-460 603 -452]] 2413:[[-500 -761 534]] 3441:[[-340 -569 -846]] 3629:[[-466 -666 -811]] 3706:[[-364 -763 -893]]]
