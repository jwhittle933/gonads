package result

type err[T any] struct {
	err error
}

func (e err[T]) Bind(Binder[T]) Result[T] {
	return e
}

func (e err[T]) BindAll(...Binder[T]) Result[T] {
	return e
}

func (e err[T]) Tee(Tee[T]) Result[T] {
	return e
}

func (err[T]) Ok() (t T) {
	return
}

func (e err[T]) Err() error {
	return e.err
}

func (err[T]) IsOk() bool {
	return false
}

func (err[T]) IsErr() bool {
	return true
}

// Expect on err panics with message.
// Only call this if you intend to crash the program
// or handle the panic with `recover`.
func (e err[T]) Expect(msg string) T {
	panic(msg)
}

// Unwrap on err panics.
// Only call this if you intend to crash the program
// or handle the panic with `recover`.
func (e err[T]) Unwrap() T {
	panic("attempted to unwrap an error")
}
