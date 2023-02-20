package main

import "fmt"

func main() {
	main1()
	main2()
	main3() 
}

func main1() {
	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}
}

/**
* using as a while loop
*/
func main2() {

	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i + i
	}
	fmt.Println(i)
}

/**
* using as a while loop
*/
func main3() {
	i := 1
	for {
		fmt.Println(i)
		i = i + i
		if i > 10 {
			break
		}
	}
	fmt.Println(i)
}