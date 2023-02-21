# Methods
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

>Can think a function that is part of the definition of the data type

>Go does not have inheritance but it does have methods.

>`Value Receiver` - When we refer to the struct directly

>`Reference Receiver` - when we refer to the pointer of struct

The difference is that, for a method declaration, you have a method receiver declared between the keyword `func` and the `name` of the method. 

The method receiver looks like a parameter declaration. It consists of a left parenthesis, and identifier, the type in which we are defining the method, and a right parenthesis. 

Whenever you want to access a field and a struct inside of this method you use this identifier. It is like the self the first parameter and a Python method declaration or the implicit in this reference that you get any Java method.

```
type Foo struct {
	A int
	B string
}

func (f Foo) String() string {
	return fmt.Sprintf("A: %d and B: %s", f.A, f.B)
}
```

In this example, the Abs method has a receiver of type Vertex named v. 
```
package main
import (
	"fmt"
	"math"
)
type Vertex struct {
	X, Y float64
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

>Remember: a method is just a function with a receiver argument. 
Here's Abs written as a regular function with no change in functionality. 
```
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

## Interfaces

>Its a type that list set of methods without their implementation

>Like a struct an interface is created using the type keyword, followed by a name and the keyword interface. But instead of defining fields, we define a “method set”. A method set is a list of methods that a type must have in order to “implement” the interface.

```
type Shape interface {
  area() float64
}
```

we can use interface types as arguments to functions:

```
type Circle struct {
  x, y, r float64
}
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

c := Circle{0, 0, 5}
r := Rectangle{0, 0, 10, 10}
```

```
func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}
```

We would call this function like this: `fmt.Println(totalArea(&c, &r))`

Interfaces can also be used as fields:

```
type MultiShape struct {
  shapes []Shape
}
```
>There's no explicit declaration of interface implementation. As long as your type implements the right methods it automatically meets the interface. 
>>And this is actually a bit like the duck typing that you see in dynamic languages like Python, or Ruby, or JavaScript. In those languages as long as the type has a method with the right name it's assumed that it's safe to call. Now Go it's not quite that loose. In addition to the name, the types of the input and output parameters must also match or your code won't compile. So you get a blend of the type safety and abstraction from other strongly typed languages combined with the flexibility of the dynamic languages.

### Empty Interface

The interface type that specifies zero methods is known as the empty interface.

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}. 

We write the empty interface as the keyword interface followed by left brace and a right brace
```
var i interface{}
i = "Hello"
j := i.("string")
```

>Type Assertion
>>j := i.("string") -- enforcing that i type is string and if it does not have string will get panic

A type assertion provides access to an interface value's underlying concrete value.
`t := i.(T)`

This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
`t, ok := i.(T)`

1. If i holds a T, then t will be the underlying value and ok will be true.
2. If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map. 

## Errors