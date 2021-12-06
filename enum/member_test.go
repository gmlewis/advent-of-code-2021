package enum

import "testing"

func TestMember_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		elem  int
		want  bool
	}{
		{
			name: "no items",
			want: false,
		},
		{
			name:  "one item member",
			items: []int{-1},
			elem:  -1,
			want:  true,
		},
		{
			name:  "one item non-member",
			items: []int{-1},
			elem:  0,
			want:  false,
		},
		{
			name:  "member at start",
			items: []int{10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1},
			elem:  10,
			want:  true,
		},
		{
			name:  "member at end",
			items: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			elem:  10,
			want:  true,
		},
		{
			name:  "member at middle",
			items: []int{0, 1, 2, 3, 4, 10, 5, 6, 7, 8, 9, -1},
			elem:  10,
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Member(tt.items, tt.elem)
			if got != tt.want {
				t.Errorf("Member = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMember_String(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		elem  string
		want  bool
	}{
		{
			name: "no items",
			want: false,
		},
		{
			name:  "one item member",
			items: []string{"yo"},
			elem:  "yo",
			want:  true,
		},
		{
			name:  "one item non-member",
			items: []string{"yo"},
			elem:  "ho",
			want:  false,
		},
		{
			name:  "member at start",
			items: []string{"yo", "ho", "and", "barrel", "of", "rum"},
			elem:  "yo",
			want:  true,
		},
		{
			name:  "member at end",
			items: []string{"a", "b", "c", "yo"},
			elem:  "yo",
			want:  true,
		},
		{
			name:  "member at middle",
			items: []string{"a", "b", "yo", "c", "d"},
			elem:  "yo",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Member(tt.items, tt.elem)
			if got != tt.want {
				t.Errorf("Member = %v, want %v", got, tt.want)
			}
		})
	}
}
