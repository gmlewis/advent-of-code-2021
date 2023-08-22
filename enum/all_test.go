package enum

import "testing"

func TestAll_Int(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []int
		f     func(int) bool
		want  bool
	}{
		{
			name: "no items",
			want: true,
		},
		{
			name:  "one item matches",
			items: []int{0},
			f:     func(v int) bool { return v >= 0 },
			want:  true,
		},
		{
			name:  "no items match",
			items: []int{-1, -2, -3},
			f:     func(v int) bool { return v >= 0 },
			want:  false,
		},
		{
			name:  "all items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  true,
		},
		{
			name:  "one item does not match at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  false,
		},
		{
			name:  "one item does not match at end",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1},
			f:     func(v int) bool { return v >= 0 },
			want:  false,
		},
		{
			name:  "one item does not match at middle",
			items: []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := All(tt.items, tt.f)
			if got != tt.want {
				t.Errorf("All = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []string
		f     func(string) bool
		want  bool
	}{
		{
			name: "no items",
			want: true,
		},
		{
			name:  "one item matches",
			items: []string{"yo"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  true,
		},
		{
			name:  "all items match",
			items: []string{"yo", "ho", "and", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  true,
		},
		{
			name:  "one item does not match",
			items: []string{"yo", "ho", "and", "a", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := All(tt.items, tt.f)
			if got != tt.want {
				t.Errorf("All = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllWithIndex_Int(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []int
		f     func(int, int) bool
		want  bool
	}{
		{
			name: "no items",
			want: true,
		},
		{
			name:  "one item matches",
			items: []int{0},
			f:     func(i, v int) bool { return v >= 0 },
			want:  true,
		},
		{
			name:  "all items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  true,
		},
		{
			name:  "one item does not match at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  false,
		},
		{
			name:  "one item does not match at end",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1},
			f:     func(i, v int) bool { return v >= 0 },
			want:  false,
		},
		{
			name:  "one item does not match at middle",
			items: []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllWithIndex(tt.items, tt.f)
			if got != tt.want {
				t.Errorf("AllWithIndex = %v, want %v", got, tt.want)
			}
		})
	}
}
