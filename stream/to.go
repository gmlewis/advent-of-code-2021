package stream

// ToChan returns a channel using the provided slice.
func ToChan[T any](items []T) <-chan T {
	size := defaultBufSize
	if len(items) < size {
		size = len(items)
	}
	ch := make(chan T, size)

	go func() {
		for _, v := range items {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

// ToSlice converts the channel values to a slice.
func ToSlice[T any](ch <-chan T) []T {
	var ret []T
	for v := range ch {
		ret = append(ret, v)
	}
	return ret
}
