// Package result is an implementation of functional Results, modeled
// on Rust's `std::result::Result`.
package result

// Binder is a callback function to operate on the underlying
// data of a `Result`. Binder must return a new `Result`.
type Binder[T any] func(data T) Result[T]

// Tee is as callback function that operates on the underlying
// data of `Result`. The function returns nothing. Tee is optimal
// for use with referenced data, i.e., `&customStruct{}`. Alternatively,
// Tee is ideal for side-effects, i.e., logging, etc.
type Tee[T any] func(data T)

// Result represents a proy operation that can succeed or fail.
// It wraps either an `ok` operation or an `error` operation.
// Becuase `Result` methods return `Result` interfaces, you can chain
// your method calls together and "happy path" a procedural chain without
// checking for an error until the end of the procedure.
type Result[T any] interface {
	// Bind calls `b` on the underlying data of
	// `Result`. The return from `b` is returned. In
	// the event of an error, `b` is not called and the
	// error Result is returned unchanged.
	Bind(b Binder[T]) Result[T]
	// Tee calls `t` on the underlying data. If the Result
	// is an error, `t` is not called and the error Result
	// is returned unchanged.
	Tee(t Tee[T]) Result[T]
	// BindAll calls every `bs` on the underlying Result.
	BindAll(bs ...Binder[T]) Result[T]
	// Ok returns the underlying data. If the Result is an error,
	// a zero-value T is returned.
	Ok() T
	// IsOk reports whether the Result is ok.
	IsOk() bool
	// Err returns the underlying error. If the Result is ok,
	// Err returns nil.
	Err() error
	// IsErr reports whether the Result is an error.
	IsErr() bool
	// Expect is an assertion that the operation was ok that
	// returns the underlying data. If not, Expect panics
	// with `msg`. Only use this if you intend for your
	// program to crash on error or if you `recover`.
	Expect(msg string) T
	// Unwrap returns the underlying data. If the Result is an error,
	// Unwrap panics. Only use if you intend for your program to crash
	// or if you `recover`.
	Unwrap() T
}

// Ok wraps `data` in an ok `Result`.
func Ok[T any](data T) Result[T] {
	return ok[T]{data: data}
}

// Err wraps `data` in an err `Result`.
func Err[T any](e error) Result[T] {
	return err[T]{err: e}
}

// Pipe creates reusable result pipelines
func Pipe[T any](r Result[T], bs ...Binder[T]) Result[T] {
	return r.BindAll(bs...)
}
