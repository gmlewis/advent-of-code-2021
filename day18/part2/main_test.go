package main

import (
	"testing"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want string
	}{
		{
			name: "part2 example",
			in: []string{
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
			},
			want: "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := Map(tt.in, func(line string) *nodeT { n, _ := parse(line); return n })
			got := sum(nums).String()
			if got != tt.want {
				t.Errorf("sum = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "part2 example",
			in:   "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]",
			want: 3993,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, _ := parse(tt.in)
			got := n.magnitude()
			if got != tt.want {
				t.Errorf("magnitude = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	want := "Solution: 3993\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`
