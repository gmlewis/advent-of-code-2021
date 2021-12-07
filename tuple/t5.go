package tuple

// T5 represents a 5-element tuple.
type T5[A any, B any, C any, D any, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

// New5 creates a 5-element tuple.
func New5[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) *T5[A, B, C, D, E] {
	return &T5[A, B, C, D, E]{a, b, c, d, e}
}
