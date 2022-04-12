package data

func NewFoo(foo string) Foo {
	f := Foo{}
	f.private = foo

	return f
}
