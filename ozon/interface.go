package main

type Foo struct{}

func (f *Foo) A() {}
func (f *Foo) B() {}
func (f *Foo) C() {}

type AB interface {
	A()
	B()
}
type BC interface {
	B()
	C()
}

func main() {
	var f AB = &Foo{}
	y := f.(BC) // сработает type-assertion
	y.A()       // не сработает
	_ = y
}
