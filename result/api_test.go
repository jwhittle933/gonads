package result

import (
	"errors"
	"testing"
)

var (
	genericError = errors.New("error!")
)

func TestOk(t *testing.T) {
	t.Run("it should contain the data", func(t *testing.T) {
		r := Ok("a string")

		if r.Ok() != "a string" {
			t.Errorf(`Expecting %s, got %s`, "a string", r.Ok())
		}
	})
}

func TestOk_Bind(t *testing.T) {
	t.Run("It should execute the fn on the data", func(t *testing.T) {
		r := Ok("a string").
			Bind(func(s string) Result[string] {
				return Ok(s[:len(s)-1])
			})

		if r.Ok() != "a strin" {
			t.Errorf(`Want "a strin", Got %s`, r.Ok())
		}
	})
}

func TestOk_BindAll(t *testing.T) {
	t.Run("It should execute the fn on the data", func(t *testing.T) {
		fn := func(i int) Result[int] {
			return Ok(i * 10)
		}

		r := Ok(10).
			BindAll(fn, fn, fn, fn)

		if r.Ok() != 100_000 {
			t.Errorf(`Want 100_000, Got %d`, r.Ok())
		}
	})
}

func TestOk_Tee(t *testing.T) {
	t.Run("Should run fn on the value", func(t *testing.T) {
		type str struct {
			val string
		}

		r := Ok(&str{"a string"}).
			Tee(func(s *str) { s.val = "a new string" })

		if r.Ok().val != "a new string" {
			t.Errorf(`Want 100_000, Got %s`, r.Ok().val)
		}
	})
}

func TestOk_IsOk(t *testing.T) {
	r := Ok("a string")
	if !r.IsOk() {
		t.Error(`Expecting Ok`)
	}
}

func TestOk_IsErr(t *testing.T) {
	r := Ok("a string")
	if r.IsErr() {
		t.Error(`Expecting Ok, not Err`)
	}
}

func TestOk_Expect(t *testing.T) {
	r := Ok("a string")
	val := r.Expect("should not panic")
	if val != "a string" {
		t.Error("Should have received val")
	}
}

func TestOk_Unwrap(t *testing.T) {
	r := Ok("a string")
	r.Unwrap()
}

func TestErr(t *testing.T) {
	t.Run("it should contain the error", func(t *testing.T) {
		r := Err[any](genericError)

		if !r.IsErr() {
			t.Errorf(`Expecting err`)
		}
		if r.Err().Error() != "error!" {
			t.Errorf(`Expecting %s, got %s`, "a string", r.Err().Error())
		}
	})
}

func TestErr_Bind(t *testing.T) {
	t.Run("It should not execute the fn on the data", func(t *testing.T) {
		called := false
		fn := func(s string) Result[string] {
			called = true
			return Ok(s[:len(s)-1])
		}

		Err[string](genericError).
			Bind(fn)

		if called {
			t.Error(`Not expecting fn to be called`)
		}
	})
}

func TestErr_BindAll(t *testing.T) {
	t.Run("It should not execute the fn on the data", func(t *testing.T) {
		called := false
		fn := func(s string) Result[string] {
			called = true
			return Ok(s[:len(s)-1])
		}

		Err[string](genericError).
			BindAll(fn, fn, fn, fn)

		if called {
			t.Errorf(`Not expecting any Bind calls`)
		}
	})
}

func TestErr_Tee(t *testing.T) {
	t.Run("Should run fn on the value", func(t *testing.T) {
		type str struct {
			val string
		}

		called := false
		Err[*str](genericError).
			Tee(func(*str) { called = true })

		if called {
			t.Errorf(`Expecting no calls`)
		}
	})
}

func TestErr_IsOk(t *testing.T) {
	r := Err[string](genericError)
	if r.IsOk() {
		t.Error(`Expecting Err, not Ok`)
	}
}

func TestErr_IsErr(t *testing.T) {
	r := Err[string](genericError)
	if !r.IsErr() {
		t.Error(`Expecting Err, not Ok`)
	}
}

func TestErr_Expect(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r.(string) != "should panic" {
				t.Error("Expecting panic message")
			}
		} else {
			t.Error("Expecting a panic but received none")
		}
	}()

	r := Err[string](genericError)
	r.Expect("should panic")
}

func TestErr_Unwrap(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r.(string) != "attempted to unwrap an error" {
				t.Error("Expecting panic message")
			}
		} else {
			t.Error("Expecting a panic but received none")
		}
	}()
	r := Err[string](genericError)
	r.Unwrap()
}
