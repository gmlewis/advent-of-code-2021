package stream

import (
	"strconv"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/google/go-cmp/cmp"
)

func TestFlatMap_IntToString(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  []string
	}{
		{
			name: "no items",
			want: []string{},
		},
		{
			name:  "one item",
			items: []int{0},
			want:  []string{"0", "0"},
		},
		{
			name:  "two items",
			items: []int{0, 1},
			want:  []string{"0", "0", "1", "1"},
		},
		{
			name:  "three items",
			items: []int{0, 1, 2},
			want:  []string{"0", "0", "1", "1", "2", "2"},
		},
	}

	f := func(v int) []string {
		vs := strconv.Itoa(v)
		return []string{vs, vs}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlatMap(ToChan(tt.items), f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FlatMap = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMap_StringToInt(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  []int
	}{
		{
			name: "no items",
			want: []int{},
		},
		{
			name:  "one item",
			items: []string{"0"},
			want:  []int{0, 0},
		},
		{
			name:  "two items",
			items: []string{"0", "1"},
			want:  []int{0, 0, 1, 1},
		},
		{
			name:  "three items",
			items: []string{"0", "1", "2"},
			want:  []int{0, 0, 1, 1, 2, 2},
		},
	}

	f := func(vs string) []int {
		v := must.Atoi(vs)
		return []int{v, v}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlatMap(ToChan(tt.items), f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FlatMap = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMapWithIndex_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  []string
	}{
		{
			name: "no items",
			want: []string{},
		},
		{
			name:  "one item",
			items: []int{0},
			want:  []string{"0", "0"},
		},
		{
			name:  "two items",
			items: []int{0, 1},
			want:  []string{"0", "0", "1", "1"},
		},
		{
			name:  "three items",
			items: []int{0, 1, 2},
			want:  []string{"0", "0", "1", "1", "2", "2"},
		},
	}

	f := func(i, v int) []string {
		vs := strconv.Itoa(v)
		return []string{vs, vs}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlatMapWithIndex(ToChan(tt.items), f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FlatMapWithIndex = %v, want %v", got, tt.want)
			}
		})
	}
}
