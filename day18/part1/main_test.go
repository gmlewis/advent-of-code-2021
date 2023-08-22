package main

import (
	"testing"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		plus string
		want string
	}{
		{
			name: "first example",
			in:   "[1,2]",
			plus: "[[3,4],5]",
			want: "[[1,2],[[3,4],5]]",
		},
		{
			name: "second example",
			in:   "[[[[4,3],4],4],[7,[[8,4],9]]]",
			plus: "[1,1]",
			want: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, _ := parse(tt.in)
			plus, _ := parse(tt.plus)
			got := n.add(plus).String()
			if got != tt.want {
				t.Errorf("add = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExplode(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "first example",
			in:   "[[[[[9,8],1],2],3],4]",
			want: "[[[[0,9],2],3],4]", // (the 9 has no regular number to its left, so it is not added to any regular number).
		},
		{
			name: "second example",
			in:   "[7,[6,[5,[4,[3,2]]]]]",
			want: "[7,[6,[5,[7,0]]]]", // (the 2 has no regular number to its right, and so it is not added to any regular number).
		},
		{
			name: "third example",
			in:   "[[6,[5,[4,[3,2]]]],1]",
			want: "[[6,[5,[7,0]]],3]",
		},
		{
			name: "fourth example",
			in:   "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			want: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", // (the pair [3,2] is unaffected because the pair [7,3] is further to the left; [3,2] would explode on the next action).
		},
		{
			name: "fifth example",
			in:   "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			want: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			name: "sixth example",
			in:   "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			want: "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
		},
		{
			name: "seventh example",
			in:   "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
			want: "[[[[0,7],4],[15,[0,13]]],[1,1]]",
		},
		{
			name: "eighth example",
			in:   "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			name: "found broken example part 1",
			in:   "[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]",
			want: "[[[[0,[3,2]],[3,3]],[4,4]],[5,5]]",
		},
		{
			name: "found broken example part 2",
			in:   "[[[[0,[3,2]],[3,3]],[4,4]],[5,5]]",
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, _ := parse(tt.in)
			out, addLeft, addRight := n.explode(0)
			if addLeft == nil && addRight == nil {
				t.Error("explode = nil, nil, want left or right explosion")
			}
			got := out.String()
			if got != tt.want {
				t.Errorf("explode = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "simple example 1",
			in:   "10",
			want: "[5,5]",
		},
		{
			name: "simple example 2",
			in:   "11",
			want: "[5,6]",
		},
		{
			name: "simple example 3",
			in:   "12",
			want: "[6,6]",
		},
		{
			name: "first example",
			in:   "[[[[0,7],4],[15,[0,13]]],[1,1]]",
			want: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		},
		{
			name: "second example",
			in:   "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			want: "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, _ := parse(tt.in)
			out, ok := n.split()
			if !ok {
				t.Error("split = false, want true")
			}
			got := out.String()
			if got != tt.want {
				t.Errorf("split = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "first example",
			in:   "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			name: "found broken example",
			in:   "[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]",
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, _ := parse(tt.in)
			got := n.reduce().String()
			if got != tt.want {
				t.Errorf("reduce = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   []string
		want string
	}{
		{
			name: "first example",
			in:   []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"},
			want: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			name: "second example",
			in:   []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"},
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			name: "third example",
			in:   []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"},
			want: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			name: "fourth example",
			in: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			},
			want: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			name: "5th example",
			in: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			},
			want: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
		{
			name: "6th example",
			in: []string{
				"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			},
			want: "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		},
		{
			name: "7th example",
			in: []string{
				"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
			},
			want: "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
		},
		{
			name: "8th example",
			in: []string{
				"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
			},
			want: "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
		},
		{
			name: "9th example",
			in: []string{
				"[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
			},
			want: "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
		},
		{
			name: "10th example",
			in: []string{
				"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
				"[2,9]",
			},
			want: "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		},
		{
			name: "11th example",
			in: []string{
				"[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
			},
			want: "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
		},
		{
			name: "12th example",
			in: []string{
				"[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
				"[[[5,[7,4]],7],1]",
			},
			want: "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
		},
		{
			name: "13th example",
			in: []string{
				"[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
				"[[[[4,2],2],6],[8,7]]",
			},
			want: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			name: "main example",
			in: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			want: "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
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
	t.Parallel()
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "first example",
			in:   "[9,1]",
			want: 29,
		},
		{
			name: "second example",
			in:   "[1,9]",
			want: 21,
		},
		{
			name: "third example",
			in:   "[[9,1],[1,9]]",
			want: 129,
		},
		{
			name: "4th example",
			in:   "[[1,2],[[3,4],5]]",
			want: 143,
		},
		{
			name: "5th example",
			in:   "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			want: 1384,
		},
		{
			name: "6th example",
			in:   "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			want: 445,
		},
		{
			name: "7th example",
			in:   "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			want: 791,
		},
		{
			name: "8th example",
			in:   "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			want: 1137,
		},
		{
			name: "9th example",
			in:   "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			want: 3488,
		},
		{
			name: "main example",
			in:   "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
			want: 4140,
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
	want := "Solution: 4140\n"
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
