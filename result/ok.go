package result

type Ok struct {
	data interface{}
}

func (o Ok) Bind(b Binder) Result {
	return b(o)
}

func (o Ok) BindAll(bs ...Binder) Result {
	var out Result
	for _, b := range bs {
		out = b(o.data)
	}

	return out
}

func (o Ok) Tee(t Tee) Result {
	t(o.data)
	return o
}

func (o Ok) Ok() interface{} {
	return o.data
}

func (Ok) IsOk() bool {
	return true
}

func (Ok) IsErr() bool {
	return false
}

func (Ok) Err() error {
	return nil
}

func (o Ok) Expect(string) interface{} {
	return o.Ok()
}

func (o Ok) Unwrap() interface{} {
	return o.Ok()
}

