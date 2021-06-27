package result

type Err struct {
	err error
}

func (e Err) Bind(Binder) Result {
	return e
}

func (e Err) BindAll(...Binder) Result {
	return e
}

func (e Err) Tee(Tee) Result {
	return e
}

func (Err) Ok() interface{} {
	return nil
}

func (e Err) Err() error {
	return e.err
}

func (Err) IsOk() bool {
	return false
}

func (Err) IsErr() bool {
	return true
}
func (e Err) Expect(msg string) interface{} {
	panic(msg)
}

func (e Err) Unwrap() interface{} {
	panic("attempted to unwrap an error")
}

