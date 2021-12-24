// -*- compile-command: "go run main.go ../input.txt"; -*-

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
	lines := must.ReadFileLines(filename)

	// digit = 0
	// code := MapWithIndex(lines, translate)
	// printf("%v\n", strings.Join(code, "\n"))

	printf("Solution: %v\n", len(lines))
}

func input(digits [14]int64) int64 {
	var w, x, y, z int64

	w = digits[0] // line 1: inp w
	x = 0         // line 2: mul x 0
	x += z        // line 3: add x z
	x %= 26       // line 4: mod x 26
	// line 5: div z 1
	x += 14 // line 6: add x 14
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 7: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 8: eql x 0
	y = 0         // line 9: mul y 0
	y += 25       // line 10: add y 25
	y *= x        // line 11: mul y x
	y += 1        // line 12: add y 1
	z *= y        // line 13: mul z y
	y = 0         // line 14: mul y 0
	y += w        // line 15: add y w
	y += 12       // line 16: add y 12
	y *= x        // line 17: mul y x
	z += y        // line 18: add z y
	w = digits[1] // line 19: inp w
	x = 0         // line 20: mul x 0
	x += z        // line 21: add x z
	x %= 26       // line 22: mod x 26
	// line 23: div z 1
	x += 10 // line 24: add x 10
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 25: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 26: eql x 0
	y = 0         // line 27: mul y 0
	y += 25       // line 28: add y 25
	y *= x        // line 29: mul y x
	y += 1        // line 30: add y 1
	z *= y        // line 31: mul z y
	y = 0         // line 32: mul y 0
	y += w        // line 33: add y w
	y += 9        // line 34: add y 9
	y *= x        // line 35: mul y x
	z += y        // line 36: add z y
	w = digits[2] // line 37: inp w
	x = 0         // line 38: mul x 0
	x += z        // line 39: add x z
	x %= 26       // line 40: mod x 26
	// line 41: div z 1
	x += 13 // line 42: add x 13
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 43: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 44: eql x 0
	y = 0         // line 45: mul y 0
	y += 25       // line 46: add y 25
	y *= x        // line 47: mul y x
	y += 1        // line 48: add y 1
	z *= y        // line 49: mul z y
	y = 0         // line 50: mul y 0
	y += w        // line 51: add y w
	y += 8        // line 52: add y 8
	y *= x        // line 53: mul y x
	z += y        // line 54: add z y
	w = digits[3] // line 55: inp w
	x = 0         // line 56: mul x 0
	x += z        // line 57: add x z
	x %= 26       // line 58: mod x 26
	z /= 26       // line 59: div z 26
	x += -8       // line 60: add x -8
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 61: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 62: eql x 0
	y = 0         // line 63: mul y 0
	y += 25       // line 64: add y 25
	y *= x        // line 65: mul y x
	y += 1        // line 66: add y 1
	z *= y        // line 67: mul z y
	y = 0         // line 68: mul y 0
	y += w        // line 69: add y w
	y += 3        // line 70: add y 3
	y *= x        // line 71: mul y x
	z += y        // line 72: add z y
	w = digits[4] // line 73: inp w
	x = 0         // line 74: mul x 0
	x += z        // line 75: add x z
	x %= 26       // line 76: mod x 26
	// line 77: div z 1
	x += 11 // line 78: add x 11
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 79: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 80: eql x 0
	y = 0         // line 81: mul y 0
	y += 25       // line 82: add y 25
	y *= x        // line 83: mul y x
	y += 1        // line 84: add y 1
	z *= y        // line 85: mul z y
	y = 0         // line 86: mul y 0
	y += w        // line 87: add y w
	y += 0        // line 88: add y 0
	y *= x        // line 89: mul y x
	z += y        // line 90: add z y
	w = digits[5] // line 91: inp w
	x = 0         // line 92: mul x 0
	x += z        // line 93: add x z
	x %= 26       // line 94: mod x 26
	// line 95: div z 1
	x += 11 // line 96: add x 11
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 97: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 98: eql x 0
	y = 0         // line 99: mul y 0
	y += 25       // line 100: add y 25
	y *= x        // line 101: mul y x
	y += 1        // line 102: add y 1
	z *= y        // line 103: mul z y
	y = 0         // line 104: mul y 0
	y += w        // line 105: add y w
	y += 11       // line 106: add y 11
	y *= x        // line 107: mul y x
	z += y        // line 108: add z y
	w = digits[6] // line 109: inp w
	x = 0         // line 110: mul x 0
	x += z        // line 111: add x z
	x %= 26       // line 112: mod x 26
	// line 113: div z 1
	x += 14 // line 114: add x 14
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 115: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 116: eql x 0
	y = 0         // line 117: mul y 0
	y += 25       // line 118: add y 25
	y *= x        // line 119: mul y x
	y += 1        // line 120: add y 1
	z *= y        // line 121: mul z y
	y = 0         // line 122: mul y 0
	y += w        // line 123: add y w
	y += 10       // line 124: add y 10
	y *= x        // line 125: mul y x
	z += y        // line 126: add z y
	w = digits[7] // line 127: inp w
	x = 0         // line 128: mul x 0
	x += z        // line 129: add x z
	x %= 26       // line 130: mod x 26
	z /= 26       // line 131: div z 26
	x += -11      // line 132: add x -11
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 133: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 134: eql x 0
	y = 0         // line 135: mul y 0
	y += 25       // line 136: add y 25
	y *= x        // line 137: mul y x
	y += 1        // line 138: add y 1
	z *= y        // line 139: mul z y
	y = 0         // line 140: mul y 0
	y += w        // line 141: add y w
	y += 13       // line 142: add y 13
	y *= x        // line 143: mul y x
	z += y        // line 144: add z y
	w = digits[8] // line 145: inp w
	x = 0         // line 146: mul x 0
	x += z        // line 147: add x z
	x %= 26       // line 148: mod x 26
	// line 149: div z 1
	x += 14 // line 150: add x 14
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 151: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 152: eql x 0
	y = 0         // line 153: mul y 0
	y += 25       // line 154: add y 25
	y *= x        // line 155: mul y x
	y += 1        // line 156: add y 1
	z *= y        // line 157: mul z y
	y = 0         // line 158: mul y 0
	y += w        // line 159: add y w
	y += 3        // line 160: add y 3
	y *= x        // line 161: mul y x
	z += y        // line 162: add z y
	w = digits[9] // line 163: inp w
	x = 0         // line 164: mul x 0
	x += z        // line 165: add x z
	x %= 26       // line 166: mod x 26
	z /= 26       // line 167: div z 26
	x += -1       // line 168: add x -1
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 169: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 170: eql x 0
	y = 0          // line 171: mul y 0
	y += 25        // line 172: add y 25
	y *= x         // line 173: mul y x
	y += 1         // line 174: add y 1
	z *= y         // line 175: mul z y
	y = 0          // line 176: mul y 0
	y += w         // line 177: add y w
	y += 10        // line 178: add y 10
	y *= x         // line 179: mul y x
	z += y         // line 180: add z y
	w = digits[10] // line 181: inp w
	x = 0          // line 182: mul x 0
	x += z         // line 183: add x z
	x %= 26        // line 184: mod x 26
	z /= 26        // line 185: div z 26
	x += -8        // line 186: add x -8
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 187: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 188: eql x 0
	y = 0          // line 189: mul y 0
	y += 25        // line 190: add y 25
	y *= x         // line 191: mul y x
	y += 1         // line 192: add y 1
	z *= y         // line 193: mul z y
	y = 0          // line 194: mul y 0
	y += w         // line 195: add y w
	y += 10        // line 196: add y 10
	y *= x         // line 197: mul y x
	z += y         // line 198: add z y
	w = digits[11] // line 199: inp w
	x = 0          // line 200: mul x 0
	x += z         // line 201: add x z
	x %= 26        // line 202: mod x 26
	z /= 26        // line 203: div z 26
	x += -5        // line 204: add x -5
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 205: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 206: eql x 0
	y = 0          // line 207: mul y 0
	y += 25        // line 208: add y 25
	y *= x         // line 209: mul y x
	y += 1         // line 210: add y 1
	z *= y         // line 211: mul z y
	y = 0          // line 212: mul y 0
	y += w         // line 213: add y w
	y += 14        // line 214: add y 14
	y *= x         // line 215: mul y x
	z += y         // line 216: add z y
	w = digits[12] // line 217: inp w
	x = 0          // line 218: mul x 0
	x += z         // line 219: add x z
	x %= 26        // line 220: mod x 26
	z /= 26        // line 221: div z 26
	x += -16       // line 222: add x -16
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 223: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 224: eql x 0
	y = 0          // line 225: mul y 0
	y += 25        // line 226: add y 25
	y *= x         // line 227: mul y x
	y += 1         // line 228: add y 1
	z *= y         // line 229: mul z y
	y = 0          // line 230: mul y 0
	y += w         // line 231: add y w
	y += 6         // line 232: add y 6
	y *= x         // line 233: mul y x
	z += y         // line 234: add z y
	w = digits[13] // line 235: inp w
	x = 0          // line 236: mul x 0
	x += z         // line 237: add x z
	x %= 26        // line 238: mod x 26
	z /= 26        // line 239: div z 26
	x += -6        // line 240: add x -6
	if x == w {
		x = 1
	} else {
		x = 0
	} // line 241: eql x w
	if x == 0 {
		x = 1
	} else {
		x = 0
	} // line 242: eql x 0
	y = 0   // line 243: mul y 0
	y += 25 // line 244: add y 25
	y *= x  // line 245: mul y x
	y += 1  // line 246: add y 1
	z *= y  // line 247: mul z y
	y = 0   // line 248: mul y 0
	y += w  // line 249: add y w
	y += 5  // line 250: add y 5
	y *= x  // line 251: mul y x
	z += y  // line 252: add z y

	return z
}

var digit int

func translate(i int, line string) string {
	prefix := line[0:5]
	suffix := strings.TrimSpace(line[5:])
	comment := fmt.Sprintf("  // line %v: %v", i+1, line)
	switch {
	case line == "inp w":
		digit++
		return fmt.Sprintf("w = digits[%v]%v", digit-1, comment)

	case line == "mul x 0":
		return "x = 0" + comment
	case line == "mul y 0":
		return "y = 0" + comment
	case line == "mul y x":
		return "y *= x" + comment
	case line == "mul z y":
		return "z *= y" + comment

	case line == "add x z":
		return "x += z" + comment
	case line == "add z y":
		return "z += y" + comment
	case prefix == "add x":
		return fmt.Sprintf("x += %v%v", suffix, comment)
	case prefix == "add y":
		return fmt.Sprintf("y += %v%v", suffix, comment)

	case line == "mod x 26":
		return "x %= 26" + comment

	case line == "div z 1":
		return comment
	case line == "div z 26":
		return "z /= 26" + comment

	case line == "eql x w":
		return "if x==w { x = 1 } else { x = 0 }" + comment
	case line == "eql x 0":
		return "if x==0 { x = 1 } else { x = 0 }" + comment
	default:
		log.Fatalf("unhandled line: %v", line)
	}
	return ""
}
