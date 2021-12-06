package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestChunkEvery_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		n     int
		step  int
		want  [][]int
	}{
		{
			name: "no items, n=0, step=0",
			want: [][]int{},
		},
		{
			name:  "5 items, n=0, step=1",
			items: []int{1, 2, 3, 4, 5},
			step:  1,
			want:  [][]int{},
		},
		{
			name:  "5 items, n=1, step=0",
			items: []int{1, 2, 3, 4, 5},
			n:     1,
			want:  [][]int{},
		},
		{
			name:  "5 items, n=1, step=1",
			items: []int{1, 2, 3, 4, 5},
			n:     1,
			step:  1,
			want:  [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name:  "one item, n=2, step=1",
			items: []int{1},
			n:     2,
			step:  1,
			want:  [][]int{},
		},
		{
			name:  "two items, n=2, step=1",
			items: []int{1, 2},
			n:     2,
			step:  1,
			want:  [][]int{{1, 2}},
		},
		{
			name:  "three items, n=2, step=1",
			items: []int{1, 2, 3},
			n:     2,
			step:  1,
			want:  [][]int{{1, 2}, {2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChunkEvery(tt.items, tt.n, tt.step)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ChunkEvery = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunkEvery_IntRange(t *testing.T) {
	tests := []struct {
		name string
		end  int
		n    int
		step int
		want [][]int
	}{
		{
			name: "1-element slice",
			n:    2,
			step: 1,
			want: [][]int{},
		},
		{
			name: "2-element slice",
			end:  1,
			n:    2,
			step: 1,
			want: [][]int{{0, 1}},
		},
		{
			name: "5-element slice, n=2, step=1",
			end:  5,
			n:    2,
			step: 1,
			want: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}},
		},
		{
			name: "5-element slice, n=3, step=1",
			end:  5,
			n:    3,
			step: 1,
			want: [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
		},
		{
			name: "5-element slice, n=4, step=1",
			end:  5,
			n:    4,
			step: 1,
			want: [][]int{{0, 1, 2, 3}, {1, 2, 3, 4}, {2, 3, 4, 5}},
		},
		{
			name: "5-element slice, n=5, step=1",
			end:  5,
			n:    5,
			step: 1,
			want: [][]int{{0, 1, 2, 3, 4}, {1, 2, 3, 4, 5}},
		},
		{
			name: "5-element slice, n=6, step=1",
			end:  5,
			n:    6,
			step: 1,
			want: [][]int{{0, 1, 2, 3, 4, 5}},
		},
		{
			name: "5-element slice, n=2, step=2",
			end:  5,
			n:    2,
			step: 2,
			want: [][]int{{0, 1}, {2, 3}, {4, 5}},
		},
		{
			name: "5-element slice, n=3, step=2",
			end:  5,
			n:    3,
			step: 2,
			want: [][]int{{0, 1, 2}, {2, 3, 4}},
		},
		{
			name: "5-element slice, n=4, step=2",
			end:  5,
			n:    4,
			step: 2,
			want: [][]int{{0, 1, 2, 3}, {2, 3, 4, 5}},
		},
		{
			name: "5-element slice, n=5, step=2",
			end:  5,
			n:    5,
			step: 2,
			want: [][]int{{0, 1, 2, 3, 4}},
		},
		{
			name: "5-element slice, n=6, step=2",
			end:  5,
			n:    6,
			step: 2,
			want: [][]int{{0, 1, 2, 3, 4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChunkEvery(Range(0, tt.end), tt.n, tt.step)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ChunkEvery = %+v, want %+v", got, tt.want)
			}
		})
	}
}
