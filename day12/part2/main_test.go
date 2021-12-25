package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/v1/test"
)

func TestExample1(t *testing.T) {
	want := "Solution: 36\n"
	test.Runner(t, example1, want, process, &printf)
}

func TestExample2(t *testing.T) {
	want := "Solution: 103\n"
	test.Runner(t, example2, want, process, &printf)
}

func TestExample3(t *testing.T) {
	want := "Solution: 3509\n"
	test.Runner(t, example3, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

var example2 = `
dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`

var example3 = `
fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`
