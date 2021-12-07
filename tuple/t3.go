package tuple

// T3 represents a 3-element tuple.
type T3[A any, B any, C any] struct {
	A A
	B B
	C C
}

// New3 creates a 3-element tuple.
func New3[A any, B any, C any](a A, b B, c C) *T3[A, B, C] {
	return &T3[A, B, C]{a, b, c}
}
