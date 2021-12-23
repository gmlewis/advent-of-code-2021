package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample1(t *testing.T) {
	want := "Solution: 39\n"
	test.Runner(t, example1, want, process, &printf)
}

func TestExample2Subset(t *testing.T) {
	subset := `
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
`

	lines := strings.Split(strings.TrimSpace(subset), "\n")
	want := []int{
		139590, 210918, 225476, 328328, 387734, 420416, 436132, 478727, 494759, 494804, 492164, 534936, 534936, 567192, 567150, 592167, 588567, 592902, 590029, 590784, 590784, 590784,
	}

	for i := 1; i <= len(lines); i++ {
		subset := strings.Join(lines[0:i], "\n")
		wantStr := fmt.Sprintf("Solution: %v\n", want[i-1])
		test.Runner(t, subset, wantStr, process, &printf)
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name  string
		start string
		add   string
		want  string
	}{
		{
			name:  "simple 2x2x2 - no expansion",
			start: "on x=0..1,y=0..1,z=0..1",
			add:   "on x=0..1,y=0..1,z=0..1",
			want:  "x=0..1,y=0..1,z=0..1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := processLine(tt.start, nil)[0]
			a := processLine(tt.add, nil)[0]
			sc := newCuboid(s.x1, s.x2, 1, s.y1, s.y2, 1, s.z1, s.z2, 1)
			ac := newCuboid(a.x1, a.x2, 1, a.y1, a.y2, 1, a.z1, a.z2, 1)
			got := sc.add(ac).String()
			if got != tt.want {
				t.Errorf("add = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name  string
		start *cuboidT
		sub   *cuboidT
		want  *cuboidT
	}{
		{
			name:  "xyPlane sub zAxis",
			start: &cuboidT{x1: -41, x2: -41, y1: 41, y2: 42, z1: -47, z2: -39, features: yzPlane},
			sub:   &cuboidT{x1: -41, x2: -41, y1: 41, y2: 42, z1: -47, z2: -39, features: zAxis},
			want:  &cuboidT{x1: -41, x2: -41, y1: 41, y2: 42, z1: -47, z2: -39, features: yzPlane},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.start.subtract(tt.sub)
			if got.features != tt.want.features {
				t.Errorf("subtract = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExample3Subset(t *testing.T) {
	subset := `
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

	lines := strings.Split(strings.TrimSpace(subset), "\n")
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

func TestInputSubset(t *testing.T) {
	subset := `
on x=-3..43,y=-40..7,z=-4..40
on x=-20..26,y=-14..40,z=-10..35
on x=-15..35,y=-41..5,z=-24..27
on x=-29..15,y=-6..44,z=-3..42
on x=-41..8,y=-44..0,z=-5..42
on x=-29..24,y=-42..12,z=-38..9
on x=-27..25,y=-29..24,z=-18..35
on x=-38..15,y=-9..44,z=-39..10
on x=-21..29,y=-37..16,z=-26..27
on x=-48..6,y=-7..44,z=-46..4
off x=-30..-13,y=17..31,z=2..17
on x=-28..21,y=-29..25,z=-25..19
off x=-7..3,y=9..25,z=-25..-10
on x=-9..45,y=-11..36,z=-45..2
off x=6..19,y=-39..-29,z=2..11
on x=-8..40,y=-28..22,z=-10..44
off x=-4..11,y=35..47,z=-41..-26
on x=-42..7,y=-16..34,z=-43..4
off x=-49..-33,y=-23..-10,z=-22..-11
on x=-8..44,y=1..46,z=-30..18
`
	want := "Solution: 609563\n"
	test.Runner(t, subset, want, process, &printf)
}

func TestExample3(t *testing.T) {
	want := "Solution: 2758514936282235\n"
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
on x=-57795..-6158,y=29564..72030,z=20435..90618
on x=36731..105352,y=-21140..28532,z=16094..90401
on x=30999..107136,y=-53464..15513,z=8553..71215
on x=13528..83982,y=-99403..-27377,z=-24141..23996
on x=-72682..-12347,y=18159..111354,z=7391..80950
on x=-1060..80757,y=-65301..-20884,z=-103788..-16709
on x=-83015..-9461,y=-72160..-8347,z=-81239..-26856
on x=-52752..22273,y=-49450..9096,z=54442..119054
on x=-29982..40483,y=-108474..-28371,z=-24328..38471
on x=-4958..62750,y=40422..118853,z=-7672..65583
on x=55694..108686,y=-43367..46958,z=-26781..48729
on x=-98497..-18186,y=-63569..3412,z=1232..88485
on x=-726..56291,y=-62629..13224,z=18033..85226
on x=-110886..-34664,y=-81338..-8658,z=8914..63723
on x=-55829..24974,y=-16897..54165,z=-121762..-28058
on x=-65152..-11147,y=22489..91432,z=-58782..1780
on x=-120100..-32970,y=-46592..27473,z=-11695..61039
on x=-18631..37533,y=-124565..-50804,z=-35667..28308
on x=-57817..18248,y=49321..117703,z=5745..55881
on x=14781..98692,y=-1341..70827,z=15753..70151
on x=-34419..55919,y=-19626..40991,z=39015..114138
on x=-60785..11593,y=-56135..2999,z=-95368..-26915
on x=-32178..58085,y=17647..101866,z=-91405..-8878
on x=-53655..12091,y=50097..105568,z=-75335..-4862
on x=-111166..-40997,y=-71714..2688,z=5609..50954
on x=-16602..70118,y=-98693..-44401,z=5197..76897
on x=16383..101554,y=4615..83635,z=-44907..18747
off x=-95822..-15171,y=-19987..48940,z=10804..104439
on x=-89813..-14614,y=16069..88491,z=-3297..45228
on x=41075..99376,y=-20427..49978,z=-52012..13762
on x=-21330..50085,y=-17944..62733,z=-112280..-30197
on x=-16478..35915,y=36008..118594,z=-7885..47086
off x=-98156..-27851,y=-49952..43171,z=-99005..-8456
off x=2032..69770,y=-71013..4824,z=7471..94418
on x=43670..120875,y=-42068..12382,z=-24787..38892
off x=37514..111226,y=-45862..25743,z=-16714..54663
off x=25699..97951,y=-30668..59918,z=-15349..69697
off x=-44271..17935,y=-9516..60759,z=49131..112598
on x=-61695..-5813,y=40978..94975,z=8655..80240
off x=-101086..-9439,y=-7088..67543,z=33935..83858
off x=18020..114017,y=-48931..32606,z=21474..89843
off x=-77139..10506,y=-89994..-18797,z=-80..59318
off x=8476..79288,y=-75520..11602,z=-96624..-24783
on x=-47488..-1262,y=24338..100707,z=16292..72967
off x=-84341..13987,y=2429..92914,z=-90671..-1318
off x=-37810..49457,y=-71013..-7894,z=-105357..-13188
off x=-27365..46395,y=31009..98017,z=15428..76570
off x=-70369..-16548,y=22648..78696,z=-1892..86821
on x=-53470..21291,y=-120233..-33476,z=-44150..38147
off x=-93533..-4276,y=-16170..68771,z=-104985..-24507
`
