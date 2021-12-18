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

	var max int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			n := sum([]*nodeT{nums[i], nums[j]}).magnitude()
			if n > max {
				max = n
			}
		}
	}

	printf("Solution: %v\n", max)
}

type nodeT struct {
	left  *nodeT
	right *nodeT

	literal int
}

func (n *nodeT) magnitude() int {
	if n.left != nil {
		left := n.left.magnitude()
		right := n.right.magnitude()
		return 3*left + 2*right
	}
	return n.literal
}

func sum(nodes []*nodeT) *nodeT {
	result := nodes[0]
	for _, node := range nodes[1:] {
		v := result.add(node)
		result = v.reduce()
	}
	return result
}

func (n *nodeT) reduce() *nodeT {
	for {
		if r, addLeft, addRight := n.explode(0); addLeft != nil || addRight != nil {
			n = r
			continue
		}
		if r, ok := n.split(); ok {
			n = r
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
			if n.right.left != nil {
				return &nodeT{left: &nodeT{literal: 0}, right: &nodeT{left: &nodeT{literal: n.right.left.literal + n.left.right.literal}, right: n.right.right}}, n.left.left, nil
			}
			return &nodeT{left: &nodeT{literal: 0}, right: &nodeT{literal: n.left.right.literal + n.right.literal}}, n.left.left, nil
		} else if n.right.left != nil { // right node explodes
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
			}
			return n, nil, &nodeT{literal: 0} // made change, no further changes allowed
		}
		right, addLeft, addRight := n.right.explode(depth + 1)
		if addLeft != nil {
			n.right = right
			if addLeft.literal != 0 {
				n.left = n.left.addToLeft(addLeft.literal)
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
	return &nodeT{left: n.copy(), right: plus.copy()}
}

func (n *nodeT) copy() *nodeT {
	if n.left != nil {
		return &nodeT{left: n.left.copy(), right: n.right.copy()}
	}
	return &nodeT{literal: n.literal}
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
