package algorithm

// ComponentGraph represents a graph that can find connected components.
type ComponentGraph[K comparable] interface {
	Less(k1, k2 K) bool
	Each(f func(key K))
	EachNeighbor(from K, f func(from, to K))
}

type neighborT[K comparable] struct {
	from K
	to   map[K]struct{}
}

func ConnectedComponents[K comparable](g ComponentGraph[K]) map[K]map[K]struct{} {
	pass1 := map[K]map[K]struct{}{}

	ch := make(chan neighborT[K], 100)
	collect := func(key K) {
		r := neighborT[K]{from: key, to: map[K]struct{}{}}
		g.EachNeighbor(key, func(from, to K) { r.to[to] = struct{}{} })
		ch <- r
	}
	go func() {
		g.Each(collect)
		close(ch)
	}()

	// first pass - find all neighbors
	for r := range ch {
		pass1[r.from] = r.to
	}

	// second pass - reduce to minimum labels
	pass2 := map[K]map[K]struct{}{}
	visited := map[K]bool{}
	for k := range pass1 {
		if visited[k] {
			continue
		}
		connectedKeys := map[K]struct{}{}
		minKey := minimize[K](g.Less, k, pass1, visited, connectedKeys)
		pass2[minKey] = connectedKeys
		// log.Printf("pass2[%v] = %v", minKey, connectedKeys)
	}

	return pass2
}

func minimize[K comparable](lessFn func(k1, k2 K) bool, key K, pass1 map[K]map[K]struct{}, visited map[K]bool, connectedKeys map[K]struct{}) K {
	minKey := key
	visited[key] = true
	connectedKeys[key] = struct{}{}
	for k := range pass1[key] {
		if lessFn(k, minKey) {
			minKey = k
		}
		connectedKeys[k] = struct{}{}
		if !visited[k] {
			m := minimize[K](lessFn, k, pass1, visited, connectedKeys)
			if lessFn(m, minKey) {
				minKey = m
			}
		}
	}
	return minKey
}
