#!/bin/bash -e
mkdir -p $@/part1 $@/part2

if [ ! -x ./cmd/get-puzzle/get-puzzle ]; then
    pushd cmd/get-puzzle && go build && popd
fi

# Note that the environment variable `AOC_COOKIES` must be
# set before calling the next command.
./cmd/get-puzzle/get-puzzle $@

cp day01/run-all.sh $@
cp day01/part1/run-go.sh $@/part1
cp day01/part2/run-go.sh $@/part2
