package option

import (
	"strings"
	"testing"
)

func TestOption_And(t *testing.T) {
	t.Run("When the option is None", func(t *testing.T) {
		t.Run("it should return the second option", func(t *testing.T) {
			n := None(0)
			got := n.And(Some(10)).Unwrap()

			if got != 10 {
				t.Errorf("Want 10, got %d", got)
			}
		})
	})

	t.Run("When the option is Some", func(t *testing.T) {
		t.Run("It should return the first option", func(t *testing.T) {
			n := Some(10)
			got := n.And(Some(100)).Unwrap()

			if got != 10 {
				t.Errorf("Want 10, got %d", got)
			}
		})
	})
}

func TestOption_AndThen(t *testing.T) {
	t.Run("When the option is Some", func(t *testing.T) {
		t.Run("it calls the Binder fn", func(t *testing.T) {
			s := Some("a string")
			andThen := func(s string) Option[string] {
				return Some(strings.Repeat(s, 2))
			}

			val := s.AndThen(andThen).Unwrap()
			if val != "a stringa string" {
				t.Errorf(`Want "a stringa string", got %s`, val)
			}
		})
	})

	t.Run("When the option is None", func(t *testing.T) {
		t.Run("it does not call the bind fn", func(t *testing.T) {
			called := false
			n := None("")
			andThen := func(string) Option[string] {
				called = true
				return Some("")
			}

			option := n.AndThen(andThen)
			if called {
				t.Error("Not expecting binder to be called")
			}

			if option.IsSome() {
				t.Error(`Want None, got Some`)
			}
		})
	})
}

func TestOption_Expect(t *testing.T) {
	t.Run("When the option is Some", func(t *testing.T) {
		t.Run("It should return the value", func(*testing.T) {
			s := Some(true)

			s.Expect("should not panic")
		})
	})

	t.Run("When the option is None", func(t *testing.T) {
		t.Run("It should panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if r.(string) != "PANIC!" {
						t.Error("Wrong panic message")
					}
				} else {
					t.Error("Expected a panic")
				}
			}()

			n := None(false)
			n.Expect("PANIC!")
		})
	})
}

func TestMap(t *testing.T) {
	t.Run("It map a string to []byte", func(t *testing.T) {
		o1 := Some("a string")
		o := Map(o1, func(s string) Option[[]byte] {
			return Some([]byte(s))
		})

		if string(o.Unwrap()) != "a string" {
			t.Error(`Expecting "a string"`)
		}
	})

	t.Run("It should return a none of new type", func(t *testing.T) {
		n := None("")
		o := Map(n, func(s string) Option[[]byte] {
			return Some([]byte(s))
		})

		if o.IsSome() {
			t.Error(`Expecting none`)
		}
	})
}

func TestOption_Unwrap(t *testing.T) {
	t.Run("When the option is Some", func(t *testing.T) {
		t.Run("It should return the value", func(*testing.T) {
			s := Some(true)

			s.Unwrap()
		})
	})

	t.Run("When the option is None", func(t *testing.T) {
		t.Run("It should panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("Expected a panic")
				}
			}()

			n := None(false)
			n.Unwrap()
		})
	})
}
