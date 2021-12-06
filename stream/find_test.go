package stream

import (
	"testing"
)

func TestFind_Int(t *testing.T) {
	tests := []struct {
		name   string
		items  []int
		f      func(int) bool
		want   int
		wantOK bool
	}{
		{
			name:   "no items",
			want:   0,
			wantOK: false,
		},
		{
			name:   "one item matches",
			items:  []int{0},
			f:      func(v int) bool { return v >= 0 },
			want:   0,
			wantOK: true,
		},
		{
			name:   "no items match",
			items:  []int{-1, -2, -3},
			f:      func(v int) bool { return v >= 0 },
			want:   0,
			wantOK: false,
		},
		{
			name:   "all items match",
			items:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:      func(v int) bool { return v >= 0 },
			want:   0,
			wantOK: true,
		},
		{
			name:   "one item does not match at start",
			items:  []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:      func(v int) bool { return v >= 0 },
			want:   0,
			wantOK: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Find(ToChan(tt.items), tt.f)
			if ok != tt.wantOK || got != tt.want {
				t.Errorf("Find = %v,%v, want %v,%v", got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestFind_String(t *testing.T) {
	tests := []struct {
		name   string
		items  []string
		f      func(string) bool
		want   string
		wantOK bool
	}{
		{
			name:   "no items",
			want:   "",
			wantOK: false,
		},
		{
			name:   "one item matches",
			items:  []string{"yo"},
			f:      func(v string) bool { return len(v) >= 2 },
			want:   "yo",
			wantOK: true,
		},
		{
			name:   "no items match",
			items:  []string{"a", "b", "c"},
			f:      func(v string) bool { return len(v) >= 2 },
			want:   "",
			wantOK: false,
		},
		{
			name:   "all items match",
			items:  []string{"yo", "ho", "and", "barrel", "of", "rum"},
			f:      func(v string) bool { return len(v) >= 2 },
			want:   "yo",
			wantOK: true,
		},
		{
			name:   "first item does not match",
			items:  []string{"a", "yo", "ho", "and", "a", "barrel", "of", "rum"},
			f:      func(v string) bool { return len(v) >= 2 },
			want:   "yo",
			wantOK: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Find(ToChan(tt.items), tt.f)
			if ok != tt.wantOK || got != tt.want {
				t.Errorf("Find = %v,%v, want %v,%v", got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestFindWithIndex_Int(t *testing.T) {
	tests := []struct {
		name   string
		items  []int
		f      func(int, int) bool
		want   int
		wantOK bool
	}{
		{
			name:   "no items",
			want:   0,
			wantOK: false,
		},
		{
			name:   "one item matches",
			items:  []int{0},
			f:      func(i, v int) bool { return v >= 0 },
			want:   0,
			wantOK: true,
		},
		{
			name:   "no items match",
			items:  []int{-1, -2, -3},
			f:      func(i, v int) bool { return v >= 0 },
			want:   0,
			wantOK: false,
		},
		{
			name:   "all items match",
			items:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:      func(i, v int) bool { return v >= 0 },
			want:   0,
			wantOK: true,
		},
		{
			name:   "one item does not match at start",
			items:  []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:      func(i, v int) bool { return v >= 0 },
			want:   0,
			wantOK: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := FindWithIndex(ToChan(tt.items), tt.f)
			if ok != tt.wantOK || got != tt.want {
				t.Errorf("FindWithIndex = %v,%v, want %v,%v", got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestFindOr_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		f     func(int) bool
		want  int
	}{
		{
			name: "no items",
			want: -99,
		},
		{
			name:  "one item matches",
			items: []int{0},
			f:     func(v int) bool { return v >= 0 },
			want:  0,
		},
		{
			name:  "no items match",
			items: []int{-1, -2, -3},
			f:     func(v int) bool { return v >= 0 },
			want:  -99,
		},
		{
			name:  "all items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  0,
		},
		{
			name:  "one item does not match at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(v int) bool { return v >= 0 },
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindOr(ToChan(tt.items), tt.f, -99)
			if got != tt.want {
				t.Errorf("FindOr = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOr_String(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		f     func(string) bool
		want  string
	}{
		{
			name: "no items",
			want: "negatory",
		},
		{
			name:  "one item matches",
			items: []string{"yo"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  "yo",
		},
		{
			name:  "no items match",
			items: []string{"a", "b", "c"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  "negatory",
		},
		{
			name:  "all items match",
			items: []string{"yo", "ho", "and", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  "yo",
		},
		{
			name:  "first item does not match",
			items: []string{"a", "yo", "ho", "and", "a", "barrel", "of", "rum"},
			f:     func(v string) bool { return len(v) >= 2 },
			want:  "yo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindOr(ToChan(tt.items), tt.f, "negatory")
			if got != tt.want {
				t.Errorf("FindOr = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOrWithIndex_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		f     func(int, int) bool
		want  int
	}{
		{
			name: "no items",
			want: -99,
		},
		{
			name:  "one item matches",
			items: []int{0},
			f:     func(i, v int) bool { return v >= 0 },
			want:  0,
		},
		{
			name:  "no items match",
			items: []int{-1, -2, -3},
			f:     func(i, v int) bool { return v >= 0 },
			want:  -99,
		},
		{
			name:  "all items match",
			items: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  0,
		},
		{
			name:  "one item does not match at start",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			f:     func(i, v int) bool { return v >= 0 },
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindOrWithIndex(ToChan(tt.items), tt.f, -99)
			if got != tt.want {
				t.Errorf("FindOrWithIndex = %v, want %v", got, tt.want)
			}
		})
	}
}
