// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/v1/enum"
	"github.com/gmlewis/advent-of-code-2021/v1/must"
)

var logf = log.Printf
var printf = fmt.Printf

func main() {
	flag.Parse()

	Each(flag.Args(), process)
}

func process(filename string) {
	logf("Processing %v ...", filename)
	buf := must.ReadFile(filename)
	bits := strings.Join(Map([]rune(buf), func(r rune) string { return hex2dec[r] }), "")

	insts := processBits(bits)
	v := insts[0].eval()

	printf("Solution: %v\n", v)
}

func processBits(bits string) []*instT {
	var insts []*instT
	for len(bits) > 7 {
		bits, insts = processNext(bits, insts)
	}
	return insts
}

func processNext(bits string, insts []*instT) (string, []*instT) {
	inst := &instT{
		version: dec2hex["0"+bits[0:3]],
		id:      dec2hex["0"+bits[3:6]],
	}

	switch {
	case inst.id == 4:
		return processLiteral(bits[6:], insts, inst)
	case bits[6] == '0':
		return processOpLenType0(bits[7:], insts, inst)
	default:
		return processOpLenType1(bits[7:], insts, inst)
	}
}

func processLiteral(bits string, insts []*instT, inst *instT) (string, []*instT) {
	var value string
	for {
		value += bits[1:5]
		if bits[0] == '0' {
			bits = bits[5:]
			break
		}
		bits = bits[5:]
	}
	inst.literal = must.ParseInt(value, 2, 64)
	return bits, append(insts, inst)
}

func processOpLenType0(bits string, insts []*instT, inst *instT) (string, []*instT) {
	totalBitLen := must.ParseInt(bits[0:15], 2, 64)
	bits = bits[15:]
	inst.subPackets = processBits(bits[0:totalBitLen])
	return bits[totalBitLen:], append(insts, inst)
}

func processOpLenType1(bits string, insts []*instT, inst *instT) (string, []*instT) {
	n := must.ParseInt(bits[0:11], 2, 64)
	bits = bits[11:]

	for i := 0; i < n; i++ {
		bits, inst.subPackets = processNext(bits, inst.subPackets)
	}
	return bits, append(insts, inst)
}

type instT struct {
	version int
	id      int

	// id=4
	literal int

	subPackets []*instT
}

func (in *instT) eval() int {
	switch in.id {
	case 0: // sum
		v := Reduce(in.subPackets, 0, func(in *instT, acc int) int { return acc + in.eval() })
		in.id, in.literal = 4, v
		return v
	case 1: // product
		v := Reduce(in.subPackets, 1, func(in *instT, acc int) int { return acc * in.eval() })
		in.id, in.literal = 4, v
		return v
	case 2: // minimum
		v := ReduceWithIndex(in.subPackets, 0, func(index int, in *instT, acc int) int {
			v := in.eval()
			if index == 0 || v < acc {
				acc = v
			}
			return acc
		})
		in.id, in.literal = 4, v
		return v
	case 3: // maximum
		v := ReduceWithIndex(in.subPackets, 0, func(index int, in *instT, acc int) int {
			v := in.eval()
			if index == 0 || v > acc {
				acc = v
			}
			return acc
		})
		in.id, in.literal = 4, v
		return v
	case 4:
		return in.literal
	case 5: // greater than
		a := in.subPackets[0].eval()
		b := in.subPackets[1].eval()
		v := 0
		if a > b {
			v = 1
		}
		in.id, in.literal = 4, v
		return v
	case 6: // less than
		a := in.subPackets[0].eval()
		b := in.subPackets[1].eval()
		v := 0
		if a < b {
			v = 1
		}
		in.id, in.literal = 4, v
		return v
	case 7: // equal to
		a := in.subPackets[0].eval()
		b := in.subPackets[1].eval()
		v := 0
		if a == b {
			v = 1
		}
		in.id, in.literal = 4, v
		return v
	}

	log.Fatalf("bad packet id=%v", in.id)
	return 0
}

var hex2dec = map[rune]string{
	'0': "0000", '1': "0001", '2': "0010", '3': "0011",
	'4': "0100", '5': "0101", '6': "0110", '7': "0111",
	'8': "1000", '9': "1001", 'A': "1010", 'B': "1011",
	'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111",
}

var dec2hex = map[string]int{
	"0000": 0, "0001": 1, "0010": 2, "0011": 3,
	"0100": 4, "0101": 5, "0110": 6, "0111": 7,
	"1000": 8, "1001": 9, "1010": 10, "1011": 11,
	"1100": 12, "1101": 13, "1110": 14, "1111": 15,
}
