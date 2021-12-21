package stream

// Map maps a channel of values from one type to another
// using the provided func f.
func Map[S any, T any](ch <-chan T, f func(v T) S) []S {
	ret := []S{}
	for v := range ch {
		ret = append(ret, f(v))
	}
	return ret
}

// MapStream maps a channel of values from one type to another
// using the provided func f.
func MapStream[S any, T any](ch <-chan T, f func(v T) S) <-chan S {
	out := make(chan S, defaultBufSize)
	go func() {
		for v := range ch {
			out <- f(v)
		}
		close(out)
	}()
	return out
}
