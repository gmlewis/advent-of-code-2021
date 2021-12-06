package enum

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/google/go-cmp/cmp"
)

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
			got := Map(tt.from, tt.f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Map(%+v) = %+v, want %+v", tt.from, got, tt.want)
			}
		})
	}
}

func TestScan(t *testing.T) {
	want := []int{1, 3, 6, 10, 15}
	got := Scan(Range(1, 5), 0, func(a, b int) int { return a + b })
	if !cmp.Equal(got, want) {
		t.Errorf("Scan = %+v, want %+v", got, want)
	}
}

func TestUniq(t *testing.T) {
	want := []int{1, 2, 3}
	got := Uniq([]int{1, 2, 3, 3, 2, 1})
	if !cmp.Equal(got, want) {
		t.Errorf("Uniq = %+v, want %+v", got, want)
	}
}

func TestZip(t *testing.T) {
	want := [][]int{{1, 3, 5}, {2, 4, 6}}
	got := Zip([][]int{{1, 2}, {3, 4}, {5, 6}})
	if !cmp.Equal(got, want) {
		t.Errorf("Zip = %+v, want %+v", got, want)
	}
}

func TestZip2(t *testing.T) {
	type ns struct {
		N int
		S string
	}
	want := []ns{{1, "a"}, {2, "b"}, {3, "c"}}
	f := func(n int, s string) ns { return ns{n, s} }
	got := Zip2([]int{1, 2, 3, 4, 5, 6}, []string{"a", "b", "c"}, f)
	if !cmp.Equal(got, want) {
		t.Errorf("Zip2 = %+v, want %+v", got, want)
	}
}
