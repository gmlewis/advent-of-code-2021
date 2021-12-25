package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/v1/test"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		name string
		hex  string
		want string
	}{
		{
			name: "example21 - sum",
			hex:  "C200B40A82",
			want: "Solution: 3\n",
		},
		{
			name: "example22 - product",
			hex:  "04005AC33890",
			want: "Solution: 54\n",
		},
		{
			name: "example23 - minimum",
			hex:  "880086C3E88112",
			want: "Solution: 7\n",
		},
		{
			name: "example24 - maximum",
			hex:  "CE00C43D881120",
			want: "Solution: 9\n",
		},
		{
			name: "example25 - less than",
			hex:  "D8005AC2A8F0",
			want: "Solution: 1\n",
		},
		{
			name: "example26 - greater than",
			hex:  "F600BC2D8F",
			want: "Solution: 0\n",
		},
		{
			name: "example27 - equal",
			hex:  "9C005AC2F8F0",
			want: "Solution: 0\n",
		},
		{
			name: "example28 - 1+1=2*2",
			hex:  "9C0141080250320F1802104A08",
			want: "Solution: 1\n",
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
