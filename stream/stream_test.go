package stream

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/gmlewis/advent-of-code-2021/strfn"
	"github.com/google/go-cmp/cmp"
)

func TestChunkEvery_Int(t *testing.T) {
	tests := []struct {
		name string
		end  int
		n    int
		step int
		want [][]int
	}{
		{
			name: "1-element channel",
			n:    2,
			step: 1,
			want: nil,
		},
		{
			name: "2-element channel",
			end:  1,
			n:    2,
			step: 1,
			want: [][]int{{0, 1}},
		},
		{
			name: "5-element channel, n=2, step=1",
			end:  5,
			n:    2,
			step: 1,
			want: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}},
		},
		{
			name: "5-element channel, n=3, step=1",
			end:  5,
			n:    3,
			step: 1,
			want: [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
		},
		{
			name: "5-element channel, n=4, step=1",
			end:  5,
			n:    4,
			step: 1,
			want: [][]int{{0, 1, 2, 3}, {1, 2, 3, 4}, {2, 3, 4, 5}},
		},
		{
			name: "5-element channel, n=5, step=1",
			end:  5,
			n:    5,
			step: 1,
			want: [][]int{{0, 1, 2, 3, 4}, {1, 2, 3, 4, 5}},
		},
		{
			name: "5-element channel, n=6, step=1",
			end:  5,
			n:    6,
			step: 1,
			want: [][]int{{0, 1, 2, 3, 4, 5}},
		},
		{
			name: "5-element channel, n=2, step=2",
			end:  5,
			n:    2,
			step: 2,
			want: [][]int{{0, 1}, {2, 3}, {4, 5}},
		},
		{
			name: "5-element channel, n=3, step=2",
			end:  5,
			n:    3,
			step: 2,
			want: [][]int{{0, 1, 2}, {2, 3, 4}},
		},
		{
			name: "5-element channel, n=4, step=2",
			end:  5,
			n:    4,
			step: 2,
			want: [][]int{{0, 1, 2, 3}, {2, 3, 4, 5}},
		},
		{
			name: "5-element channel, n=5, step=2",
			end:  5,
			n:    5,
			step: 2,
			want: [][]int{{0, 1, 2, 3, 4}},
		},
		{
			name: "5-element channel, n=6, step=2",
			end:  5,
			n:    6,
			step: 2,
			want: [][]int{{0, 1, 2, 3, 4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSlice(ChunkEvery(Range(0, tt.end), tt.n, tt.step))
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ChunkEvery = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestFlatMap_IntToString(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		f      func(int) []string
		want   []string
	}{
		{
			name:   "empty int to string",
			values: []int{},
			want:   []string{},
		},
		{
			name:   "doc example",
			values: []int{1, 2, 3},
			f:      func(v int) []string { s := fmt.Sprintf("%v", v); return []string{s, s} },
			want:   []string{"1", "1", "2", "2", "3", "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlatMap(ToChan(tt.values), tt.f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FlatMap(%+v) = %+v, want %+v", tt.values, got, tt.want)
			}
		})
	}
}

func TestGroupBy_StringToInt(t *testing.T) {
	tests := []struct {
		name      string
		values    []string
		keyFunc   func(string) int
		valueFunc func(string) string
		want      map[int][]string
	}{
		{
			name:   "empty int to string",
			values: []string{},
			want:   map[int][]string{},
		},
		{
			name:      "doc example",
			values:    []string{"ant", "buffalo", "cat", "dingo"},
			keyFunc:   strfn.Length,
			valueFunc: enum.Identity[string],
			want: map[int][]string{
				3: {"ant", "cat"}, 5: {"dingo"}, 7: {"buffalo"},
			},
		},
		{
			name:      "doc example",
			values:    []string{"ant", "buffalo", "cat", "dingo"},
			keyFunc:   strfn.Length,
			valueFunc: strfn.First,
			want: map[int][]string{
				3: {"a", "c"}, 5: {"d"}, 7: {"b"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupBy(ToChan(tt.values), tt.keyFunc, tt.valueFunc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GroupBy(%+v) = %+v, want %+v", tt.values, got, tt.want)
			}
		})
	}
}

func TestMap_StringToInt(t *testing.T) {
	tests := []struct {
		name string
		from []string
		f    func(string) int
		want []int
	}{
		{
			name: "empty string to int",
			from: []string{},
			want: []int{},
		},
		{
			name: "simple decimal string to int",
			from: []string{"1", "2", "3", "100"},
			f:    must.Atoi,
			want: []int{1, 2, 3, 100},
		},
		{
			name: "simple binary string to int",
			from: []string{"1", "10", "11", "10000000"},
			f:    func(s string) int { return must.ParseInt(s, 2, 64) },
			want: []int{1, 2, 3, 128},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(ToChan(tt.from), tt.f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Map(%+v) = %+v, want %+v", tt.from, got, tt.want)
			}
		})
	}
}

func TestRepeatedly(t *testing.T) {
	rand.Seed(0)
	want := []int32{2029793274, 526058514, 1408655353, 116702506, 789387515,
		621654496, 413258767, 1407315077, 1926657288, 359390928}
	got := Take(Repeatedly(rand.Int31), 10)
	if !cmp.Equal(got, want) {
		t.Errorf("Repeatedly = %+v, want %+v", got, want)
	}
}

func TestScan(t *testing.T) {
	want := []int{1, 3, 6, 10, 15}
	got := ToSlice(Scan(Range(1, 5), 0, func(a, b int) int { return a + b }))
	if !cmp.Equal(got, want) {
		t.Errorf("Scan = %+v, want %+v", got, want)
	}
}

func TestUniq(t *testing.T) {
	want := []int{1, 2, 3}
	got := ToSlice(Uniq(ToChan([]int{1, 2, 3, 3, 2, 1})))
	if !cmp.Equal(got, want) {
		t.Errorf("Uniq = %+v, want %+v", got, want)
	}
}

func TestZip2(t *testing.T) {
	type ns struct {
		N int
		S string
	}
	want := []ns{{1, "a"}, {2, "b"}, {3, "c"}}
	f := func(n int, s string) ns { return ns{n, s} }
	got := ToSlice(Zip2(ToChan([]int{1, 2, 3, 4, 5, 6}), ToChan([]string{"a", "b", "c"}), f))
	if !cmp.Equal(got, want) {
		t.Errorf("Zip2 = %+v, want %+v", got, want)
	}
}
