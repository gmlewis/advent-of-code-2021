// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"

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
	buf := must.ReadFile(filename)
	parts := strings.Split(buf, "=")
	xp := strings.Split(strings.TrimSuffix(parts[1], ", y"), "..")
	yp := strings.Split(parts[2], "..")
	xmin := must.Atoi(xp[0])
	xmax := must.Atoi(xp[1])
	ymin := must.Atoi(yp[0])
	ymax := must.Atoi(yp[1])
	logf("target area: x=%v..%v, y=%v..%v", xmin, xmax, ymin, ymax)

	failed := map[keyT]bool{}
	success := map[keyT]bool{}
	maxHeight := math.MinInt
	for steps := 1; steps < 1000; steps++ {
		height, ok := calcHeight(steps, xmin, xmax, ymin, ymax, failed, success)
		if !ok {
			continue
		}
		// logf("steps=%v, height=%v", steps, height)
		if height < maxHeight {
			// logf("breaking at steps=%v", steps)
			break
		}
		maxHeight = height
	}

	printf("Solution: %v\n", len(success))
}

type keyT [2]int

func shoot(steps, xmin, ymin, xmax, ymax int) (xvMin, yvMin, xvMax, yvMax int) {
	xvMin = math.MaxInt
	for xv := 1; true; xv++ {
		x := xv * (xv + 1) / 2
		if steps < xv {
			d := xv - steps
			x -= d * (d + 1) / 2
		}
		if x >= xmin && xv < xvMin {
			xvMin = xv
		}
		if x > xmax {
			break
		}
		xvMax = xv
	}

	sum := steps * (steps - 1) / 2
	yvMin = int(math.Round(float64(ymin+sum) / float64(steps)))
	yvMax = int(math.Round(float64(ymax+sum) / float64(steps)))
	if yvMin > yvMax {
		yvMin, yvMax = yvMax, yvMin
	}
	return xvMin, yvMin, xvMax, yvMax
}

func calcHeight(steps, xmin, xmax, ymin, ymax int, failed, success map[keyT]bool) (int, bool) {
	xv1, yv1, xv2, yv2 := shoot(steps, xmin, ymin, xmax, ymax)

	maxHeight := math.MinInt
	var foundOne bool
	// var bestXV, bestYV int
	for yv := yv1; yv <= yv2; yv++ {
		for xv := xv1; xv <= xv2; xv++ {
			k := keyT{xv, yv}
			if failed[k] || success[k] {
				continue
			}
			height, ok := simulate(xv, yv, xmin, xmax, ymin, ymax)
			if !ok {
				failed[k] = true
				continue
			}
			success[k] = true
			foundOne = true
			if height > maxHeight {
				maxHeight = height
				// bestXV, bestYV = xv, yv
			}
		}
	}
	// if foundOne {
	// 	logf("calcHeight(steps=%v): best initial velocity: (%v,%v) with height %v", steps, bestXV, bestYV, maxHeight)
	// }

	return maxHeight, foundOne
}

// simulate performs the following actions after each step:
// The probe's x position increases by its x velocity.
// The probe's y position increases by its y velocity.
// Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
// Due to gravity, the probe's y velocity decreases by 1.
func simulate(xv, yv, xmin, xmax, ymin, ymax int) (int, bool) {
	// initXV, initYV := xv, yv
	xp, yp := 0, 0
	maxHeight := 0 // starting height
	var steps int
	for xp <= xmax && yp >= ymin {
		xp += xv
		yp += yv
		if xv > 0 {
			xv--
		}
		yv--
		if yp > maxHeight {
			maxHeight = yp
		}
		steps++
		// logf("after step #%v: pos=(%v,%v), vel=(%v,%v)", steps, xp, yp, xv, yv)

		if xp >= xmin && xp <= xmax && yp >= ymin && yp <= ymax {
			// logf("simulate found solution after %v steps with initial velocity (%v,%v)", steps, initXV, initYV)
			return maxHeight, true
		}
	}
	return 0, false
}
