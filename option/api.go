// Package option is an implementation of functional Options,
// loosely modeled on Rust's `std::option::Option`.
package option

type Binder[T any] func(data T) Option[T]
type Mapper[T any, U any] func(data T) Option[U]

// Type Option represents an optional value, either `some` or `none`.
// This is a monadic replacement for nil reference checks.
type Option[T any] interface {
	And(Option[T]) Option[T]
	// AndThen calls `b` on the underlying data. If the Option is `none`,
	// AndThen returns the `none` unchanged.
	AndThen(b Binder[T]) Option[T]
	// IsSome reports whether the Option contains a value.
	IsSome() bool
	// IsNone reports whether the Option contains no value.
	IsNone() bool
	// Expect returns the underlying data. If the Option is an error,
	// Expect panics with `msg`. Only use if you intend your program
	// to crash or if you `recover`.
	Expect(msg string) T
	// Unwrap returns the underlying data. If the Option is an error,
	// Unwrap panics. Only use if you intend your program to crash
	// or if you `recover`.
	Unwrap() T
}

// Some creates a new optional Some of type T.
func Some[T any](data T) Option[T] {
	return some[T]{data}
}

// None create a new optional None of type T.
// None can be called with `n` number of arguments.
// This helps the compiler infer the type argument.
// Alternatively, you can specify it:
//
//    n1 := None("")
//    n2 := None(0)
//    n3 := None[string]()
//    n4 := None[int64]()
//
// The type that you specify will affect the remainder
// of the Optional chain. Any values given as arguments are
// discarded.
func None[T any](...T) Option[T] {
	return none[T]{}
}

// Map applies `fn` to Option[T], which converts to an
// Option[U].
func Map[T, U any](o Option[T], fn Mapper[T, U]) Option[U] {
	if o.IsNone() {
		return None[U]()
	}

	return fn(o.Unwrap())
}
