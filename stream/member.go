package stream

// Member checks if elem exists within the channel of values.
// The channel stops being read if the elem is found.
func Member[T comparable](ch <-chan T, elem T) bool {
	for v := range ch {
		if elem == v {
			return true
		}
	}
	return false
}
