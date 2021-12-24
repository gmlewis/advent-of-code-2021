# advent-of-code-2021

Here are my solutions for [Advent of Code 2021](https://adventofcode.com/2021).

This year, I chose to write my solutions using Go 1.18 with generics
(by building the Go compiler from the latest master branch):

```
$ go version
go version devel go1.18-deb988a286 Fri Dec 3 18:09:19 2021 +0000 linux/amd64
```

## Benchmarks

Benchmarks were run on an Intel i7 laptop running Linux Mint 19.3.

| Puzzle      | Benchmark          | Iters   | ns/op             |
|-------------|--------------------|     --: |               --: |
| day01/part1 | BenchmarkInput-4   |    8048 |      144071 ns/op |
| day01/part2 | BenchmarkInput-4   |    6345 |      160099 ns/op |
|-------------|--------------------|---------|-------------------|
| day02/part1 | BenchmarkInput-4   |    7518 |      142567 ns/op |
| day02/part2 | BenchmarkInput-4   |    7508 |      140629 ns/op |
|-------------|--------------------|---------|-------------------|
| day03/part1 | BenchmarkInput-4   |   24511 |       43900 ns/op |
| day03/part2 | BenchmarkInput-4   |    5559 |      191525 ns/op |
|-------------|--------------------|---------|-------------------|
| day04/part1 | BenchmarkInput-4   |     270 |     3782730 ns/op |
| day04/part2 | BenchmarkInput-4   |     310 |     3772155 ns/op |
|-------------|--------------------|---------|-------------------|
| day05/part1 | BenchmarkInput-4   |      26 |    44208473 ns/op |
| day05/part2 | BenchmarkInput-4   |      13 |    88025595 ns/op |
|-------------|--------------------|---------|-------------------|
| day06/part1 | BenchmarkInput-4   |   14320 |       77127 ns/op |
| day06/part2 | BenchmarkInput-4   |    5192 |      212946 ns/op |
|-------------|--------------------|---------|-------------------|
| day07/part1 | BenchmarkInput-4   |     340 |     3504080 ns/op |
| day07/part2 | BenchmarkInput-4   |     288 |     4131633 ns/op |
|-------------|--------------------|---------|-------------------|
| day08/part1 | BenchmarkInput-4   |   11214 |      107337 ns/op |
| day08/part2 | BenchmarkInput-4   |     630 |     1870067 ns/op |
|-------------|--------------------|---------|-------------------|
| day09/part1 | BenchmarkInput-4   |     595 |     2000520 ns/op |
| day09/part2 | BenchmarkInput-4   |     297 |     4048426 ns/op |
|-------------|--------------------|---------|-------------------|
| day10/part1 | BenchmarkInput-4   |    8643 |      119493 ns/op |
| day10/part2 | BenchmarkInput-4   |    8388 |      127750 ns/op |
|-------------|--------------------|---------|-------------------|
| day11/part1 | BenchmarkInput-4   |     874 |     1359490 ns/op |
| day11/part2 | BenchmarkInput-4   |     324 |     3681322 ns/op |
|-------------|--------------------|---------|-------------------|
| day12/part1 | BenchmarkInput-4   |      27 |    39919589 ns/op |
| day12/part2 | BenchmarkInput-4   |       1 |  1396313129 ns/op |
|-------------|--------------------|---------|-------------------|
| day13/part1 | BenchmarkInput-4   |    3220 |      344567 ns/op |
| day13/part2 | BenchmarkInput-4   |    1161 |      865072 ns/op |
|-------------|--------------------|---------|-------------------|
| day14/part1 | BenchmarkInput-4   |     376 |     3065596 ns/op |
| day14/part2 | BenchmarkInput-4   |    1483 |      740865 ns/op |
|-------------|--------------------|---------|-------------------|
| day15/part1 | BenchmarkInput-4   |      32 |    37535837 ns/op |
| day15/part2 | BenchmarkInput-4   |       1 |  1414612059 ns/op |
|-------------|--------------------|---------|-------------------|
| day16/part1 | BenchmarkInput-4   |    8124 |      134785 ns/op |
| day16/part2 | BenchmarkInput-4   |    7513 |      136187 ns/op |
|-------------|--------------------|---------|-------------------|
| day17/part1 | BenchmarkInput-4   |    1644 |      723567 ns/op |
| day17/part2 | BenchmarkInput-4   |    1454 |      816308 ns/op |
|-------------|--------------------|---------|-------------------|
| day18/part1 | BenchmarkInput-4   |     390 |     3518750 ns/op |
| day18/part2 | BenchmarkInput-4   |      81 |    13735733 ns/op |
|-------------|--------------------|---------|-------------------|
| day19/part1 | BenchmarkInput-4   |       8 |   139914205 ns/op |
| day19/part2 | BenchmarkInput-4   |       8 |   141170683 ns/op |
|-------------|--------------------|---------|-------------------|
| day20/part1 | BenchmarkInput-4   |     146 |     8597374 ns/op |
| day20/part2 | BenchmarkInput-4   |       3 |   435822786 ns/op |
|-------------|--------------------|---------|-------------------|
| day21/part1 | BenchmarkInput-4   |   21853 |       59397 ns/op |
| day21/part2 | BenchmarkInput-4   |       2 |   854468106 ns/op |
|-------------|--------------------|---------|-------------------|
| day22/part1 | BenchmarkInput-4   |       7 |   161074391 ns/op |
| day22/part2 | BenchmarkInput-4   |       1 | 91770973916 ns/op |
|-------------|--------------------|---------|-------------------|
| day23/part1 | BenchmarkInput-4   |       1 | 19647640120 ns/op |
| day23/part2 | BenchmarkInput-4   |       1 |110287766926 ns/op |
|-------------|--------------------|---------|-------------------|

----------------------------------------------------------------------

# License

Copyright 2021 Glenn M. Lewis. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
