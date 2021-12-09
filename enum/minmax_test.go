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

func TestMaxFunc_Key(t *testing.T) {
	type keyT [2]int
	tests := []struct {
		name  string
		items []keyT
		want  keyT
	}{
		{
			name: "no items",
			want: keyT{},
		},
		{
			name:  "one item",
			items: []keyT{{-1, -2}},
			want:  keyT{-1, -2},
		},
		{
			name:  "max at start",
			items: []keyT{{2, 10}, {-1, -2}, {10, 1}, {1, 10}},
			want:  keyT{2, 10},
		},
		{
			name:  "max at end",
			items: []keyT{{-1, -2}, {10, 1}, {1, 10}, {2, 10}},
			want:  keyT{2, 10},
		},
		{
			name:  "max at middle",
			items: []keyT{{-1, -2}, {3, 4}, {2, 10}, {10, 1}, {1, 10}},
			want:  keyT{2, 10},
		},
		{
			name:  "multiple max",
			items: []keyT{{2, 10}, {-1, -2}, {10, 1}, {2, 10}, {1, 10}, {2, 10}},
			want:  keyT{2, 10},
		},
	}

	lessFunc := func(a, b keyT) bool {
		if a[1] == b[1] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxFunc(tt.items, lessFunc)
			if got != tt.want {
				t.Errorf("MaxFunc = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFunc_Key(t *testing.T) {
	type keyT [2]int
	tests := []struct {
		name  string
		items []keyT
		want  keyT
	}{
		{
			name: "no items",
			want: keyT{},
		},
		{
			name:  "one item",
			items: []keyT{{-1, -2}},
			want:  keyT{-1, -2},
		},
		{
			name:  "min at start",
			items: []keyT{{-2, -10}, {-1, -2}, {10, 1}, {-1, -10}},
			want:  keyT{-2, -10},
		},
		{
			name:  "min at end",
			items: []keyT{{-1, -2}, {10, 1}, {-1, -10}, {-2, -10}},
			want:  keyT{-2, -10},
		},
		{
			name:  "min at middle",
			items: []keyT{{-1, -2}, {3, 4}, {-2, -10}, {10, 1}, {-1, -10}},
			want:  keyT{-2, -10},
		},
		{
			name:  "multiple min",
			items: []keyT{{-2, -10}, {-1, -2}, {10, 1}, {-2, -10}, {-1, -10}, {-2, -10}},
			want:  keyT{-2, -10},
		},
	}

	lessFunc := func(a, b keyT) bool {
		if a[1] == b[1] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinFunc(tt.items, lessFunc)
			if got != tt.want {
				t.Errorf("MinFunc = %v, want %v", got, tt.want)
			}
		})
	}
}
