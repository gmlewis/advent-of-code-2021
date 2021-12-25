package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/v1/test"
)

func TestExample1(t *testing.T) {
	want := "Solution: 39\n"
	test.Runner(t, example1, want, process, &printf)
}

func TestExample2(t *testing.T) {
	want := "Solution: 590784\n"
	test.Runner(t, example2, want, process, &printf)
}

func TestExample2Subsets(t *testing.T) {
	lines := strings.Split(strings.TrimSpace(example2), "\n")
	want := []int{
		139590, 210918, 225476, 328328, 387734, 420416, 436132, 478727, 494759, 494804, 492164, 534936, 534936, 567192, 567150, 592167, 588567, 592902, 590029, 590784, 590784, 590784,
	}

	for i := 1; i <= len(lines); i++ {
		subset := strings.Join(lines[0:i], "\n")
		wantStr := fmt.Sprintf("Solution: %v\n", want[i-1])
		test.Runner(t, subset, wantStr, process, &printf)
	}
}

func TestExample3Subsets(t *testing.T) {
	lines := strings.Split(strings.TrimSpace(example3), "\n")
	want := []int{
		151686,
		248314,
		310956,
		389786,
		389786,
		421952,
		421700,
		433638,
		433638,
		474140,
	}

	for i := 1; i <= len(lines); i++ {
		subset := strings.Join(lines[0:i], "\n")
		wantStr := fmt.Sprintf("Solution: %v\n", want[i-1])
		test.Runner(t, subset, wantStr, process, &printf)
	}
}

func TestExample3(t *testing.T) {
	want := "Solution: 474140\n"
	test.Runner(t, example3, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
on x=10..12,y=10..12,z=10..12
on x=11..13,y=11..13,z=11..13
off x=9..11,y=9..11,z=9..11
on x=10..10,y=10..10,z=10..10
`

var example2 = `
on x=-20..26,y=-36..17,z=-47..7
on x=-20..33,y=-21..23,z=-26..28
on x=-22..28,y=-29..23,z=-38..16
on x=-46..7,y=-6..46,z=-50..-1
on x=-49..1,y=-3..46,z=-24..28
on x=2..47,y=-22..22,z=-23..27
on x=-27..23,y=-28..26,z=-21..29
on x=-39..5,y=-6..47,z=-3..44
on x=-30..21,y=-8..43,z=-13..34
on x=-22..26,y=-27..20,z=-29..19
off x=-48..-32,y=26..41,z=-47..-37
on x=-12..35,y=6..50,z=-50..-2
off x=-48..-32,y=-32..-16,z=-15..-5
on x=-18..26,y=-33..15,z=-7..46
off x=-40..-22,y=-38..-28,z=23..41
on x=-16..35,y=-41..10,z=-47..6
off x=-32..-23,y=11..30,z=-14..3
on x=-49..-5,y=-3..45,z=-29..18
off x=18..30,y=-20..-8,z=-3..13
on x=-41..9,y=-7..43,z=-33..15
on x=-54112..-39298,y=-85059..-49293,z=-27449..7877
on x=967..23432,y=45373..81175,z=27513..53682
`

var example3 = `
on x=-5..47,y=-31..22,z=-19..33
on x=-44..5,y=-27..21,z=-14..35
on x=-49..-1,y=-11..42,z=-10..38
on x=-20..34,y=-40..6,z=-44..1
off x=26..39,y=40..50,z=-2..11
on x=-41..5,y=-41..6,z=-36..8
off x=-43..-33,y=-45..-28,z=7..25
on x=-33..15,y=-32..19,z=-34..11
off x=35..47,y=-46..-34,z=-11..5
on x=-14..36,y=-6..44,z=-16..29
`
