package main

import "fmt"

func main() {
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	fmt.Println(vals, vals[0], vals[1], vals[2])

	//** START_1: cannot use vals (variable of type [3]int) as [4]int value in variable declaration
	//var vals2 [4]int = vals
	//fmt.Println(vals, vals2) 
	//** END_1: 
}