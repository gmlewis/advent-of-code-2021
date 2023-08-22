package stream

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestChunkEvery_Int(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		end  int
		n    int
		step int
		want [][]int
	}{
		{
			name: "1-element channel",
			n:    2,
			step: 1,
			want: nil,
		},
		{
			name: "2-element channel",
			end:  1,
			n:    2,
			step: 1,
			want: [][]int{{0, 1}},
		},
		{
			name: "5-element channel, n=2, step=1",
			end:  5,
			n:    2,
			step: 1,
			want: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}},
		},
		{
			name: "5-element channel, n=3, step=1",
			end:  5,
			n:    3,
			step: 1,
			want: [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
		},
		{
			name: "5-element channel, n=4, step=1",
			end:  5,
			n:    4,
			step: 1,
			want: [][]int{{0, 1, 2, 3}, {1, 2, 3, 4}, {2, 3, 4, 5}},
		},
		{
			name: "5-element channel, n=5, step=1",
			end:  5,
			n:    5,
			step: 1,
			want: [][]int{{0, 1, 2, 3, 4}, {1, 2, 3, 4, 5}},
		},
		{
			name: "5-element channel, n=6, step=1",
			end:  5,
			n:    6,
			step: 1,
			want: [][]int{{0, 1, 2, 3, 4, 5}},
		},
		{
			name: "5-element channel, n=2, step=2",
			end:  5,
			n:    2,
			step: 2,
			want: [][]int{{0, 1}, {2, 3}, {4, 5}},
		},
		{
			name: "5-element channel, n=3, step=2",
			end:  5,
			n:    3,
			step: 2,
			want: [][]int{{0, 1, 2}, {2, 3, 4}},
		},
		{
			name: "5-element channel, n=4, step=2",
			end:  5,
			n:    4,
			step: 2,
			want: [][]int{{0, 1, 2, 3}, {2, 3, 4, 5}},
		},
		{
			name: "5-element channel, n=5, step=2",
			end:  5,
			n:    5,
			step: 2,
			want: [][]int{{0, 1, 2, 3, 4}},
		},
		{
			name: "5-element channel, n=6, step=2",
			end:  5,
			n:    6,
			step: 2,
			want: [][]int{{0, 1, 2, 3, 4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSlice(ChunkEvery(Range(0, tt.end), tt.n, tt.step))
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ChunkEvery = %+v, want %+v", got, tt.want)
			}
		})
	}
}
