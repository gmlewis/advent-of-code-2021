// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/maps"
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
		turns := board.play(numbers)
		if acc.board == nil || turns > acc.turns {
			acc.turns = turns
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
	turns int
}

type BoardT struct {
	unmarked map[string]string // number => "y,x"
	marked   map[string]string // "y,x" => number
	row      map[string][]string
	col      map[string][]string
	lastNum  int
}

func (b *BoardT) unmarkedSum() int {
	return maps.Reduce(b.unmarked, 0, func(k, v string, acc int) int {
		return acc + must.Atoi(k)
	})
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
	f := func(m map[string][]string) bool {
		return enum.Any(maps.Map(m, maps.ValueLen[string, string]), enum.Equals(5))
	}
	return f(b.row) || f(b.col)
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
