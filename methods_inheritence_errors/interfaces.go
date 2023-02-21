package main

import "fmt"

type tester interface {
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

type rangeTest struct { //struct definition
	min int
	max int
}

func (rt rangeTest) test(i int) bool { //method on struct, same as interface method
	return rt.min <= i && i <= rt.max
}

type divTest int //custom type definition

func (dt divTest) test(i int) bool { //method of custom type, same as interface method
	return i%int(dt) == 0
}

func main() {
	result := runTests(10, []tester{
		rangeTest{min: 5, max: 20},
		divTest(5),
	})
	fmt.Println(result)
}