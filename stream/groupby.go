package stream

// GroupBy splits the items into groups based on keyFunc and valueFunc.
//
// For example:
//    GroupBy([]string{"ant", "buffalo", "cat", "dingo"}, StrLength, Identity[string])
// returns:
//    {3: {"ant", "cat"}, 5: {"dingo"}, 7: {"buffalo"}}
// and
//    GroupBy([]string{ant buffalo cat dingo}, StrLength, StrFirst)
// returns:
//    {3: {"a", "c"}, 5: {"d"}, 7: {"b"}}
func GroupBy[K comparable, V any, T any](ch <-chan T, keyFunc func(T) K, valueFunc func(T) V) map[K][]V {
	result := map[K][]V{}
	for item := range ch {
		k := keyFunc(item)
		result[k] = append(result[k], valueFunc(item))
	}
	return result
}
