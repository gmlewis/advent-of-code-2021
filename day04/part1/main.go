// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

func main() {
	flag.Parse()

	enum.Each(flag.Args(), process)
}

func process(filename string) {
	log.Printf("Processing %v ...", filename)
	buf := must.ReadFile(filename)
	parts := strings.Split(buf, "\n\n")
	numbers := strings.Split(parts[0], ",")
	boards := enum.Map(parts[1:], parseBoard)

	best := enum.Reduce(boards, &tup{}, func(board *BoardT, acc *tup) *tup {
		count := board.play(numbers)
		if acc.board == nil || count < acc.count {
			acc.count = count
			acc.board = board
		}
		return acc
	})

	sum := best.board.unmarkedSum()
	log.Printf("unmarkedSum=%v, lastNum=%v", sum, best.board.lastNum)

	fmt.Printf("Solution: %v\n", sum*best.board.lastNum)
}

type tup struct {
	board *BoardT
	count int
}

type BoardT struct {
	unmarked map[string]string
	marked   map[string]string
	row      map[string][]string
	col      map[string][]string
	lastNum  int
}

func (b *BoardT) unmarkedSum() int {
	var sum int
	for k := range b.unmarked {
		v := must.Atoi(k)
		sum += v
	}
	return sum
}

func (b *BoardT) play(numbers []string) int {
	b.row = map[string][]string{}
	b.col = map[string][]string{}
	for i, num := range numbers {
		if b.bingo() {
			return i
		}
		if k, ok := b.unmarked[num]; ok {
			b.lastNum = must.Atoi(num)
			b.marked[k] = num
			parts := strings.Split(k, ",")
			y := parts[0]
			x := parts[1]
			b.row[y] = append(b.row[y], num)
			b.col[x] = append(b.col[x], num)
			delete(b.unmarked, num)
		}
	}
	return len(numbers)
}

func (b *BoardT) bingo() bool {
	if len(b.marked) < 5 {
		return false
	}
	for _, v := range b.row {
		if len(v) == 5 {
			return true
		}
	}
	for _, v := range b.col {
		if len(v) == 5 {
			return true
		}
	}
	return false
}

var whitespaceRE = regexp.MustCompile(`\s+`)

func parseBoard(puz string) *BoardT {
	b := &BoardT{unmarked: map[string]string{}, marked: map[string]string{}}
	lines := strings.Split(puz, "\n")
	for y, line := range lines {
		line = whitespaceRE.ReplaceAllString(line, " ")
		cols := strings.Split(strings.TrimSpace(line), " ")
		for x, v := range cols {
			b.unmarked[v] = fmt.Sprintf("%v,%v", y, x)
		}
	}
	// log.Printf("board: %#v", b.unmarked)
	return b
}
