package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFilter_Int(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []int
		f     func(int) bool
		want  []int
	}{
		{
			name: "no items",
			want: nil,
		},
		{
			name:  "one item matches",
			items: []int{0},
			f:     func(v int) bool { return v >= 0 },
			want:  []int{0},
		},
		{
			name:  "no items match",
			items: []int{-1, -2, -3},
			f:     func(v int) bool { return v >= 0 },
			want:  nil,
		},
		{
			name:  "all items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "one item does not match at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "one item does not match at end",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1},
			f:     func(v int) bool { return v >= 0 },
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "one item does not match at middle",
			items: []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.items, tt.f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Filter = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []string
		f     func(string) bool
		want  []string
	}{
		{
			name: "no items",
			want: nil,
		},
		{
			name:  "one item matches",
			items: []string{"yo"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  []string{"yo"},
		},
		{
			name:  "no items match",
			items: []string{"a", "b", "c"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  nil,
		},
		{
			name:  "all items match",
			items: []string{"yo", "ho", "and", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  []string{"yo", "ho", "and", "barrel", "of", "rum"},
		},
		{
			name:  "one item does not match",
			items: []string{"yo", "ho", "and", "a", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  []string{"yo", "ho", "and", "barrel", "of", "rum"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.items, tt.f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Filter = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWithIndex_Int(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []int
		f     func(int, int) bool
		want  []int
	}{
		{
			name: "no items",
			want: nil,
		},
		{
			name:  "one item matches",
			items: []int{0},
			f:     func(i, v int) bool { return v >= 0 },
			want:  []int{0},
		},
		{
			name:  "no items match",
			items: []int{-1, -2, -3},
			f:     func(i, v int) bool { return v >= 0 },
			want:  nil,
		},
		{
			name:  "all items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "one item does not match at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "one item does not match at end",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1},
			f:     func(i, v int) bool { return v >= 0 },
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "one item does not match at middle",
			items: []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterWithIndex(tt.items, tt.f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FilterWithIndex = %v, want %v", got, tt.want)
			}
		})
	}
}
