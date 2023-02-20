package main

import "fmt"

func main() {

	example1()

	example2()

	example3()

	example4()

	//generatePanic();
}

func example1() {
	a := 10
	b := &a
	c := a
	fmt.Println(a, b, *b, c) //10 0xc000018090 10 10

	a = 20
	fmt.Println(a, b, *b, c) //20 0xc000018090 20 10

	*b = 30
	fmt.Println(a, b, *b, c) //30 0xc000018090 30 10

	c = 40
	fmt.Println(a, b, *b, c) //30 0xc000018090 30 40
}

func generatePanic() {
   /*When you declare a variable this way it's assigned the 0 value for a pointer 
   *and all pointers, matter what type they point to, have the same zero value, 
   it's called nil*/
   var b *int //0 value of pointer is Nil

   fmt.Println(b) //accessing Nil will generate a panic that is runtime error
   fmt.Println(*b) //If you try to read or modify the value pointed to by a nil pointer you will get what's called a panic.

   //Output -> invalid memory address or nil pointer dereference
}

func example2() {

	/**
	* First we use builtin new function to create a pointer for the type passed in.
	* Second, new doesn't just make be a pointer it also allocates memory as well.
	*/
	b := new(int)

	fmt.Println(b) //memory location
	fmt.Println(*b) //0 as the default int type value
}

func setTo10(a *int) {
	*a = 10 //will actually change as change is happening at the memory location
}
func example3() {
	a := 20
	fmt.Println(a) //20
	setTo10(&a) //still pass by value but value is the memory address of a
	fmt.Println(a) //10 as
}


func setTo10Fail(a *int) {
	a = new(int) //because its pass by value so it will override old memory location to new one
	*a = 10
}
func example4() {
	a := 20
	fmt.Println(a)
	setTo10Fail(&a)
	fmt.Println(a)
}