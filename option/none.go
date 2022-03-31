package option

type none[T any] struct{}

func (n none[T]) And(o Option[T]) Option[T] {
	return o
}

func (n none[T]) AndThen(Binder[T]) Option[T] {
	return n
}

func (none[T]) IsSome() bool {
	return false
}

func (none[T]) IsNone() bool {
	return true
}

func (none[T]) Expect(msg string) T {
	panic(msg)
}

func (none[T]) Unwrap() T {
	panic("attempted to unwrap a none value")
}
