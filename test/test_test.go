package test

import (
	"fmt"
	"testing"
)

var printf = fmt.Printf

func TestRunner(t *testing.T) {
	process := func(string) { printf("output") }
	Runner(t, "input", "output", process, &printf)
}
