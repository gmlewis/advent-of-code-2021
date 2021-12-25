// -*- compile-command: "go run main.go ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
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

	code, f := translate(0, 0, lines, func([14]int, regT) regT { return regT{} })
	logf("%v lines of code", len(code))

	max := [14]int{3, 9, 9, 9, 9, 6, 9, 8, 7, 9, 9, 4, 2, 9}
	min := [14]int{1, 8, 1, 1, 6, 1, 2, 1, 1, 3, 4, 1, 1, 7}

	printf("Part 1 Solution: %v f=%v\n", max, f(max, regT{}))
	printf("Part 2 Solution: %v f=%v\n", min, f(max, regT{}))
}

type regT struct {
	w, x, y, z int
}

func translate(digit, lineCount int, lines []string, f func([14]int, regT) regT) ([]string, func([14]int, regT) regT) {
	var result []string

	for j, line := range lines {
		i := lineCount + j
		prefix := line[0:5]
		suffix := strings.TrimSpace(line[5:])
		num, numErr := strconv.Atoi(suffix)
		comment := fmt.Sprintf("  // line %v: %v", i+1, line)
		f2 := f

		switch {
		case line == "inp w":
			digit++
			line = fmt.Sprintf("w = digits[%v]%v", digit-1, comment)
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.w = digits[digit-1]
				// logf("r.w=%v%v", r.w, comment)
				return r
			}

		case line == "mul x 0":
			line = "x = 0" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.x = 0
				// logf("r.x=%v%v", r.x, comment)
				return r
			}

		case line == "mul y 0":
			line = "y = 0" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.y = 0
				// logf("r.y=%v%v", r.y, comment)
				return r
			}

		case line == "mul y x":
			line = "y *= x" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.y *= r.x
				// logf("r.y *= r.x(%v) = %v%v", r.x, r.y, comment)
				return r
			}

		case line == "mul z y":
			line = "z *= y" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.z *= r.y
				// logf("r.z *= r.y(%v) = %v%v", r.y, r.z, comment)
				return r
			}

		case line == "add x z":
			line = "x += z" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.x += r.z
				// logf("r.x += r.z(%v) = %v%v", r.z, r.x, comment)
				return r
			}

		case line == "add z y":
			line = "z += y" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.z += r.y
				// logf("r.z += r.y(%v) = %v%v", r.y, r.z, comment)
				return r
			}

		case line == "add y w":
			line = "y += w" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.y += r.w
				// logf("r.y += r.w(%v) = %v%v", r.w, r.y, comment)
				return r
			}

		case prefix == "add x":
			if numErr != nil {
				log.Fatalf("unhandled line: %v", line)
			}
			line = fmt.Sprintf("x += %v%v", suffix, comment)
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.x += num
				// logf("r.x += %v = %v%v", num, r.x, comment)
				return r
			}

		case prefix == "add y":
			if numErr != nil {
				log.Fatalf("unhandled line: %v", line)
			}
			line = fmt.Sprintf("y += %v%v", suffix, comment)
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.y += num
				// logf("r.y += %v = %v%v", num, r.y, comment)
				return r
			}

		case line == "mod x 26":
			line = "x %= 26" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.x %= 26
				// logf("r.x %%= 26 = %v%v", r.x, comment)
				return r
			}

		case line == "div z 1":
			line = comment

		case line == "div z 26":
			line = "z /= 26" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				r.z /= 26
				// logf("r.z /= 26 = %v%v", r.z, comment)
				return r
			}

		case line == "eql x w":
			line = "if x==w { x = 1 } else { x = 0 }" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				if r.x == r.w {
					r.x = 1
					// logf("x==w: r.x = 1%v", comment)
				} else {
					r.x = 0
					// logf("x!=w: r.x = 0%v", comment)
				}
				return r
			}

		case line == "eql x 0":
			line = "if x==0 { x = 1 } else { x = 0 }" + comment
			f2 = func(digits [14]int, regs regT) regT {
				r := f(digits, regs)
				if r.x == 0 {
					r.x = 1
					// logf("x==0: r.x = 1%v", comment)
				} else {
					r.x = 0
					// logf("x!=0: r.x = 0%v", comment)
				}
				return r
			}

		default:
			log.Fatalf("unhandled line: %v", line)
		}

		result = append(result, line)
		code, nf := translate(digit, i+1, lines[j+1:], f2)
		return append(result, code...), nf
	}

	return result, f
}
