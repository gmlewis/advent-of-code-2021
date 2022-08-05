package algorithm

import (
	"golang.org/x/exp/constraints"
)

// Number is a number.
type Number interface {
	constraints.Integer | constraints.Unsigned | constraints.Float
}

// Graph represents a graph that can use Dijkstra's algorithm.
type Graph[K comparable, T Number] interface {
	Distance(from, to K) T
	Each(func(key K))
	EachNeighbor(from K, f func(from, to K))
}

// Dijkstra performs Dijkstra's algorithm to find the shortest path
// from source to target. If target is nil, then all distances are
// calculated.
func Dijkstra[K comparable, T Number](g Graph[K, T], source K, target *K, maxT T) map[K]T {
	inQ := map[K]bool{}
	dist := map[K]T{}
	less := func(a, b K) bool {
		va, okA := dist[a]
		vb, okB := dist[b]
		switch {
		case okA && okB:
			return va < vb
		case okA:
			return true
		default:
			return false
		}
	}
	q := NewPriorityQueue(less)
	prev := map[K]K{}

	g.Each(func(k K) {
		dist[k] = maxT
		if k == source {
			dist[k] = 0
		}
		q.Push(k)
		inQ[k] = true
	})

	f := func(u, v K) {
		if !inQ[v] {
			return
		}
		alt := dist[u] + g.Distance(u, v)
		if alt < dist[v] {
			dist[v] = alt
			prev[v] = u
			q.Fix(v)
		}
	}

	for q.Len() > 0 {
		u := q.Pop()
		delete(inQ, u)

		if target != nil && u == *target {
			break
		}

		g.EachNeighbor(u, f)
	}

	return dist
}

// PathTo finds the shortest path from one node to another but instead of returning all
// the distances, it only returns only the nodes (and their distances) along that path.
func PathTo[K comparable, T Number](g Graph[K, T], source K, target K, maxT T) map[K]T {
	distances := Dijkstra[K, T](g, source, &target, maxT)

	breadcrumbs := map[K]T{target: distances[target]}
	for target != source {
		myDist := distances[target]
		g.EachNeighbor(target, func(from, to K) {
			if distances[to] < myDist {
				breadcrumbs[to] = distances[to]
				target = to
				myDist = distances[to] // find the smallest distance
			}
		})
	}
	return breadcrumbs
}

// Max finds the node with the maximum distance and returns its key and distance.
func Max[K comparable, T Number](distances map[K]T, maxT T) (maxKey K, maxDistance T) {
	for k, d := range distances {
		if d >= maxT {
			continue
		}
		if d > maxDistance {
			maxKey = k
			maxDistance = d
		}
	}
	return maxKey, maxDistance
}
