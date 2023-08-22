package algorithm

import (
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/google/go-cmp/cmp"
)

type keyT [2]int
type puzT struct {
	grid map[keyT]bool
}

// puzT implements the ComponentGraph interface.
var _ ComponentGraph[keyT] = &puzT{}

func (p *puzT) Less(k1, k2 keyT) bool {
	if k1[1] == k2[1] {
		return k1[0] < k2[0]
	}
	return k1[1] < k2[1]
}

func (p *puzT) Each(f func(key keyT)) {
	for k := range p.grid {
		f(k)
	}
}

func (p *puzT) EachNeighbor(from keyT, f func(from, to keyT)) {
	fn := func(k keyT) {
		if p.grid[k] {
			f(from, k)
		}
	}
	fn(keyT{from[0] - 1, from[1]})
	fn(keyT{from[0] + 1, from[1]})
	fn(keyT{from[0], from[1] - 1})
	fn(keyT{from[0], from[1] + 1})
}

func TestConnectedComponents_2Dgrid(t *testing.T) {
	t.Parallel()
	grid := `
####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##
`

	puz := &puzT{grid: map[keyT]bool{}}
	lines := strings.Split(strings.TrimSpace(grid), "\n")
	parseLine := func(y int, line string, acc int) int {
		for x, r := range line {
			if r == '#' {
				puz.grid[keyT{x, y}] = true
			}
		}
		return acc
	}
	enum.ReduceWithIndex(lines, 0, parseLine)

	got := ConnectedComponents[keyT](puz)

	f := func(keys ...keyT) map[keyT]struct{} {
		result := map[keyT]struct{}{}
		for _, k := range keys {
			result[k] = struct{}{}
		}
		return result
	}
	want := map[keyT]map[keyT]struct{}{
		keyT{0, 0}:  f(keyT{0, 0}, keyT{1, 0}, keyT{2, 0}, keyT{3, 0}),
		keyT{1, 2}:  f(keyT{1, 2}, keyT{0, 3}, keyT{1, 3}, keyT{2, 3}, keyT{1, 4}),
		keyT{2, 6}:  f(keyT{2, 6}, keyT{2, 7}, keyT{0, 8}, keyT{1, 8}, keyT{2, 8}),
		keyT{0, 10}: f(keyT{0, 10}, keyT{0, 11}, keyT{0, 12}, keyT{0, 13}),
		keyT{0, 15}: f(keyT{0, 15}, keyT{1, 15}, keyT{0, 16}, keyT{1, 16}),
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ConnectedComponents mismatch (-want +got):\n%v", diff)
	}
}
