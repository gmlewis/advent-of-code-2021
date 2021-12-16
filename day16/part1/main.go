// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
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
	logf("hex %v: bits: %v", buf, bits)

	insts := processBits(bits)

	sumVersions := Reduce(insts, 0, func(inst *instT, acc int) int { return acc + inst.sumVersions() })

	printf("Solution: %v\n", sumVersions)
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
	logf("literal: inst=%#v", *inst)
	return bits, append(insts, inst)
}

func processOpLenType0(bits string, insts []*instT, inst *instT) (string, []*instT) {
	inst.totalBitLen = must.ParseInt(bits[0:15], 2, 64)
	bits = bits[15:]
	inst.subPackets = processBits(bits[0:inst.totalBitLen])
	logf("opType0: inst=%#v", *inst)
	return bits[inst.totalBitLen:], append(insts, inst)
}

func processOpLenType1(bits string, insts []*instT, inst *instT) (string, []*instT) {
	inst.numSubPackets = must.ParseInt(bits[0:11], 2, 64)
	bits = bits[11:]

	for i := 0; i < inst.numSubPackets; i++ {
		bits, inst.subPackets = processNext(bits, inst.subPackets)
	}
	logf("opType1: inst=%#v", *inst)
	return bits, append(insts, inst)
}

type instT struct {
	version int
	id      int

	// id=4
	literal int

	// id!=4
	lenTypeID int

	// lenTypeID=0
	totalBitLen int // (15 bits)

	// lenTypeID=1
	numSubPackets int // (11 bits)

	subPackets []*instT
}

func (in *instT) sumVersions() int {
	sum := in.version
	for _, sub := range in.subPackets {
		sum += sub.sumVersions()
	}
	return sum
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
