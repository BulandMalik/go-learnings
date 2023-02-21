package main

import "fmt"

type myInt int //newInt is a new type of int

func (mi myInt) isEven() bool { //define method of newInt
	return mi%2 == 0
}

func (mi *myInt) Double() { //define method of newInt
	*mi = *mi * 2
}

func main() {
	m := myInt(10) //type conversion, converting an int 10 to myInt type
	fmt.Println(m)
	fmt.Println(m.isEven())
	m.Double()
	fmt.Println(m)
}