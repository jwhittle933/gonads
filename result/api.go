package result

type Binder func(data interface{}) Result
type Tee func(data interface{})

// Result represents a proxy operation that can succeed or fail
type Result interface {
	Bind(b Binder) Result
	Tee(t Tee) Result
	BindAll(bs ...Binder) Result
	Ok() interface{}
	IsOk() bool
	Err() error
	IsErr() bool
	Expect(msg string) interface{}
	Unwrap() interface{}
}

// Wrap converts data to a Result to perform resilient operations on data
func Wrap(data interface{}) Result {
	return Ok{data: data}
}

func WrapErr(e error) Result {
	return Err{err: e}
}

// Pipe creates reusable result pipelines
func Pipe(r Result, bs ...Binder) Result {
	return r.BindAll(bs...)
}