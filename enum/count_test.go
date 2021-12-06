package enum

import "testing"

func TestCount_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		f     func(int) bool
		want  int
	}{
		{
			name: "no items",
			want: 0,
		},
		{
			name:  "one item matches",
			items: []int{0},
			f:     func(v int) bool { return v >= 0 },
			want:  1,
		},
		{
			name:  "multiple items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  11,
		},
		{
			name:  "one item does not match at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  11,
		},
		{
			name:  "one item does not match at end",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1},
			f:     func(v int) bool { return v >= 0 },
			want:  11,
		},
		{
			name:  "one item does not match at middle",
			items: []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Count(tt.items, tt.f)
			if got != tt.want {
				t.Errorf("Count = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount_String(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		f     func(string) bool
		want  int
	}{
		{
			name: "no items",
			want: 0,
		},
		{
			name:  "one item matches",
			items: []string{"yo"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  1,
		},
		{
			name:  "multiple items match",
			items: []string{"yo", "ho", "and", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  6,
		},
		{
			name:  "one item does not match",
			items: []string{"yo", "ho", "and", "a", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Count(tt.items, tt.f)
			if got != tt.want {
				t.Errorf("Count = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountWithIndex_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		f     func(int, int) bool
		want  int
	}{
		{
			name: "no items",
			want: 0,
		},
		{
			name:  "one item matches",
			items: []int{0},
			f:     func(i, v int) bool { return v >= 0 },
			want:  1,
		},
		{
			name:  "multiple items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  11,
		},
		{
			name:  "one item does not match at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  11,
		},
		{
			name:  "one item does not match at end",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1},
			f:     func(i, v int) bool { return v >= 0 },
			want:  11,
		},
		{
			name:  "one item does not match at middle",
			items: []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWithIndex(tt.items, tt.f)
			if got != tt.want {
				t.Errorf("CountWithIndex = %v, want %v", got, tt.want)
			}
		})
	}
}
