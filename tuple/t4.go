package tuple

// T4 represents a 4-element tuple.
type T4[A any, B any, C any, D any] struct {
	A A
	B B
	C C
	D D
}

// New4 creates a 4-element tuple.
func New4[A any, B any, C any, D any](a A, b B, c C, d D) *T4[A, B, C, D] {
	return &T4[A, B, C, D]{a, b, c, d}
}
