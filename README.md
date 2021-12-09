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

| Puzzle      | Benchmark          | Iters   | ns/op           |
|-------------|--------------------|     --: |             --: |
| day01/part1 | BenchmarkExample-4 |  176708 |      6524 ns/op |
| day01/part1 | BenchmarkInput-4   |    8048 |    144071 ns/op |
| day01/part2 | BenchmarkExample-4 |  183540 |      6406 ns/op |
| day01/part2 | BenchmarkInput-4   |    6345 |    160099 ns/op |
|-------------|--------------------|---------|-----------------|
| day02/part1 | BenchmarkExample-4 |  188508 |      6339 ns/op |
| day02/part1 | BenchmarkInput-4   |    7518 |    142567 ns/op |
| day02/part2 | BenchmarkExample-4 |  169826 |      6375 ns/op |
| day02/part2 | BenchmarkInput-4   |    7508 |    140629 ns/op |
|-------------|--------------------|---------|-----------------|
| day03/part1 | BenchmarkExample-4 |  193821 |      5864 ns/op |
| day03/part1 | BenchmarkInput-4   |   24511 |     43900 ns/op |
| day03/part2 | BenchmarkExample-4 |  124622 |      9356 ns/op |
| day03/part2 | BenchmarkInput-4   |    5559 |    191525 ns/op |
|-------------|--------------------|---------|-----------------|
| day04/part1 | BenchmarkExample-4 |   15457 |     77290 ns/op |
| day04/part1 | BenchmarkInput-4   |     270 |   3782730 ns/op |
| day04/part2 | BenchmarkExample-4 |   15550 |     76737 ns/op |
| day04/part2 | BenchmarkInput-4   |     310 |   3772155 ns/op |
|-------------|--------------------|---------|-----------------|
| day05/part1 | BenchmarkExample-4 |   60306 |     19457 ns/op |
| day05/part1 | BenchmarkInput-4   |      26 |  44208473 ns/op |
| day05/part2 | BenchmarkExample-4 |   42327 |     28219 ns/op |
| day05/part2 | BenchmarkInput-4   |      13 |  88025595 ns/op |
|-------------|--------------------|---------|-----------------|
| day06/part1 | BenchmarkExample-4 |   19146 |     62354 ns/op |
| day06/part1 | BenchmarkInput-4   |   14320 |     77127 ns/op |
| day06/part2 | BenchmarkExample-4 |    5108 |    197662 ns/op |
| day06/part2 | BenchmarkInput-4   |    5192 |    212946 ns/op |
|-------------|--------------------|---------|-----------------|
| day07/part1 | BenchmarkExample-4 |  184621 |      6264 ns/op |
| day07/part1 | BenchmarkInput-4   |     340 |   3504080 ns/op |
| day07/part2 | BenchmarkExample-4 |  184597 |      6401 ns/op |
| day07/part2 | BenchmarkInput-4   |     288 |   4131633 ns/op |
|-------------|--------------------|---------|-----------------|
| day08/part1 | BenchmarkExample-4 |  114986 |     10227 ns/op |
| day08/part1 | BenchmarkInput-4   |   11214 |    107337 ns/op |
| day08/part2 | BenchmarkExample-4 |   12390 |     96351 ns/op |
| day08/part2 | BenchmarkInput-4   |     630 |   1870067 ns/op |
|-------------|--------------------|---------|-----------------|
| day09/part1 | BenchmarkExample-4 |   78338 |     15168 ns/op |
| day09/part1 | BenchmarkInput-4   |     595 |   2000520 ns/op |
| day09/part2 | BenchmarkExample-4 |   36229 |     33133 ns/op |
| day09/part2 | BenchmarkInput-4   |     297 |   4048426 ns/op |
|-------------|--------------------|---------|-----------------|

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
