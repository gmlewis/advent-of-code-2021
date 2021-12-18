// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

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
	nums := Map(lines, func(line string) *nodeT { n, _ := parse(line); return n })

	printf("Solution: %v\n", len(nums))
}

type nodeT struct {
	left  *nodeT
	right *nodeT

	literal int
}

func (n *nodeT) explode(depth int) (result, addLeft, addRight *nodeT) {
	logf("explode(%v): %v", depth, n.prettyPrint())
	if depth == 3 && n.left != nil {
		if n.left.left != nil { // left node explodes
			logf("left node explodes: %v", n.prettyPrint())
			return &nodeT{left: &nodeT{literal: 0}, right: &nodeT{literal: n.left.right.literal + n.right.literal}}, n.left.left, nil
		} else if n.right.left != nil { // right node explodes
			logf("right node explodes: %v", n.prettyPrint())
			return &nodeT{left: &nodeT{literal: n.left.literal + n.right.left.literal}, right: &nodeT{literal: 0}}, nil, n.right.right
		}
	}
	if n.left != nil {
		left, addLeft, addRight := n.left.explode(depth + 1)
		if addLeft != nil {
			n.left = left
			return n, addLeft, nil
		}
		if addRight != nil {
			n.left = left
			if addRight.literal != 0 {
				n.right = n.right.addToRight(addRight.literal)
				logf("parent: left exploded, addRight=%v, this.right=%v", addRight.prettyPrint(), n.right.prettyPrint())
			}
			return n, nil, &nodeT{literal: 0} // made change, no further changes allowed
		}
		right, addLeft, addRight := n.right.explode(depth + 1)
		if addLeft != nil {
			n.right = right
			if addLeft.literal != 0 {
				n.left = n.left.addToLeft(addLeft.literal)
				logf("parent: right exploded, addLeft=%v, this.left=%v", addLeft.prettyPrint(), n.left.prettyPrint())
			}
			return n, &nodeT{literal: 0}, nil // made change, no further changes allowed
		}
		if addRight != nil {
			n.right = right
			return n, nil, addRight
		}
	}
	return n, nil, nil
}

func (n *nodeT) addToRight(v int) *nodeT {
	if n.left != nil {
		n.left = n.left.addToRight(v)
		return n
	}
	n.literal += v
	return n
}

func (n *nodeT) addToLeft(v int) *nodeT {
	if n.right != nil {
		n.right = n.right.addToLeft(v)
		return n
	}
	n.literal += v
	return n
}

func (n *nodeT) add(plus *nodeT) *nodeT {
	return &nodeT{left: n, right: plus}
}

func parse(s string) (*nodeT, string) {
	if s[0] >= '0' && s[0] <= '9' {
		return &nodeT{literal: int(s[0] - '0')}, s[1:]
	}
	if s[0] != '[' {
		log.Fatalf("expected '[', got: %q", s)
	}
	n := &nodeT{}
	n.left, s = parse(s[1:])
	if s[0] != ',' {
		log.Fatalf("expected ',', got: %q", s)
	}
	n.right, s = parse(s[1:])
	if s[0] != ']' {
		log.Fatalf("expected ']', got: %q", s)
	}
	return n, s[1:]
}

func (n *nodeT) prettyPrint() string {
	if n.left == nil {
		return fmt.Sprintf("%v", n.literal)
	}
	left := n.left.prettyPrint()
	right := n.right.prettyPrint()
	return fmt.Sprintf("[%v,%v]", left, right)
}
