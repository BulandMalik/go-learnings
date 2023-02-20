package main

import "fmt"

func main() {
	var a int = 1;
	fmt.Println("a:",a)

	//compiler use type inference to determine b type is int, its a short version
	b := a * 9
	fmt.Println("b:",b)

	//c := b * 3 //its an error as we are not reading c anywhere `c declared and not used`

	var d int //assign 0 as its initial value
	fmt.Println("d => before assignment:",d)
	//d := b * 3 // common mistake, as := is used to declare but not assign

	d = b * 3
	fmt.Println("d => after assignment:",d)
}

/*
Output of this program will be following.
a: 1
b: 9
d => before assignment: 0
d => after assignment: 27
*/