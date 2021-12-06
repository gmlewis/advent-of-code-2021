package stream

// Map maps a channel of values from one type to another
// using the provided func f.
func Map[S any, T any](ch <-chan T, f func(a T) S) []S {
	ret := []S{}
	for v := range ch {
		ret = append(ret, f(v))
	}
	return ret
}
