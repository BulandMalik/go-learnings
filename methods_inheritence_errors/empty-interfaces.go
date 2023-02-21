package main

import "fmt"

func main() {
	example1()
	//example2()
	example3_type_assertions()
}

func example1() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func example2() {
	var i interface{}
	describe(i)
	
	i = 42
	j := i.(string) //panic: runtime exception: interface {} is int, not string
	describe(i)
	describe(j)

	i = "hello"
	describe(i)
}

func example3_type_assertions() {
	//var i interface{} = "hello"

	var i interface{}
	i = "hello"

	s := i.(string)
	fmt.Println(s) //hello

	s, ok := i.(string)
	fmt.Println(s, ok) //hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) //0 false

	f = i.(float64) // panic
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i) //%v for value, %T for type
}
