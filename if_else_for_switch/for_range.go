package main

import "fmt"

func main() {
	main1()
	main2()
}

func main1() {
	s := "Hello, world!"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}

func main2() {
	s := "👋 🌍"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}

	/**
	* Output 
	* 0 128075 👋 (first emoji is four characters long)
	* 4 32  (the space is one character long)
	* 5 127757 🌍 (last emoji is four characters long)
	*/
}