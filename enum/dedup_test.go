package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDedup_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  []int
	}{
		{
			name: "no items",
			want: []int{},
		},
		{
			name:  "one item",
			items: []int{0},
			want:  []int{0},
		},
		{
			name:  "multiple items, no dups",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "one consecutive dup",
			items: []int{0, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "multiple consecutive dups",
			items: []int{0, 0, 1, 1, 1, 2, 2, 2, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9, 10, 10, 10},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dedup(tt.items)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Dedup = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDedup_String(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  []string
	}{
		{
			name: "no items",
			want: []string{},
		},
		{
			name:  "one item",
			items: []string{"a"},
			want:  []string{"a"},
		},
		{
			name:  "multiple items, no dups",
			items: []string{"a", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "one consecutive dup",
			items: []string{"a", "b", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "multiple consecutive dups",
			items: []string{"a", "a", "a", "b", "b", "c", "c", "c", "c"},
			want:  []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dedup(tt.items)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Dedup = %v, want %v", got, tt.want)
			}
		})
	}
}
