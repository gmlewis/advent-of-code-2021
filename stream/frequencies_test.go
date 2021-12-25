package stream

import (
	"strconv"
	"testing"

	"github.com/gmlewis/advent-of-code-2021/must"
	"github.com/google/go-cmp/cmp"
)

func TestFrequencies_Int(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  map[int]int
	}{
		{
			name: "no items",
			want: map[int]int{},
		},
		{
			name:  "one item",
			items: []int{0},
			want:  map[int]int{0: 1},
		},
		{
			name:  "3 items, various times",
			items: []int{0, 1, 2, 0, 1, 0},
			want:  map[int]int{0: 3, 1: 2, 2: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Frequencies(ToChan(tt.items))
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Frequencies = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrequencies_String(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  map[string]int
	}{
		{
			name: "no items",
			want: map[string]int{},
		},
		{
			name:  "one item",
			items: []string{"0"},
			want:  map[string]int{"0": 1},
		},
		{
			name:  "3 items, various times",
			items: []string{"0", "1", "2", "0", "1", "0"},
			want:  map[string]int{"0": 3, "1": 2, "2": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Frequencies(ToChan(tt.items))
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Frequencies = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrequenciesBy_IntToString(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  map[string]int
	}{
		{
			name: "no items",
			want: map[string]int{},
		},
		{
			name:  "one item",
			items: []int{0},
			want:  map[string]int{"0": 1},
		},
		{
			name:  "3 items, various times",
			items: []int{0, 1, 2, 0, 1, 0},
			want:  map[string]int{"0": 3, "1": 2, "2": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FrequenciesBy(ToChan(tt.items), strconv.Itoa)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FrequenciesBy = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrequenciesBy_StringToInt(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		want  map[int]int
	}{
		{
			name: "no items",
			want: map[int]int{},
		},
		{
			name:  "one item",
			items: []string{"0"},
			want:  map[int]int{0: 1},
		},
		{
			name:  "3 items, various times",
			items: []string{"0", "1", "2", "0", "1", "0"},
			want:  map[int]int{0: 3, 1: 2, 2: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FrequenciesBy(ToChan(tt.items), must.Atoi)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("FrequenciesBy = %v, want %v", got, tt.want)
			}
		})
	}
}
