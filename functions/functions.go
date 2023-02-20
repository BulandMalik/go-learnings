package main

import "fmt"

func main() {
	oneAntonymousFunc();

	closure1();

	callMakeAdder();

	callMakeAdderAndDoubler()
}

func oneAntonymousFunc() {

	fmt.Println("inside oneAntonymousFunc......")
	myAddOne := func(a int) int {
		return a + 1
	}
	fmt.Println(myAddOne(1))
}

func closure1() {
	fmt.Println("inside closure1......")
	b := 2
	myAddOne := func(a int) int {
		b = b + a
		return b
	}

	fmt.Println(myAddOne(1))
	fmt.Println(b) //b value is modified
}


func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}
func callMakeAdder() {
	fmt.Println("inside callMakeAdder......")
	addOne := makeAdder(1)
	addTwo := makeAdder(2)

	fmt.Println(addOne(1))
	fmt.Println(addTwo(1))
}

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func callMakeAdderAndDoubler() {
	fmt.Println("inside callMakeAdderAndDoubler......")
	addOne := makeAdder(1)
	doubleAddOne := makeDoubler(addOne) //func chaining

	fmt.Println(addOne(1))
	fmt.Println(doubleAddOne(1))
}