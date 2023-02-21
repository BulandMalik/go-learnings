package main

import "fmt"

func main() {
	example1()
	example2()
}

func example1() {
	fmt.Println("inside example1..............")
	s := make([]string, 0)
	fmt.Println("length of s:", len(s))
	s = append(s, "hello") //put hello string on 0th index so now length is 1
	fmt.Println("length of s:", len(s))
	fmt.Println("contents of s[0]:", s[0])
	s[0] = "goodbye"
	fmt.Println("contents of s[0]:", s[0])

	s2 := make([]string, 2) //creates a slice of size 2 with empty string at 0 & 1 index
	fmt.Println("contents of s2[0]:", s2[0])
	s2 = append(s2, "hello")
	fmt.Println("contents of s2[0]:", s2[0])
	fmt.Println("contents of s2[2]:", s2[2])
	fmt.Println("length of s2:", len(s2))
}

func example2() {
	fmt.Println("\ninside example2..............")
	s3 := []string{"a", "b", "c"}
	s4 := s3[0:2]
	fmt.Println("s4:", s4)
	s3[0] = "d"
	fmt.Println("s4:", s4)

	var s5 []string
	s5 = s3 //pointer reference
	s5[1] = "camel" //will change to s3 as its a pointer reference to s3 mem address
	fmt.Println("s3:", s3)

	modSlice(s3) //passing the slice s3
	fmt.Println("s3[0]:", s3[0])
}
func modSlice(s []string) {
	s[0] = "hello"
}