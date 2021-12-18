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
	final := sum(nums)

	printf("Solution: %v\n", final)
}

type nodeT struct {
	left  *nodeT
	right *nodeT

	literal int
}

func sum(nodes []*nodeT) *nodeT {
	result := nodes[0]
	for i, node := range nodes[1:] {
		v := result.add(node)
		logf("after sum #%v: %v", i+1, v)
		result = v.reduce()
		logf("after sum+reduce #%v: %v", i+1, result)
	}
	return result
}

func (n *nodeT) reduce() *nodeT {
	for {
		if r, addLeft, addRight := n.explode(0); addLeft != nil || addRight != nil {
			n = r
			logf("after explode: %v", n)
			continue
		}
		if r, ok := n.split(); ok {
			n = r
			logf("after split: %v", n)
			continue
		}
		break
	}
	return n
}

func (n *nodeT) split() (*nodeT, bool) {
	if n.left != nil {
		left, ok := n.left.split()
		if ok {
			n.left = left
			return n, true
		}
		right, ok := n.right.split()
		if ok {
			n.right = right
			return n, true
		}
	}
	if n.literal >= 10 {
		left := &nodeT{literal: n.literal / 2}
		right := &nodeT{literal: (n.literal + 1) / 2}
		return &nodeT{left: left, right: right}, true
	}
	return n, false
}

func (n *nodeT) explode(depth int) (result, addLeft, addRight *nodeT) {
	if depth == 3 && n.left != nil {
		if n.left.left != nil { // left node explodes
			logf("left node explodes: %v, n.left=%v, n.right.left=%v, n.left.right=%v", n, n.left, n.right.left, n.left.right)
			if n.right.left != nil {
				return &nodeT{left: &nodeT{literal: 0}, right: &nodeT{left: &nodeT{literal: n.right.left.literal + n.left.right.literal}, right: n.right.right}}, n.left.left, nil
			}
			return &nodeT{left: &nodeT{literal: 0}, right: &nodeT{literal: n.left.right.literal + n.right.literal}}, n.left.left, nil
		} else if n.right.left != nil { // right node explodes
			logf("left right explodes: %v", n)
			return &nodeT{left: &nodeT{literal: n.left.literal + n.right.left.literal}, right: &nodeT{literal: 0}}, nil, n.right.right
		}
	}
	if n.left != nil {
		left, addLeft, addRight := n.left.explode(depth + 1)
		if addLeft != nil {
			n.left = left
			logf("NOOP parent: left exploded, addLeft=%v, this.left=%v, addRight=%v", addLeft, n.left, addRight)
			return n, addLeft, nil
		}
		if addRight != nil {
			n.left = left
			if addRight.literal != 0 {
				n.right = n.right.addToRight(addRight.literal)
				logf("parent: left exploded, addRight=%v, this.right=%v", addRight, n.right)
			}
			return n, nil, &nodeT{literal: 0} // made change, no further changes allowed
		}
		right, addLeft, addRight := n.right.explode(depth + 1)
		if addLeft != nil {
			n.right = right
			if addLeft.literal != 0 {
				n.left = n.left.addToLeft(addLeft.literal)
				logf("parent: right exploded, addLeft=%v, this.left=%v, addRight=%v", addLeft, n.left, addRight)
			}
			return n, &nodeT{literal: 0}, nil // made change, no further changes allowed
		}
		if addRight != nil {
			n.right = right
			logf("NOOP parent: right exploded, addRight=%v, this.right=%v", addRight, n.right)
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
		if len(s) > 1 && s[1] >= '0' && s[1] <= '9' {
			return &nodeT{literal: 10*int(s[0]-'0') + int(s[1]-'0')}, s[2:]
		}
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

func (n *nodeT) String() string {
	if n == nil {
		return ""
	}
	if n.left == nil {
		return fmt.Sprintf("%v", n.literal)
	}
	return fmt.Sprintf("[%v,%v]", n.left, n.right)
}
