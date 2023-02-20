package main

import "fmt"

func main() {
	testIntegers();
	testStrings();
	testBooleans();
	testFloats();
}

func testIntegers() {
	fmt.Println("\n************** test Integers ***************\n")
	a := 1;
	var b int = 2
	fmt.Println(a,"+",b," = ", a+b )
	fmt.Println("1 + 1 =", 1.0 + 1.0)	
}

func testStrings() {
	fmt.Println("\n************** test Strings ***************\n")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	
	message := `My Name is 
	Buland Malik and I am 40 years old.
	I am here to learn Go!!!!!`
	fmt.Println("message=",message);
}

func testBooleans() {
	fmt.Println("\n************** test Booleans ***************\n")
	fmt.Println("true && true:",true && true)
	fmt.Println("true && false:",true && false)
	fmt.Println("true || true:",true || true)
	fmt.Println("true || false:",true || false)
	fmt.Println("!true:",!true)
  }

  func testFloats() {
	fmt.Println("\n************** test Floats ***************\n")
	var i int8 = 20
	var f float32 = 5.6
	fmt.Println(float32(i) + f)

	fmt.Println(i + int8(f+1.9))	

	var ii int8 = 20
	var jj int32 = 40
	//fmt.Println(ii + jj)	//Error as go does not automatically converts ||| invalid operation: ii + jj (mismatched types int8 and int32)
	fmt.Println(int32(ii) + jj)
}  

/**
* Output of the Program
************** test Integers ***************

1 + 2  =  3
1 + 1 = 2

************** test Strings ***************

11
101 (you see 101 instead of e when you run this program. This is because the character is represented by a byte (remember a byte is an integer).)
Hello World
message= My Name is 
        Buland Malik and I am 40 years old.
        I am here to learn Go!!!!!

************** test Booleans ***************

true && true: true
true && false: false
true || true: true
true || false: true
!true: false

************** test Floats ***************

25.6
27
60
*/