package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	tests := []struct {
		name string
		hex  string
		want string
	}{
		{
			name: "example1",
			hex:  "D2FE28",
			want: "Solution: 6\n",
		},
		{
			name: "example2",
			hex:  "38006F45291200",
			want: "Solution: 9\n",
		},
		{
			name: "example3",
			hex:  "EE00D40C823060",
			want: "Solution: 14\n",
		},
		{
			name: "example4",
			hex:  "8A004A801A8002F478",
			want: "Solution: 16\n",
		},
		{
			name: "example5",
			hex:  "620080001611562C8802118E34",
			want: "Solution: 12\n",
		},
		{
			name: "example6",
			hex:  "C0015000016115A2E0802F182340",
			want: "Solution: 23\n",
		},
		{
			name: "example7",
			hex:  "A0016C880162017C3686B18A3D4780",
			want: "Solution: 31\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test.Runner(t, tt.hex, tt.want, process, &printf)
		})
	}
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}
