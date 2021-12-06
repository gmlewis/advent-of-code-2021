package enum

import "testing"

func TestAny_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		f     func(int) bool
		want  bool
	}{
		{
			name: "no items",
			want: false,
		},
		{
			name:  "one item does not match",
			items: []int{0},
			f:     func(v int) bool { return v < 0 },
			want:  false,
		},
		{
			name:  "no items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v < 0 },
			want:  false,
		},
		{
			name:  "one item matches at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v < 0 },
			want:  true,
		},
		{
			name:  "one item matches at end",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1},
			f:     func(v int) bool { return v < 0 },
			want:  true,
		},
		{
			name:  "one item matches at middle",
			items: []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v < 0 },
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Any(tt.items, tt.f)
			if got != tt.want {
				t.Errorf("Any = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_String(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		f     func(string) bool
		want  bool
	}{
		{
			name: "no items",
			want: false,
		},
		{
			name:  "one item does not match",
			items: []string{"yo"},
			f:     func(v string) bool { return len(v) < 2 },
			want:  false,
		},
		{
			name:  "no items match",
			items: []string{"yo", "ho", "and", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) < 2 },
			want:  false,
		},
		{
			name:  "one item matches",
			items: []string{"yo", "ho", "and", "a", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) < 2 },
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Any(tt.items, tt.f)
			if got != tt.want {
				t.Errorf("Any = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyWithIndex_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		f     func(int, int) bool
		want  bool
	}{
		{
			name: "no items",
			want: false,
		},
		{
			name:  "one item does not match",
			items: []int{0},
			f:     func(i, v int) bool { return v < 0 },
			want:  false,
		},
		{
			name:  "no items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v < 0 },
			want:  false,
		},
		{
			name:  "one item matches at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v < 0 },
			want:  true,
		},
		{
			name:  "one item matches at end",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1},
			f:     func(i, v int) bool { return v < 0 },
			want:  true,
		},
		{
			name:  "one item matches at middle",
			items: []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v < 0 },
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnyWithIndex(tt.items, tt.f)
			if got != tt.want {
				t.Errorf("AnyWithIndex = %v, want %v", got, tt.want)
			}
		})
	}
}
