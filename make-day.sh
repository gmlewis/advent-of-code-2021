#!/bin/bash -e
mkdir -p $@/part1
mkdir -p $@/part2
touch $@/example1.txt $@/input.txt
cp day01/run-all.sh $@
cp day01/part1/run-go.sh $@/part1
cp day01/part2/run-go.sh $@/part2
MAIN=$(cat <<EOF
// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

var printf = fmt.Printf

func main() {
	flag.Parse()

	enum.Each(flag.Args(), process)
}

func process(filename string) {
	log.Printf("Processing %v ...", filename)
	buf := must.ReadFile(filename)

	printf("Solution: %v\n", len(buf))
}
EOF
)

TEST=$(cat <<EOF
package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 0\n"
	test.Runner(t, example1, want, process, &printf)
}

var example1 = \`

\`
EOF
)

echo "${MAIN}" > $@/part1/main.go
echo "${TEST}" > $@/part1/main_test.go
echo "${MAIN}" > $@/part2/main.go
echo "${TEST}" > $@/part2/main_test.go
