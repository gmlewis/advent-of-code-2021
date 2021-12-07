#!/bin/bash -ex
mkdir -p $@/part1
mkdir -p $@/part2
touch $@/example1.txt $@/input.txt
cp day01/run-all.sh $@
cp day01/part1/run-go.sh day01/part1/main_test.go $@/part1
cp day01/part2/run-go.sh day01/part2/main_test.go $@/part2
MAIN=<<EOF
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

echo ${MAIN} > $@/part1/main.go
echo ${MAIN} > $@/part2/main.go
