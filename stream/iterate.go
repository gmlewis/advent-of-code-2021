package stream

// Iterate returns a channel of values using the provided `f(index, last)` function.
// f returns false if this is the last value generated, and the channel will
// be closed.
func Iterate[T any](start T, f func(index int, last T) (val T, ok bool)) <-chan T {
	ch := make(chan T, defaultBufSize)

	go func() {
		ch <- start
		i := 1
		for {
			v, ok := f(i, start)
			ch <- v
			start = v
			i++
			if !ok {
				close(ch)
				break
			}
		}
	}()

	return ch
}
