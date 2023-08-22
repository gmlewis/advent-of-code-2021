package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEach_Int(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []int
		want  []int
	}{
		{
			name: "no items",
			want: nil,
		},
		{
			name:  "one item",
			items: []int{0},
			want:  []int{0},
		},
		{
			name:  "multiple items",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			f := func(v int) { got = append(got, v) }
			Each(tt.items, f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Each = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEach_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []string
		want  []string
	}{
		{
			name: "no items",
			want: nil,
		},
		{
			name:  "one item",
			items: []string{"a"},
			want:  []string{"a"},
		},
		{
			name:  "multiple items",
			items: []string{"a", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			f := func(v string) { got = append(got, v) }
			Each(tt.items, f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Each = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEachWithIndex_Int(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []int
		want  []int
	}{
		{
			name: "no items",
			want: nil,
		},
		{
			name:  "one item",
			items: []int{1},
			want:  []int{0, 1},
		},
		{
			name:  "multiple items",
			items: []int{1, 2, 3},
			want:  []int{0, 1, 1, 2, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			f := func(i, v int) { got = append(got, i, v) }
			EachWithIndex(tt.items, f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("EachWithIndex = %v, want %v", got, tt.want)
			}
		})
	}
}
