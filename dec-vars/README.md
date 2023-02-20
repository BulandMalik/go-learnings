## Declaring, Assigning And using Variables

Go allows developers to create variables to hold data, and like many programming languages variable declarations in Go require all variables to represent the type.

In Go the keyword var is used to indicate that a variable is being declared. In variables and Go can be declared within a function or within a package.

The general structure is follow.

1. Variable declarations start with the keyword var
2. The name of the variable
3. Type of the variable
4. = sign (Optional)
5. value of teh variable (optional)

> var year int = 2023
> var year int (assign `0`)
> year := 2023 (type inference)

## Type Inference
Type inference simply means that in some cases Go can guess the type of a value making unnecessary for the programmer to specify it explicitly. Taking advantage of type inference makes your code less repetitive, and it's considered a good form and Go programs.

Why we may need `var`, few reasons.

1. The first is that you cannot use the shorthand syntax outside of the function, so if you declare a package level variable you must use the VAR keyword, although you can be about the type if the type can be inferred from the right-hand side
2. The second reason is that we don't always want to use the type the compiler in first from the right-hand side of the equal sign.
3. One other reasons why we have the `var` keyword is that sometimes you don't actually have a value from which we can infer a type. Let's say that we don't know what value we want for `index` just yet but we wanted to clear it. 

Last but not the least, once you declare the variable, you have to use it in `Go` otherwise its a compile time error.

## Constants
o also has support for constants. Constants are basically variables whose values cannot be changed later. They are created in the same way you create variables but instead of using the var keyword we use the const keyword:
> const x string = "Hello World"

## Defining Multiple Variables
Go also has another shorthand when you need to define multiple variables:
```
var (
  a = 5
  b = 10
  c = 15
)
```
Use the keyword var (or const) followed by parentheses with each variable on its own line.

## An Example Program
Here's an example program which takes in a number entered by the user and doubles it:

```
package main

import "fmt"

func main() {
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}
```
We use another function from the fmt package to read the user input (Scanf). &input will be explained in a later chapter, for now all we need to know is that Scanf fills input with the number we enter.

