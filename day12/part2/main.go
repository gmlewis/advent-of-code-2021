// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
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

	edges := Reduce(lines, mapT{}, func(line string, acc mapT) mapT {
		p := strings.Split(line, "-")
		if p[1] != "start" {
			acc[p[0]] = append(acc[p[0]], p[1])
		}
		if p[0] != "start" {
			acc[p[1]] = append(acc[p[1]], p[0])
		}
		return acc
	})

	allPaths := findPaths("start", nil, edges)

	printf("Solution: %v\n", len(allPaths))
}

type mapT map[string][]string

type pathT struct {
	visited map[string]int
}

func findPaths(root string, acc []*pathT, edges mapT) []*pathT {
	for _, node := range edges[root] {
		p := &pathT{visited: map[string]int{}}
		acc = append(acc, p.completeAllPaths(node, edges)...)
	}
	return acc
}

func (p *pathT) completeAllPaths(node string, edges mapT) []*pathT {
	if node[0] >= 'a' && p.visited[node] >= 1 {
		if maps.Any(p.visited, func(k string, v int) bool { return k[0] >= 'a' && v >= 2 }) {
			return nil
		}
	}
	if node == "end" {
		return []*pathT{p}
	}

	var ret []*pathT
	for _, n := range edges[node] {
		visited := map[string]int{node: 1}
		for k, v := range p.visited {
			visited[k] += v
		}
		p2 := &pathT{visited: visited}
		ret = append(ret, p2.completeAllPaths(n, edges)...)
	}
	return ret
}
