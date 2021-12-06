package enum

// FlatMap maps the given f(value) and flattens the result.
//
// For example:
//   FlatMap([]int{1,2,3}, func (v int) []string {
//     s := fmt.Sprintf("%v", v)
//     return []string{s,s}
//   })
// returns:
//   []string{"1","1","2","2","3","3"}
func FlatMap[S any, T any](values []S, f func(S) []T) []T {
	result := []T{}
	for _, v := range values {
		result = append(result, f(v)...)
	}
	return result
}

// FlatMapWithIndex maps the given f(index, value) and flattens the result.
func FlatMapWithIndex[S any, T any](values []S, f func(int, S) []T) []T {
	result := []T{}
	for i, v := range values {
		result = append(result, f(i, v)...)
	}
	return result
}
