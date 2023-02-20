
package main

import "fmt"

func main() {
	main1()
	main2()
}

func main1() {
	a := 10
	if b := a / 2; b > 5 {
		fmt.Println("b is bigger than 5")
	} else {
		fmt.Println("b is less than or equal to 5")
	}
}

func main2() {
	a := 10
	if b := a / 2; b > 5 {
		fmt.Println("b is bigger than 5:", b)
	} else {
		fmt.Println("b is less than or equal to 5:", b)
	}
	//fmt.Println(b) //undefined: b
}