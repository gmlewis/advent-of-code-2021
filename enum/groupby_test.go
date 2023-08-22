package enum

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/strfn"
	"github.com/google/go-cmp/cmp"
)

func TestGroupBy_StringToInt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		values    []string
		keyFunc   func(string) int
		valueFunc func(string) string
		want      map[int][]string
	}{
		{
			name:   "empty int to string",
			values: []string{},
			want:   map[int][]string{},
		},
		{
			name:      "doc example",
			values:    []string{"ant", "buffalo", "cat", "dingo"},
			keyFunc:   strfn.Length,
			valueFunc: Identity[string],
			want: map[int][]string{
				3: {"ant", "cat"}, 5: {"dingo"}, 7: {"buffalo"},
			},
		},
		{
			name:      "doc example",
			values:    []string{"ant", "buffalo", "cat", "dingo"},
			keyFunc:   strfn.Length,
			valueFunc: strfn.First,
			want: map[int][]string{
				3: {"a", "c"}, 5: {"d"}, 7: {"b"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupBy(tt.values, tt.keyFunc, tt.valueFunc)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("GroupBy(%+v) = %+v, want %+v", tt.values, got, tt.want)
			}
		})
	}
}
