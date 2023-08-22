package stream

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/google/go-cmp/cmp"
)

func TestMap_IntToString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		items []int
		f     func(int) string
		want  []string
	}{
		{
			name: "empty int to string",
			want: []string{},
		},
		{
			name:  "simple decimal int to string",
			items: []int{1, 2, 3, 100},
			f:     strconv.Itoa,
			want:  []string{"1", "2", "3", "100"},
		},
		{
			name:  "simple decimal int to hex string",
			items: []int{1, 2, 3, 128},
			f:     func(v int) string { return fmt.Sprintf("0x%02x", v) },
			want:  []string{"0x01", "0x02", "0x03", "0x80"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(ToChan(tt.items), tt.f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Map(%+v) = %+v, want %+v", tt.items, got, tt.want)
			}
		})
	}
}

func TestMap_StringToInt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		from []string
		f    func(string) int
		want []int
	}{
		{
			name: "empty string to int",
			from: []string{},
			want: []int{},
		},
		{
			name: "simple decimal string to int",
			from: []string{"1", "2", "3", "100"},
			f:    must.Atoi,
			want: []int{1, 2, 3, 100},
		},
		{
			name: "simple binary string to int",
			from: []string{"1", "10", "11", "10000000"},
			f:    func(s string) int { return must.ParseInt(s, 2, 64) },
			want: []int{1, 2, 3, 128},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(ToChan(tt.from), tt.f)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Map(%+v) = %+v, want %+v", tt.from, got, tt.want)
			}
		})
	}
}
