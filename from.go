package piper

func From[T any](n ...T) (<-chan *W[T], int) {
	out := make(chan *W[T])
	go func() {
		defer close(out)

		for _, i := range n {
			out <- &W[T]{
				Data: i,
			}
		}
	}()

	return out, len(n)
}

func FromValues[K comparable, T any](m map[K]T) (<-chan *W[T], int) {
	out := make(chan *W[T])
	go func() {
		defer close(out)

		for _, i := range m {
			out <- &W[T]{
				Data: i,
			}
		}
	}()

	return out, len(m)
}
