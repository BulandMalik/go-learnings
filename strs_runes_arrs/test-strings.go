
package main

import "fmt"

func main() {
	testStrings()
}

func testStrings() {
	fmt.Println("\n************** test Strings ***************\n")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println(string("Hello World"[1]))
	fmt.Println("Hello " + "World")
	
	s := "Hello, \n\t\"world!\" with a backslash \\"
	fmt.Println(s)

	message := `My Name is 
	Buland Malik and I am 40 years old.
	I am here to learn Go!!!!!`
	fmt.Println("backtik message=",message);

	s = "👋 🌍"
	fmt.Println(s)

	s = "Hello, world!"
	s2 := "👋 🌍"
	s3 := s + " " + s2
	fmt.Println(s3)	

	//s = "Hello, world!"
	b := s[0]
	b2 := s[4]
	fmt.Println(s, b, string(b), b2, string(b2))	

	//s = "Hello, world!"
	s2 = s[0:5]
	s3 = s[7:12]
	s4 := s[:5]
	s5 := s[7:]
	fmt.Println(s, " ___ ", s2, " ___ ",s3, " ___ ", s4, " ___ ", s5)	
	fmt.Println(s, " ___ ", len(s), " ___ ", s2, " ___ ", len(s2), " ___ ", s3, " ___ ", len(s3))

	testUTF8Chars();
}

func testUTF8Chars() {
	s := "👋 🌍"
	s2 := s[:1]
	s3 := s[2:]
	fmt.Println(s, " ___ ", len(s), " ___ ", s2, " ___ ", len(s2), " ___ ", s3, " ___ ", len(s3))

	//👋 🌍  ___  9  ___  �  ___  1  ___  �� 🌍  ___  7

	// declaring and initializing a Unicode character
	emoji := '😀' 

	// %T represents the type of the variable num
	fmt.Printf("Data type of %c is %T and the rune value is %U \n", emoji, emoji, emoji);	

	ss := "Hello, "
	var r rune = 127757 //UTF8 code for 🌍 emoji
	ss = ss + string(r)
	fmt.Println(ss)

	ss = "Hello, "
	r = '🌍'
	ss = ss + string(r)
	fmt.Println(ss)	
}

