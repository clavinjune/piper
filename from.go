package piper

func From[T any](n ...T) (<-chan T, int) {
	out := make(chan T)
	go func() {
		defer close(out)

		for _, i := range n {
			out <- i
		}
	}()

	return out, len(n)
}

func FromValues[K comparable, T any](m map[K]T) (<-chan T, int) {
	out := make(chan T)
	go func() {
		defer close(out)

		for _, i := range m {
			out <- i
		}
	}()

	return out, len(m)
}

func FromKeys[K comparable, T any](m map[K]T) (<-chan K, int) {
	out := make(chan K)
	go func() {
		defer close(out)

		for i := range m {
			out <- i
		}
	}()

	return out, len(m)
}
