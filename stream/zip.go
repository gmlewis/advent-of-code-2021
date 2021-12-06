package stream

// Zip2 zips corresponding elements from different types into a channel.
//
// The zipping finishes as soon as either channel ends.
func Zip2[S any, T any, KV any](sCh <-chan S, tCh <-chan T, f func(S, T) KV) <-chan KV {
	outCh := make(chan KV, defaultBufSize)

	go func() {
		for s := range sCh {
			t, ok := <-tCh
			if !ok {
				break
			}
			outCh <- f(s, t)
		}

		close(outCh)
	}()

	return outCh
}
