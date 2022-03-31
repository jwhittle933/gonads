package option

type some[T any] struct {
	data T
}

func (s some[T]) And(o Option[T]) Option[T] {
	return s
}

func (s some[T]) AndThen(fn Binder[T]) Option[T] {
	return fn(s.data)
}

func (some[T]) IsSome() bool {
	return true
}

func (some[T]) IsNone() bool {
	return false
}

func (s some[T]) Expect(msg string) T {
	return s.data
}

func (s some[T]) Unwrap() T {
	return s.data
}
