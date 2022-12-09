package piper

type W[T any] struct {
	_    struct{}
	Data T
	Err  error
}

func (w *W[T]) Error() string {
	if w.Err != nil {
		return w.Err.Error()
	}
	return ""
}
