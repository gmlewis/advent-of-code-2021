package stream

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPermutationsOf_String(t *testing.T) {
	t.Parallel()
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

func TestBigPermutationsOf(t *testing.T) {
	t.Parallel()
	// bigSeq := []string{"UU", "EW", "WJ", "OZ", "ZM", "RU", "WH", "MJ", "UD", "FD", "CO", "DW", "PL", "YJ", "ZI"}  // 15 items is too big => 1,307,674,368,000
	bigSeq := []string{"UU", "EW", "WJ", "OZ", "ZM", "RU", "WH", "MJ", "UD", "FD"}
	// 10 items => 1.27 seconds, 11 items => 55.38 seconds.

	ch := PermutationsOf(bigSeq)
	var got int
	for range ch {
		got++
	}

	if want := 3628800; got != want {
		t.Errorf("PermutationsOf(%v items) = %v, want %v", len(bigSeq), got, want)
	}
}
