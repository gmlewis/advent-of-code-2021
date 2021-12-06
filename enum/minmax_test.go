package enum

import "testing"

func TestMax_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  int
	}{
		{
			name: "no items",
			want: 0,
		},
		{
			name:  "one item",
			items: []int{-1},
			want:  -1,
		},
		{
			name:  "max at start",
			items: []int{10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1},
			want:  10,
		},
		{
			name:  "max at end",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  10,
		},
		{
			name:  "max at middle",
			items: []int{0, 1, 2, 3, 4, 10, 5, 6, 7, 8, 9, -1},
			want:  10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.items)
			if got != tt.want {
				t.Errorf("Max = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax_String(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  string
	}{
		{
			name: "no items",
			want: "",
		},
		{
			name:  "one item",
			items: []string{"yo"},
			want:  "yo",
		},
		{
			name:  "max at start",
			items: []string{"yo", "ho", "and", "barrel", "of", "rum"},
			want:  "yo",
		},
		{
			name:  "max at end",
			items: []string{"a", "b", "c", "yo"},
			want:  "yo",
		},
		{
			name:  "max at middle",
			items: []string{"a", "b", "yo", "c", "d"},
			want:  "yo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.items)
			if got != tt.want {
				t.Errorf("Max = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  int
	}{
		{
			name: "no items",
			want: 0,
		},
		{
			name:  "one item",
			items: []int{-1},
			want:  -1,
		},
		{
			name:  "min at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  -1,
		},
		{
			name:  "min at end",
			items: []int{10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1},
			want:  -1,
		},
		{
			name:  "min at middle",
			items: []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10},
			want:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Min(tt.items)
			if got != tt.want {
				t.Errorf("Min = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin_String(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  string
	}{
		{
			name: "no items",
			want: "",
		},
		{
			name:  "one item",
			items: []string{"yo"},
			want:  "yo",
		},
		{
			name:  "min at start",
			items: []string{"a", "b", "c", "yo"},
			want:  "a",
		},
		{
			name:  "min at end",
			items: []string{"d", "b", "yo", "c", "a"},
			want:  "a",
		},
		{
			name:  "min at middle",
			items: []string{"yo", "ho", "and", "barrel", "of", "rum"},
			want:  "and",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Min(tt.items)
			if got != tt.want {
				t.Errorf("Min = %v, want %v", got, tt.want)
			}
		})
	}
}
