package algorithm

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDijkstra(t *testing.T) {
	g := graphT{
		"A": {"B": 3, "C": 1},
		"B": {"A": 3, "C": 7, "D": 5, "E": 1},
		"C": {"A": 1, "B": 7, "D": 2},
		"D": {"C": 2, "B": 5, "E": 7},
		"E": {"B": 1, "D": 7},
	}

	got := Dijkstra[string, int](g, "C", nil, math.MaxInt)

	want := map[string]int{
		"A": 1,
		"B": 4,
		"C": 0,
		"D": 2,
		"E": 5,
	}

	if !cmp.Equal(got, want) {
		t.Errorf("Dijkstra = %+v, want %+v", got, want)
	}
}

type graphT map[string]map[string]int

func (g graphT) Distance(from, to string) int {
	return g[from][to]
}

func (g graphT) Each(f func(key string)) {
	for k := range g {
		f(k)
	}
}

func (g graphT) EachNeighbor(from string, f func(from, to string)) {
	for k := range g[from] {
		f(from, k)
	}
}
