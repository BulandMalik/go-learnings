package main

import "fmt"

type Foo struct {
	A int
	B string
}

func main() {
	example1()
}

func (f Foo) String() string {
	return fmt.Sprintf("A: %d and B: %s", f.A, f.B)
}

func (f *Foo) Double() {
	f.A = f.A * 2
}

func (f Foo) DoubleWithoutPointer() {
	f.A = f.A * 2
}

func example1() {
	fmt.Println("inside example1...........")
	f := Foo{
		A: 10,
		B: "Hello",
	}
	fmt.Println(f.String())
	f.Double() //Go knows that it needs to pass the pointer reference so pass the f mem location
	fmt.Println(f.String())

	f.DoubleWithoutPointer() //it will not modify as it will pass by value
	fmt.Println(f.String())
}