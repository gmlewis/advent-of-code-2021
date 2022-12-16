package stream

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPermutationsOf_String(t *testing.T) {
	tests := []struct {
		name string
		seq  []string
		want [][]string
	}{
		{
			name: "nil",
		},
		{
			name: "empty",
			seq:  []string{},
		},
		{
			name: "one element",
			seq:  []string{"A"},
			want: [][]string{{"A"}},
		},
		{
			name: "two elements",
			seq:  []string{"A", "B"},
			want: [][]string{
				{"A", "B"},
				{"B", "A"},
			},
		},
		{
			name: "three elements",
			seq:  []string{"A", "B", "C"},
			want: [][]string{
				{"A", "B", "C"},
				{"B", "A", "C"},
				{"B", "C", "A"},
				{"A", "C", "B"},
				{"C", "A", "B"},
				{"C", "B", "A"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := PermutationsOf(tt.seq)
			var got [][]string
			for seq := range ch {
				got = append(got, seq)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("PermutationsOf mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
