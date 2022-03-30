package result

type ok[T any] struct {
	data T
}

func (o ok[T]) Bind(fn Binder[T]) Result[T] {
	return fn(o.data)
}

func (o ok[T]) BindAll(fns ...Binder[T]) Result[T] {
	var out Result[T] = o
	for _, fn := range fns {
		out = out.Bind(fn)
	}

	return out
}

func (o ok[T]) Tee(fn Tee[T]) Result[T] {
	fn(o.data)
	return o
}

func (o ok[T]) Ok() T {
	return o.data
}

func (ok[T]) IsOk() bool {
	return true
}

func (ok[T]) IsErr() bool {
	return false
}

func (ok[T]) Err() error {
	return nil
}

func (o ok[T]) Expect(string) T {
	return o.Ok()
}

func (o ok[T]) Unwrap() T {
	return o.Ok()
}
