package tuple

// T2 represents a 2-element tuple.
type T2[A any, B any] struct {
	A A
	B B
}

// New2 creates a 2-element tuple.
func New2[A any, B any](a A, b B) *T2[A, B] {
	return &T2[A, B]{a, b}
}
