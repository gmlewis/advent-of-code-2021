package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 0\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `Syntax error in navigation subsystem on line: <span title="Some days, that's just how it is.">all of them</span>

[({(&lt;(())[]&gt;[[{[]{&lt;()&lt;&gt;&gt;
[(()[&lt;&gt;])]({[&lt;{&lt;&lt;[]&gt;&gt;(
{([(&lt;{}[&lt;&gt;[]}&gt;{[]{[(&lt;()&gt;
(((({&lt;&gt;}&lt;{&lt;{&lt;&gt;}{[]{[]{}
[[&lt;[([]))&lt;([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{&lt;[[]]&gt;}&lt;{[{[{[]{()[[[]
[&lt;(&lt;(&lt;(&lt;{}))&gt;&lt;([]([]()
&lt;{([([[(&lt;&gt;()){}]&gt;(&lt;&lt;{{
&lt;{([{{}}[&lt;[[[&lt;&gt;{}]]]&gt;[]]
`
