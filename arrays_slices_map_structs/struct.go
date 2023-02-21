package main

import (
	"fmt"
	"math"
    "encoding/json"	
)

type Foo struct {
	A int
	b string
}

func main() {
	example1()
	example2()
	example3()
	example4()

	//struct tags
	example5()
	example6()
}

func example1() {
	fmt.Println("inside example1...........")

	f := Foo{}
	fmt.Println(f) //A=0, b=empty ({0 })

	g := Foo{10, "Hello"}
	fmt.Println(g) //{10 Hello}

	h := Foo{
		b: "Goodbye",
	}
	fmt.Println(h) //A=0,b=Goodbye .... {0, Goodbye}

	h.A = 1000
	fmt.Println(h.A) //1000
}

func example2() {
	fmt.Println("inside example2...........")

	f := Foo{
		A: 20,
	}
	var f2 Foo //create a pointer with o/nil
	f2 = f //structs in Go are value type, its deeply copying f values to f2
	f2.A = 100
	fmt.Println(f2.A) //100
	fmt.Println(f.A) //20 (because f2 gets copy of f.A value and f2 is a separate memory area)

	var f3 *Foo = &f //create a pointer and assign f memory location
	f3.A = 200
	fmt.Println(f3.A) 
	fmt.Println(f.A) //because f3 points to f memory area
}

type Circle struct {
	x, y, r float64
}
func circleArea(c *Circle) float64 { //can change incoming param value
	return math.Pi * c.r*c.r
}	
func example3() {
	fmt.Println("inside example3...........")
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(&c))	//pass memory location
}

//--------------------------------- Struct Embedding

type Foo1 struct {
	A int
	b string
}
type Bar struct {
	C string
	F Foo1
}
type Baz struct {
	D string
	Foo1
}
func example4() {
	fmt.Println("inside example4...........")
	f := Foo1{A: 10, b: "Hello"}
	b1 := Bar{C: "Fred", F: f}
	fmt.Println(b1.F.A)
	b2 := Baz{D: "Nancy", Foo1: f}
	fmt.Println(b2.A) //can refers to Foo fields directly as its defined within Baz

	var f2 Foo1 = b2.Foo1
	fmt.Println(f2)
}

//--------------------------------- Struct Tags
type User struct {
    Id         int    `json:"id"`
    Name       string `json:"name"`
    Occupation string `json:"occupation,omitempty"`
}
func (p User) String() string {

    return fmt.Sprintf("User id=%v, name=%v, occupation=%v",
        p.Id, p.Name, p.Occupation)
}
func example5() {
	fmt.Println("inside example5...........")
    user := User{Id: 1, Name: "John Doe", Occupation: "gardener"}
    res, _ := json.MarshalIndent(user, " ", "  ")

    fmt.Println(string(res))

    user2 := User{Id: 1, Name: "John Doe"}
    res2, _ := json.MarshalIndent(user2, " ", "  ")

    fmt.Println(string(res2))
}


type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
func example6() {
	fmt.Println("inside example6...........")
	bob := `{ "name": "Bob", "age": 30}`
	var b Person
	json.Unmarshal([]byte(bob), &b)
	fmt.Println(b)
	bob2, _ := json.Marshal(b)
	fmt.Println(string(bob2))
}