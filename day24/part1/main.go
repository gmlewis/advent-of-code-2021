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

func assert(v bool, fmtStr string, args ...interface{}) {
	if !v {
		log.Fatalf(fmtStr, args...)
	}
}

func input(digits [14]int64) int64 {
	var w, x, y, z int64

	w = digits[0] // line 1: inp w
	x = 0         // line 2: mul x 0
	x += z        // line 3: add x z
	x %= 26       // line 4: mod x 26
	// line 5: div z 1
	x += 14 // line 6: add x 14

	assert(x == 14, "x!=14 (%v)", x)
	assert(y == 0, "y!=0 (%v)", y)
	assert(z == 0, "z!=0 (%v)", z)

	logf("digits[0]=%v: x(%v)==w(%v): %v\n", digits[0], x, w, x == w)
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

	assert(x == 1, "x!=1 (%v)", x) // cannot change this
	assert(y == 0, "y!=0 (%v)", y)
	assert(z == 0, "z!=0 (%v)", z)

	y = 0   // line 9: mul y 0
	y += 25 // line 10: add y 25
	y *= x  // line 11: mul y x
	y += 1  // line 12: add y 1
	z *= y  // line 13: mul z y
	y = 0   // line 14: mul y 0
	y += w  // line 15: add y w
	y += 12 // line 16: add y 12
	y *= x  // line 17: mul y x
	z += y  // line 18: add z y
	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[1] // line 19: inp w
	x = 0         // line 20: mul x 0
	x += z        // line 21: add x z
	x %= 26       // line 22: mod x 26
	// line 23: div z 1
	x += 10 // line 24: add x 10

	assert(x == digits[0]+22, "x!=digits[0]+12+10 (%v!=%v)", x, digits[0]+12+10)
	assert(y == digits[0]+12, "y!=digits[0]+12 (%v!=%v)", y, digits[0]+12)
	assert(z == digits[0]+12, "line 24: z!=digits[0]+12 (%v!=%v)", z, digits[0]+12)

	logf("digits[1]=%v: x(%v)==w(%v): %v\n", digits[1], x, w, x == w)
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

	assert(x == 1, "x!=1 (%v)", x) // cannot change this

	y = 0   // line 27: mul y 0
	y += 25 // line 28: add y 25
	y *= x  // line 29: mul y x
	y += 1  // line 30: add y 1
	z *= y  // line 31: mul z y
	y = 0   // line 32: mul y 0
	y += w  // line 33: add y w
	y += 9  // line 34: add y 9
	y *= x  // line 35: mul y x
	z += y  // line 36: add z y
	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[2] // line 37: inp w
	x = 0         // line 38: mul x 0
	x += z        // line 39: add x z
	x %= 26       // line 40: mod x 26
	// line 41: div z 1
	x += 13 // line 42: add x 13

	assert(x == digits[0]+22, "x!=digits[0]+12+10 (%v!=%v)", x, digits[0]+12+10)
	assert(y == digits[1]+9, "y!=digits[1]+9 (%v!=%v)", y, digits[1]+9)
	assert(z == 26*(digits[0]+12)+(digits[1]+9), "line 42: z!=26*(digits[0]+12)+(digits[1]+9) (%v!=%v)", z, 26*(digits[0]+12)+(digits[1]+9))

	logf("digits[2]=%v: x(%v)==w(%v): %v\n", digits[2], x, w, x == w)
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

	assert(x == 1, "x!=1 (%v)", x) // cannot change this

	y = 0   // line 45: mul y 0
	y += 25 // line 46: add y 25
	y *= x  // line 47: mul y x
	y += 1  // line 48: add y 1
	assert(y == 26, "y!=26 (%v)", y)
	z *= y // line 49: mul z y
	assert(z == 676*(digits[0]+12)+26*(digits[1]+9), "line 49: z!=26*(26*(digits[0]+12)+(digits[1]+9)) (%v!=%v)", z, 26*(26*(digits[0]+12)+(digits[1]+9)))
	y = 0  // line 50: mul y 0
	y += w // line 51: add y w
	y += 8 // line 52: add y 8
	y *= x // line 53: mul y x
	assert(y == digits[2]+8, "y!=digits[2]+8 (%v!=%v)", y, digits[2]+8)
	z += y // line 54: add z y
	assert(z == 676*(digits[0]+12)+26*(digits[1]+9)+(digits[2]+8), "line 54: z!=26*(26*(digits[0]+12)+(digits[1]+9))+(digits[2]+8) (%v!=%v)", z, 26*(26*(digits[0]+12)+(digits[1]+9))+(digits[2]+8))
	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[3] // line 55: inp w
	x = 0         // line 56: mul x 0
	x += z        // line 57: add x z
	x %= 26       // line 58: mod x 26
	assert(x == digits[2]+8, "x!=(digits[2]+8) (%v!=%v)", x, (digits[2] + 8))
	z /= 26 // line 59: div z 26
	assert(z == 26*(digits[0]+12)+(digits[1]+9), "line 59: z!=(26*(digits[0]+12)+(digits[1]+9)) (%v!=%v)", z, (26*(digits[0]+12) + (digits[1] + 9)))
	x += -8 // line 60: add x -8
	assert(x == digits[2], "x!=digits[2] (%v!=%v)", x, digits[2])

	logf("digits[3]=%v: x(%v)==w(%v): %v\n", digits[3], x, w, x == w)
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
	if digits[2] == digits[3] {
		assert(x == 0, "x!=0 (%v)", x)
	} else {
		assert(x == 1, "x!=1 (%v)", x)
	}
	y = 0   // line 63: mul y 0
	y += 25 // line 64: add y 25
	y *= x  // line 65: mul y x
	y += 1  // line 66: add y 1
	if digits[2] == digits[3] {
		assert(y == 1, "y!=1 (%v)", y)
	} else {
		assert(y == 26, "y!=26 (%v)", y)
	}
	z *= y // line 67: mul z y
	if digits[2] == digits[3] {
		assert(z == 26*(digits[0]+12)+(digits[1]+9), "line 67: z!=(26*(digits[0]+12)+(digits[1]+9)) (%v!=%v)", z, (26*(digits[0]+12) + (digits[1] + 9)))
	} else {
		assert(z == (676*(digits[0]+12)+26*(digits[1]+9)), "line 67: z!=(26*(26*(digits[0]+12)+(digits[1]+9))) (%v!=%v)", z, (26 * (26*(digits[0]+12) + (digits[1] + 9))))
	}
	y = 0  // line 68: mul y 0
	y += w // line 69: add y w
	y += 3 // line 70: add y 3
	y *= x // line 71: mul y x
	if digits[2] == digits[3] {
		assert(y == 0, "y!=0 (%v)", y)
	} else {
		assert(y == digits[3]+3, "y!=digits[3]+3 (%v!=%v)", y, digits[3]+3)
	}
	z += y // line 72: add z y

	if digits[2] == digits[3] {
		assert(z == 26*(digits[0]+12)+(digits[1]+9), "line 72: z!=(26*(digits[0]+12)+(digits[1]+9)) (%v!=%v)", z, (26*(digits[0]+12) + (digits[1] + 9)))
	} else {
		assert(z == (676*(digits[0]+12)+26*(digits[1]+9))+(digits[3]+3), "line 72: z!=(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3) (%v!=%v)", z, (26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))
	}

	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[4] // line 73: inp w
	x = 0         // line 74: mul x 0
	x += z        // line 75: add x z
	x %= 26       // line 76: mod x 26

	if digits[2] == digits[3] {
		assert(x == (digits[1]+9), "x!=(digits[1]+9) (%v!=%v)", x, (digits[1] + 9))
	} else {
		assert(x == (digits[3]+3), "x!=(digits[3]+3) (%v!=%v)", x, (digits[3] + 3))
	}

	// line 77: div z 1
	x += 11 // line 78: add x 11

	if digits[2] == digits[3] {
		assert(x == (digits[1]+20), "x!=(digits[1]+9+11) (%v!=%v)", x, (digits[1] + 9 + 11))
	} else {
		assert(x == (digits[3]+14), "x!=(digits[3]+3+11) (%v!=%v)", x, (digits[3] + 3 + 11))
	}

	logf("digits[4]=%v: x(%v)==w(%v): %v\n", digits[4], x, w, x == w)
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

	assert(x == 1, "x!=1 (%v)", x) // cannot change this

	y = 0   // line 81: mul y 0
	y += 25 // line 82: add y 25
	y *= x  // line 83: mul y x
	y += 1  // line 84: add y 1

	assert(y == 26, "y!=26 (%v)", y)

	z *= y // line 85: mul z y

	if digits[2] == digits[3] {
		assert(z == 676*(digits[0]+12)+26*(digits[1]+9), "line 85: z!=(26*(26*(digits[0]+12)+(digits[1]+9))) (%v!=%v)", z, (26 * (26*(digits[0]+12) + (digits[1] + 9))))
	} else {
		assert(z == (17576*(digits[0]+12)+676*(digits[1]+9))+26*(digits[3]+3), "line 85: z!=(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3)) (%v!=%v)", z, (26*(26*(26*(digits[0]+12)+(digits[1]+9))) + (digits[3] + 3)))
	}

	y = 0  // line 86: mul y 0
	y += w // line 87: add y w
	y += 0 // line 88: add y 0
	y *= x // line 89: mul y x

	assert(y == digits[4], "y!=digits[4] (%v!=%v)", y, digits[4])

	z += y // line 90: add z y

	if digits[2] == digits[3] {
		assert(z == 676*(digits[0]+12)+26*(digits[1]+9)+(digits[4]), "line 90: z!=(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]) (%v!=%v)", z, (26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))
	} else {
		assert(z == (17576*(digits[0]+12)+676*(digits[1]+9))+26*(digits[3]+3)+(digits[4]), "line 90: z!=(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]) (%v!=%v)", z, (26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))
	}

	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[5] // line 91: inp w
	x = 0         // line 92: mul x 0
	x += z        // line 93: add x z
	x %= 26       // line 94: mod x 26

	assert(x == (digits[4]), "x!=(digits[4]) (%v!=%v)", x, (digits[4]))

	// line 95: div z 1
	x += 11 // line 96: add x 11

	assert(x == digits[4]+11, "x!=digits[4]+11: (%v!=%v)", x, digits[4]+11)

	logf("digits[5]=%v: x(%v)==w(%v): %v\n", digits[5], x, w, x == w)
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

	assert(x == 1, "line 98: x!=1 (%v)", x)

	y = 0   // line 99: mul y 0
	y += 25 // line 100: add y 25
	y *= x  // line 101: mul y x
	y += 1  // line 102: add y 1

	assert(y == 26, "y!=26 (%v)", y)

	z *= y // line 103: mul z y

	if digits[2] == digits[3] {
		assert(z == 17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4]), "line 103: z!=(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4])) (%v!=%v)", z, (26*(26*(26*(digits[0]+12)+(digits[1]+9))) + (digits[4])))
	} else {
		assert(z == (456976*(digits[0]+12)+17576*(digits[1]+9))+676*(digits[3]+3)+26*(digits[4]), "line 103: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4])) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3)) + (digits[4])))
	}

	y = 0   // line 104: mul y 0
	y += w  // line 105: add y w
	y += 11 // line 106: add y 11

	assert(y == digits[5]+11, "y!=digits[5]+11 (%v!=%v)", y, digits[5]+11)

	y *= x // line 107: mul y x

	assert(y == digits[5]+11, "y!=digits[5]+11 (%v!=%v)", y, digits[5]+11)

	z += y // line 108: add z y

	if digits[2] == digits[3] {
		assert(z == 17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11), "line 108: z!=(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))
	} else {
		assert(z == (456976*(digits[0]+12)+17576*(digits[1]+9))+676*(digits[3]+3)+26*(digits[4])+(digits[5]+11), "line 108: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))
	}

	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[6] // line 109: inp w
	x = 0         // line 110: mul x 0
	x += z        // line 111: add x z
	x %= 26       // line 112: mod x 26

	assert(x == (digits[5]+11), "x!=(digits[5]+11) (%v!=%v)", x, (digits[5] + 11))

	// line 113: div z 1
	x += 14 // line 114: add x 14

	assert(x == digits[5]+25, "x!=digits[5]+11+14: (%v!=%v)", x, digits[5]+11+14)

	logf("digits[6]=%v: x(%v)==w(%v): %v\n", digits[6], x, w, x == w)
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

	assert(x == 1, "line 116: x!=1 (%v)", x)

	y = 0   // line 117: mul y 0
	y += 25 // line 118: add y 25
	y *= x  // line 119: mul y x
	y += 1  // line 120: add y 1

	assert(y == 26, "y!=26 (%v)", y)

	z *= y // line 121: mul z y

	if digits[2] == digits[3] {
		assert(z == 456976*(digits[0]+12)+17576*(digits[1]+9)+676*(digits[4])+26*(digits[5]+11), "line 121: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11)) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4])) + (digits[5] + 11)))
	} else {
		assert(z == (11881376*(digits[0]+12)+456976*(digits[1]+9))+17576*(digits[3]+3)+676*(digits[4])+26*(digits[5]+11), "line 121: z!=(26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11)) (%v!=%v)", z, (26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4])) + (digits[5] + 11)))
	}

	y = 0   // line 122: mul y 0
	y += w  // line 123: add y w
	y += 10 // line 124: add y 10

	assert(y == digits[6]+10, "y!=digits[6]+10 (%v!=%v)", y, digits[6]+10)

	y *= x // line 125: mul y x

	assert(y == digits[6]+10, "y!=digits[6]+10 (%v!=%v)", y, digits[6]+10)

	z += y // line 126: add z y

	if digits[2] == digits[3] {
		assert(z == 456976*(digits[0]+12)+17576*(digits[1]+9)+676*(digits[4])+26*(digits[5]+11)+(digits[6]+10), "line 126: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[6]+10) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[6]+10))
	} else {
		assert(z == (11881376*(digits[0]+12)+456976*(digits[1]+9))+17576*(digits[3]+3)+676*(digits[4])+26*(digits[5]+11)+(digits[6]+10), "line 126: z!=(26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))+(digits[6]+10) (%v!=%v)", z, (26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))+(digits[6]+10))
	}

	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[7] // line 127: inp w
	x = 0         // line 128: mul x 0
	x += z        // line 129: add x z
	x %= 26       // line 130: mod x 26

	assert(x == (digits[6]+10), "x!=(digits[6]+10) (%v!=%v)", x, (digits[6] + 10))

	z /= 26 // line 131: div z 26

	if digits[2] == digits[3] {
		assert(z == 17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11), "line 131: z!=(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))
	} else {
		assert(z == (456976*(digits[0]+12)+17576*(digits[1]+9))+676*(digits[3]+3)+26*(digits[4])+(digits[5]+11), "line 131: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))
	}

	x += -11 // line 132: add x -11

	assert(x == digits[6]-1, "x!=digits[6]+10-11: (%v!=%v)", x, digits[6]+10-11)

	logf("digits[7]=%v: x(%v)==w(%v): %v\n", digits[7], x, w, x == w)
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

	if digits[6]-1 == digits[7] {
		assert(x == 0, "line 134: x!=0 (%v)", x)
	} else {
		assert(x == 1, "line 134: x!=1 (%v)", x)
	}

	y = 0   // line 135: mul y 0
	y += 25 // line 136: add y 25
	y *= x  // line 137: mul y x
	y += 1  // line 138: add y 1

	if digits[6]-1 == digits[7] {
		assert(y == 1, "line 138: y!=1 (%v)", y)
	} else {
		assert(y == 26, "line 138: y!=26 (%v)", y)
	}

	z *= y // line 139: mul z y

	if digits[6]-1 == digits[7] {
		if digits[2] == digits[3] {
			assert(z == 17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11), "A: line 139: z!=(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))
		} else {
			assert(z == (456976*(digits[0]+12)+17576*(digits[1]+9))+676*(digits[3]+3)+26*(digits[4])+(digits[5]+11), "B: line 139: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))
		}
	} else {
		if digits[2] == digits[3] {
			assert(z == 26*(17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11)), "C: line 139: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11)) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4])) + (digits[5] + 11)))
		} else {
			assert(z == (11881376*(digits[0]+12)+456976*(digits[1]+9))+16576*(digits[3]+3)+676*(digits[4])+26*(digits[5]+11), "D: line 139: z: (%v!=%v)", z, (26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4])) + (digits[5] + 11)))
		}
	}

	y = 0   // line 140: mul y 0
	y += w  // line 141: add y w
	y += 13 // line 142: add y 13
	assert(y == digits[7]+13, "line 142: y!=digits[7]+13 (%v!=%v)", y, digits[7]+13)
	y *= x // line 143: mul y x

	if digits[6]-1 == digits[7] {
		assert(y == 0, "line 143: y!=0 (%v)", y)
	} else {
		assert(y == digits[7]+13, "line 143: y!=digits[7]+13 (%v!=%v)", y, digits[7]+13)
	}

	z += y // line 144: add z y

	if digits[6]-1 == digits[7] {
		if digits[2] == digits[3] {
			assert(z == 17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11), "A: line 144: z!=(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))
		} else {
			assert(z == (456976*(digits[0]+12)+17576*(digits[1]+9))+676*(digits[3]+3)+26*(digits[4])+(digits[5]+11), "B: line 144: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))
		}
	} else {
		if digits[2] == digits[3] {
			assert(z == 26*(17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11))+(digits[7]+13), "C: line 144: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[7]+13) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[7]+13))
		} else {
			assert(z == (11881376*(digits[0]+12)+456976*(digits[1]+9))+16576*(digits[3]+3)+676*(digits[4])+26*(digits[5]+11)+(digits[7]+13), "D: line 144: z: (%v!=%v)", z, (26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))+(digits[7]+13))
		}
	}

	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[8] // line 145: inp w
	x = 0         // line 146: mul x 0
	x += z        // line 147: add x z
	x %= 26       // line 148: mod x 26

	if digits[6]-1 == digits[7] {
		assert(x == (digits[5]+11), "A: line 148: x!=(digits[5]+11) (%v!=%v)", x, (digits[5] + 11))
	} else {
		assert(x == (digits[7]+13), "B: line 148: x!=(digits[7]+13) (%v!=%v)", x, (digits[7] + 13))
	}

	// line 149: div z 1
	x += 14 // line 150: add x 14

	if digits[6]-1 == digits[7] {
		assert(x == (digits[5]+25), "A: line 150: x!=(digits[5]+11+14) (%v!=%v)", x, (digits[5] + 11 + 14))
	} else {
		assert(x == (digits[7]+27), "B: line 150: x!=(digits[7]+13+14) (%v!=%v)", x, (digits[7] + 13 + 14))
	}

	logf("digits[8]=%v: x(%v)==w(%v): %v\n", digits[8], x, w, x == w)
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

	assert(x == 1, "x!=1 (%v)", x) // cannot change this

	y = 0   // line 153: mul y 0
	y += 25 // line 154: add y 25
	y *= x  // line 155: mul y x
	y += 1  // line 156: add y 1

	assert(y == 26, "y!=26 (%v)", y)

	z *= y // line 157: mul z y

	if digits[6]-1 == digits[7] {
		if digits[2] == digits[3] {
			assert(z == (26*(17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11))), "A: line 157: z: (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4])) + (digits[5] + 11)))
		} else {
			assert(z == (26*((456976*(digits[0]+12)+17576*(digits[1]+9))+676*(digits[3]+3)+26*(digits[4])+(digits[5]+11))), "B: line 157: z: (%v!=%v)", z, (26 * ((26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3)) + (digits[4])) + (digits[5] + 11))))
		}
	} else {
		if digits[2] == digits[3] {
			assert(z == (26*(26*(17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11))+(digits[7]+13))), "C: line 157: z: (%v!=%v)", z, (26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11)) + (digits[7] + 13)))
		} else {
			assert(z == (26*(11881376*(digits[0]+12)+456976*(digits[1]+9))+16576*(digits[3]+3)+676*(digits[4])+26*(digits[5]+11)+(digits[7]+13)), "D: line 157: z: (%v!=%v)", z, (26*(26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11)) + (digits[7] + 13)))
		}
	}

	y = 0  // line 158: mul y 0
	y += w // line 159: add y w
	y += 3 // line 160: add y 3
	assert(y == digits[8]+3, "y!=digits[8]+3 (%v!=%v)", y, digits[8]+3)
	y *= x // line 161: mul y x
	assert(y == digits[8]+3, "y!=digits[8]+3 (%v!=%v)", y, digits[8]+3)

	z += y // line 162: add z y

	if digits[6]-1 == digits[7] {
		if digits[2] == digits[3] {
			assert(z == (26*(17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11)))+(digits[8]+3), "A: line 162: z: (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[8]+3))
		} else {
			assert(z == (26*((456976*(digits[0]+12)+17576*(digits[1]+9))+676*(digits[3]+3)+26*(digits[4])+(digits[5]+11)))+(digits[8]+3), "B: line 162: z: (%v!=%v)", z, (26*((26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11)))+(digits[8]+3))
		}
	} else {
		if digits[2] == digits[3] {
			assert(z == (26*(26*(17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11))+(digits[7]+13)))+(digits[8]+3), "C: line 162: z: (%v!=%v)", z, (26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[7]+13))+(digits[8]+3))
		} else {
			assert(z == (26*(11881376*(digits[0]+12)+456976*(digits[1]+9))+16576*(digits[3]+3)+676*(digits[4])+26*(digits[5]+11)+(digits[7]+13))+(digits[8]+3), "D: line 162: z: (%v!=%v)", z, (26*(26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))+(digits[7]+13))+(digits[8]+3))
		}
	}

	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[9] // line 163: inp w
	x = 0         // line 164: mul x 0
	x += z        // line 165: add x z
	x %= 26       // line 166: mod x 26

	assert(x == (digits[8]+3), "line 166: x: (%v!=%v)", x, (digits[8] + 3))

	z /= 26 // line 167: div z 26

	if digits[6]-1 == digits[7] {
		if digits[2] == digits[3] {
			assert(z == 17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11), "A: line 167: z!=(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))
		} else {
			assert(z == (456976*(digits[0]+12)+17576*(digits[1]+9))+676*(digits[3]+3)+26*(digits[4])+(digits[5]+11), "B: line 167: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))
		}
	} else {
		if digits[2] == digits[3] {
			assert(z == 26*(17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11))+(digits[7]+13), "C: line 167: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[7]+13) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[7]+13))
		} else {
			assert(z == (11881376*(digits[0]+12)+456976*(digits[1]+9))+16576*(digits[3]+3)+676*(digits[4])+26*(digits[5]+11)+(digits[7]+13), "D: line 167: z: (%v!=%v)", z, (26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))+(digits[7]+13))
		}
	}

	x += -1 // line 168: add x -1

	assert(x == digits[8]+2, "x != digits[8]+3-1: (%v!=%v)", x, digits[8]+3-1)

	logf("digits[9]=%v: x(%v)==w(%v): %v\n", digits[9], x, w, x == w)
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

	if digits[8]+2 == digits[9] {
		assert(x == 0, "x!=0: (%v)", x)
	} else {
		assert(x == 1, "x!=1: (%v)", x)
	}

	y = 0   // line 171: mul y 0
	y += 25 // line 172: add y 25
	y *= x  // line 173: mul y x
	y += 1  // line 174: add y 1
	if digits[8]+2 == digits[9] {
		assert(y == 1, "y!=1: (%v)", y)
	} else {
		assert(y == 26, "y!=26: (%v)", y)
	}
	z *= y // line 175: mul z y

	if digits[8]+2 == digits[9] {
		if digits[6]-1 == digits[7] {
			if digits[2] == digits[3] {
				assert(z == 17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11), "A: line 167: z!=(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))
			} else {
				assert(z == (456976*(digits[0]+12)+17576*(digits[1]+9))+676*(digits[3]+3)+26*(digits[4])+(digits[5]+11), "B: line 167: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))
			}
		} else {
			if digits[2] == digits[3] {
				assert(z == 26*(17576*(digits[0]+12)+676*(digits[1]+9)+26*(digits[4])+(digits[5]+11))+(digits[7]+13), "C: line 167: z!=(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[7]+13) (%v!=%v)", z, (26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))+(digits[5]+11))+(digits[7]+13))
			} else {
				assert(z == (11881376*(digits[0]+12)+456976*(digits[1]+9))+16576*(digits[3]+3)+676*(digits[4])+26*(digits[5]+11)+(digits[7]+13), "D: line 167: z: (%v!=%v)", z, (26*(26*(26*(26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[3]+3))+(digits[4]))+(digits[5]+11))+(digits[7]+13))
			}
		}
	}

	y = 0   // line 176: mul y 0
	y += w  // line 177: add y w
	y += 10 // line 178: add y 10
	assert(y == digits[9]+10, "y!=digits[9]+10: (%v)", digits[9]+10)
	y *= x // line 179: mul y x
	if digits[8]+2 == digits[9] {
		assert(y == 0, "y!=0: (%v)", y)
	}
	z += y // line 180: add z y
	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[10] // line 181: inp w
	x = 0          // line 182: mul x 0
	x += z         // line 183: add x z
	x %= 26        // line 184: mod x 26

	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] {
		assert(x == digits[5]+11, "line 184: x != digits[5]+11: (%v!=%v)", x, digits[5]+11)
	}

	z /= 26 // line 185: div z 26
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[2] == digits[3] {
		assert(z == 676*(digits[0]+12)+26*(digits[1]+9)+(digits[4]), "line 185: z: (%v!=%v)", z, (26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))
	}
	x += -8 // line 186: add x -8

	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] {
		assert(x == digits[5]+3, "line 186: x != digits[5]+11-8: (%v!=%v)", x, digits[5]+11-8)
	}

	logf("digits[10]=%v: x(%v)==w(%v): %v\n", digits[10], x, w, x == w)
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

	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[5]+3 == digits[10] {
		assert(x == 0, "line 188: x!=0: (%v)", x)
	}

	y = 0   // line 189: mul y 0
	y += 25 // line 190: add y 25
	y *= x  // line 191: mul y x
	y += 1  // line 192: add y 1
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[5]+3 == digits[10] {
		assert(y == 1, "line 192: y!=1: (%v)", y)
	}
	z *= y // line 193: mul z y
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[2] == digits[3] && digits[5]+3 == digits[10] {
		assert(z == 676*(digits[0]+12)+26*(digits[1]+9)+(digits[4]), "line 193: z: (%v!=%v)", z, (26*(26*(digits[0]+12)+(digits[1]+9)))+(digits[4]))
	}
	y = 0   // line 194: mul y 0
	y += w  // line 195: add y w
	y += 10 // line 196: add y 10
	y *= x  // line 197: mul y x
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[5]+3 == digits[10] {
		assert(y == 0, "line 197: y!=0: (%v)", y)
	}
	z += y // line 198: add z y
	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[11] // line 199: inp w
	x = 0          // line 200: mul x 0
	x += z         // line 201: add x z
	x %= 26        // line 202: mod x 26
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[5]+3 == digits[10] {
		assert(x == (digits[4]), "line 202: x: (%v!=%v)", x, digits[4])
	}
	z /= 26 // line 203: div z 26
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[2] == digits[3] && digits[5]+3 == digits[10] {
		assert(z == 26*(digits[0]+12)+(digits[1]+9), "line 203: z: (%v!=%v)", z, (26*(digits[0]+12) + (digits[1] + 9)))
	}
	x += -5 // line 204: add x -5
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[5]+3 == digits[10] {
		assert(x == (digits[4]-5), "line 204: x: (%v!=%v)", x, digits[4]-5)
	}

	logf("digits[11]=%v: x(%v)==w(%v): %v\n", digits[11], x, w, x == w)
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
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[5]+3 == digits[10] && digits[4]-5 == digits[11] {
		assert(x == 0, "line 206: x!=0: (%v)", x)
	}
	y = 0   // line 207: mul y 0
	y += 25 // line 208: add y 25
	y *= x  // line 209: mul y x
	y += 1  // line 210: add y 1
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[5]+3 == digits[10] && digits[4]-5 == digits[11] {
		assert(y == 1, "line 210: y!=1: (%v)", y)
	}
	z *= y // line 211: mul z y
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[2] == digits[3] && digits[5]+3 == digits[10] && digits[4]-5 == digits[11] {
		assert(z == 26*(digits[0]+12)+(digits[1]+9), "line 211: z: (%v!=%v)", z, (26*(digits[0]+12) + (digits[1] + 9)))
	}
	y = 0   // line 212: mul y 0
	y += w  // line 213: add y w
	y += 14 // line 214: add y 14
	assert(y == digits[11]+14, "line 214: y!=digits[11]+14: (%v!=%v)", y, digits[11]+14)
	y *= x // line 215: mul y x
	if digits[8]+2 == digits[9] && digits[6]-1 == digits[7] && digits[5]+3 == digits[10] && digits[4]-5 == digits[11] {
		assert(y == 0, "line 215: y!=0: (%v)", y)
	}
	z += y // line 216: add z y
	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[12] // line 217: inp w
	x = 0          // line 218: mul x 0
	x += z         // line 219: add x z
	x %= 26        // line 220: mod x 26
	z /= 26        // line 221: div z 26
	x += -16       // line 222: add x -16

	assert(x == 4 || x == digits[11]-2, "x!=4 && x!=digits[11]-2: (%v!=%v)", x, digits[11]-2)

	logf("digits[12]=%v: x(%v)==w(%v): %v\n", digits[12], x, w, x == w)
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
	y = 0   // line 225: mul y 0
	y += 25 // line 226: add y 25
	y *= x  // line 227: mul y x
	y += 1  // line 228: add y 1
	z *= y  // line 229: mul z y
	y = 0   // line 230: mul y 0
	y += w  // line 231: add y w
	y += 6  // line 232: add y 6
	y *= x  // line 233: mul y x
	z += y  // line 234: add z y
	logf("x=%v, y=%v, z=%v", x, y, z)
	w = digits[13] // line 235: inp w
	x = 0          // line 236: mul x 0
	x += z         // line 237: add x z
	x %= 26        // line 238: mod x 26
	z /= 26        // line 239: div z 26
	x += -6        // line 240: add x -6
	logf("digits[13]=%v: x(%v)==w(%v): %v\n", digits[13], x, w, x == w)
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
	logf("FINAL: x=%v, y=%v, z=%v", x, y, z)

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
