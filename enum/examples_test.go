package enum_test

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

func ExampleAll_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(value int) bool { return value >= 0 }
	fmt.Println(enum.All(items, f))
	// Output: false
}

func ExampleAll_string() {
	items := []string{"yo", "ho", "and", "barrel", "of", "rum"}
	f := func(value string) bool { return len(value) >= 2 }
	fmt.Println(enum.All(items, f))
	// Output: true
}

func ExampleAllWithIndex_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(index, value int) bool { return index < 6 || value >= 0 }
	fmt.Println(enum.AllWithIndex(items, f))
	// Output: true
}

func ExampleAny_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(value int) bool { return value >= 0 }
	fmt.Println(enum.Any(items, f))
	// Output: true
}

func ExampleAny_string() {
	items := []string{"yo", "ho", "and", "barrel", "of", "rum"}
	f := func(value string) bool { return len(value) >= 20 }
	fmt.Println(enum.Any(items, f))
	// Output: false
}

func ExampleAnyWithIndex_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(index, value int) bool { return index < 6 || value >= 0 }
	fmt.Println(enum.AnyWithIndex(items, f))
	// Output: true
}

func ExampleChunkEvery_int() {
	items := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(enum.ChunkEvery(items, 2, 1))
	// Output: [[1 2] [2 3] [3 4] [4 5] [5 6] [6 7]]
}

func ExampleCount_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(value int) bool { return value >= 0 }
	fmt.Println(enum.Count(items, f))
	// Output: 11
}

func ExampleCountWithIndex_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(index, value int) bool { return index < 6 || value >= 0 }
	fmt.Println(enum.CountWithIndex(items, f))
	// Output: 12
}

func ExampleDedup_int() {
	items := []int{0, 0, 1, 1, 1, 2, 2, 2, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9, 10, 10, 10}
	fmt.Println(enum.Dedup(items))
	// Output: [0 1 2 3 4 5 6 7 8 9 10]
}

func ExampleEach_string() {
	items := []string{"a", "b", "c"}
	upcased := make([]string, 0, len(items))
	f := func(value string) { upcased = append(upcased, strings.ToUpper(value)) }
	enum.Each(items, f)
	fmt.Println(upcased)
	// Output:  [A B C]
}

func ExampleEachWithIndex_string() {
	items := []string{"a", "b", "c"}
	upcased := make([]string, 0, len(items))
	f := func(index int, value string) {
		v := fmt.Sprintf("%v%v", strings.ToUpper(value), index)
		upcased = append(upcased, v)
	}
	enum.EachWithIndex(items, f)
	fmt.Println(upcased)
	// Output:  [A0 B1 C2]
}

func ExampleEquals_string() {
	f := enum.Equals("A")
	fmt.Println(f("a"), f("A"))
	// Output: false true
}

func ExampleIdentity_string() {
	fmt.Println(enum.Identity("a"), enum.Identity("A"))
	// Output: a A
}

func ExampleLength_string() {
	items := []string{"1", "2", "3"}
	fmt.Println(enum.Length(items))
	// Output: 3
}

func ExampleFirst_string() {
	var empty []string
	items := []string{"a", "b", "c"}
	fmt.Printf("%q %q", enum.First(empty), enum.First(items))
	// Output: "" "a"
}

func ExampleLonger_int() {
	a, b := []int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5}
	fmt.Println(enum.Longer(a, b))
	// Output: [1 2 3 4 5]
}

func ExampleShorter_int() {
	a, b := []int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5}
	fmt.Println(enum.Shorter(a, b))
	// Output: [1 2 3 4]
}

func ExampleAverage_int() {
	a := []int{1, 2, 3, 4}
	fmt.Println(enum.Average(a))
	// Output: 2
}

func ExampleAverage_float64() {
	a := []float64{1, 2, 3, 4}
	fmt.Println(enum.Average(a))
	// Output: 2.5
}

func ExampleSum_int() {
	a := []int{1, 2, 3, 4}
	fmt.Println(enum.Sum(a))
	// Output: 10
}

func ExampleProduct_int() {
	a := []int{1, 2, 3, 4}
	fmt.Println(enum.Product(a))
	// Output: 24
}

func ExampleFilter_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(value int) bool { return value >= 0 }
	fmt.Println(enum.Filter(items, f))
	// Output: [0 1 2 3 4 5 6 7 8 9 10]
}

func ExampleFilterMap_string2int() {
	items := []string{"0", "1", "2", "3", "4", "-1", "5", "6", "7", "8", "9", "10"}
	mapFunc := func(value string) int { return must.Atoi(value) }
	f := func(value int) bool { return value >= 0 }
	fmt.Println(enum.FilterMap(items, mapFunc, f))
	// Output: [0 1 2 3 4 5 6 7 8 9 10]
}

func ExampleFilterWithIndex_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(index, value int) bool { return index < 5 || value >= 0 }
	fmt.Println(enum.FilterWithIndex(items, f))
	// Output: [0 1 2 3 4 5 6 7 8 9 10]
}

func ExampleFindFirst_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(value int) bool { return value < 0 }
	fmt.Println(enum.FindFirst(items, f))
	// Output: 5
}

// FindLast returns the index of the last element for which f(value) returns true
// or -1 if none found.
func ExampleFindLast_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(value int) bool { return value >= 0 }
	fmt.Println(enum.FindLast(items, f))
	// Output: 11
}

func ExampleFind_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(value int) bool { return value >= 0 }
	fmt.Println(enum.Find(items, f))
	// Output: 0 true
}

func ExampleFindOr_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(value int) bool { return value >= 20 }
	fmt.Println(enum.FindOr(items, f, 42))
	// Output: 42
}

func ExampleFindWithIndex_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(index, value int) bool { return index > 4 && value >= 0 }
	fmt.Println(enum.FindWithIndex(items, f))
	// Output: 5 true
}

func ExampleFindOrWithIndex_int() {
	items := []int{0, 1, 2, 3, 4, -1, 5, 6, 7, 8, 9, 10}
	f := func(index, value int) bool { return index > 4 && value >= 0 }
	fmt.Println(enum.FindOrWithIndex(items, f, 42))
	// Output: 5
}

func ExampleFlatMap_int() {
	items := []int{1, 2, 3}
	f := func(value int) []string {
		s := strconv.Itoa(value)
		return []string{s, s}
	}
	fmt.Println(enum.FlatMap(items, f))
	// Output: [1 1 2 2 3 3]
}

func ExampleFlatMapWithIndex_int() {
	items := []int{10, 20, 30}
	f := func(index, value int) []string {
		s := strconv.Itoa(value)
		msg := fmt.Sprintf("%v=%v", index, s)
		return []string{msg, msg}
	}
	fmt.Println(enum.FlatMapWithIndex(items, f))
	// Output: [0=10 0=10 1=20 1=20 2=30 2=30]
}

func ExampleFrequencies_string() {
	items := []string{"0", "1", "2", "0", "1", "0"}
	fmt.Printf("%#v", enum.Frequencies(items))
	// Output: map[string]int{"0":3, "1":2, "2":1}
}

func ExampleFrequenciesBy_string() {
	items := []string{"0", "1", "2", "0", "1", "0"}
	fmt.Printf("%#v", enum.FrequenciesBy(items, must.Atoi))
	// Output: map[int]int{0:3, 1:2, 2:1}
}

func ExampleGroupBy_string() {
	items := []string{"ant", "buffalo", "cat", "dingo"}
	strLength := func(s string) int { return len(s) }
	fmt.Println(enum.GroupBy(items, strLength, enum.Identity))
	// Output: map[3:[ant cat] 5:[dingo] 7:[buffalo]]
}

func ExampleMap_string() {
	items := []string{"0", "1", "2", "3", "4", "5", "6"}
	fmt.Printf("%#v", enum.Map(items, must.Atoi))
	// Output: []int{0, 1, 2, 3, 4, 5, 6}
}

func ExampleMapWithIndex_string() {
	items := []string{"a", "b", "c", "d", "e", "f"}
	f := func(index int, value string) string { return fmt.Sprintf("%v:%v", index, value) }
	fmt.Printf("%#v", enum.MapWithIndex(items, f))
	// Output: []string{"0:a", "1:b", "2:c", "3:d", "4:e", "5:f"}
}

func ExampleMember_string() {
	items := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Printf("%#v", enum.Member(items, "yo"))
	// Output: false
}

func ExampleMaxFunc_int() {
	type keyT [2]int
	items := []keyT{{-1, -2}, {3, 4}, {2, 10}, {10, 1}, {1, 10}}
	lessFunc := func(a, b keyT) bool {
		if a[1] == b[1] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	}
	fmt.Println(enum.MaxFunc(items, lessFunc))
	// Output: [2 10]
}

func ExampleMinFunc_int() {
	type keyT [2]int
	items := []keyT{{-1, -2}, {3, 4}, {2, 10}, {10, 1}, {1, 10}}
	lessFunc := func(a, b keyT) bool {
		if a[1] == b[1] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	}
	fmt.Println(enum.MinFunc(items, lessFunc))
	// Output: [-1 -2]
}

func ExampleRange_up() {
	fmt.Println(enum.Range(0, 2))
	// Output: [0 1 2]
}

func ExampleRange_down() {
	fmt.Println(enum.Range(2, 0))
	// Output: [2 1 0]
}

func ExampleRanges() {
	start, end := []int{0, 3, 0}, []int{2, 1, 0}
	fmt.Println(enum.Ranges(start, end))
	// Output: [[0 3 0] [1 2 0] [2 1 0]]
}

func ExampleReduce() {
	f := func(value, acc int) int { return acc + value*value }
	items := []int{1, 2, 3}
	fmt.Println(enum.Reduce(items, 0, f))
	// Output: 14
}

func ExampleReduceWithIndex() {
	f := func(i, v, acc int) int { return acc + v*v + i*i*i }
	items := []int{1, 2, 3}
	fmt.Println(enum.ReduceWithIndex(items, 0, f))
	// Output: 23
}

func ExampleScan() {
	f := func(a, b int) int { return a + b }
	fmt.Println(enum.Scan(enum.Range(1, 5), 0, f))
	// Output: [1 3 6 10 15]
}

func ExampleUniq_int() {
	items := []int{1, 2, 3, 3, 2, 1}
	fmt.Println(enum.Uniq(items))
	// Output: [1 2 3]
}

func ExampleZip() {
	items := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(enum.Zip(items))
	// Output: [[1 3 5] [2 4 6]]
}

func ExampleZip2() {
	items1 := []int{1, 2, 3, 4, 5, 6}
	items2 := []string{"a", "b", "c"}
	type ns struct {
		N int
		S string
	}
	f := func(n int, s string) ns { return ns{n, s} }
	fmt.Println(enum.Zip2(items1, items2, f))
	// Output: [{1 a} {2 b} {3 c}]
}
