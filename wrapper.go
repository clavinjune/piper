package piper

type W[T any] struct {
	_    struct{}
	Data T
	Err  error
}

func (r *W[T]) Error() string {
	return r.Err.Error()
}
