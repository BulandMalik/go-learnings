# Structs

>A struct is a type which contains named fields.

>A struct is not an object and there is not inheritance you can do at all in Go.

>Structs in Go are value type, its deeply copying one striuct values to another struct
>>`AStruct = BStruct`

```
type Circle struct {
  x float64
  y float64
  r float64
}
```

The `type` keyword introduces a new type. It's followed by the name of the type `(Circle)`, the keyword `struct` to indicate that we are defining a struct type and a list of fields inside of curly braces. Each field has a name and a type.

## Initialization

We can create an instance of our new Circle type in a variety of ways:

`var c Circle`

Like with other data types, this will create a local Circle variable that is by default set to zero. For a struct zero means each of the fields is set to their corresponding zero value (0 for ints, 0.0 for floats, "" for strings, nil for pointers, …) We can also use the new function:

`c := new(Circle)`

This allocates memory for all the fields, sets each of them to their zero value and returns a pointer. `(*Circle)`.

More often we want to give each of the fields a value. We can do this in two ways. Like this:

`c := Circle{x: 0, y: 0, r: 5}`

Or we can leave off the field names if we know the order they were defined:

`c := Circle{0, 0, 5}`

## Accessing Fields
We can access fields using the . operator:

```
fmt.Println(c.x, c.y, c.r)
c.x = 10
c.y = 5
```

## Methods

```
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}
```

In between the keyword `func` and the `name of the function` we've added a `receiver`. The receiver is like a parameter – it has a name and a type – but by creating the function in this way it allows us to call the function using the . operator:

`fmt.Println(c.area())`

we no longer need the & operator (Go automatically knows to pass a pointer to the circle for this method) 

## Embedded Types

>Go doesn't support inheritance in the classical sense; instead, in encourages composition as a way to extend the functionality of types.

>Embedding is an important Go feature making composition more convenient and useful.

>Embedded fields are also called as anonymous fields. 

>Embedding seems like inheritance but its not. Embedding provides delegation

While Go strives to be simple, embedding is one place where the essential complexity of the problem leaks somewhat.

A struct's fields usually represent the has-a relationship. For example a Circle has a radius. 

```
type Person struct {
  Name string
}
func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}
```

```
type Base struct {
  b int
}

type Container struct {     // Container is the embedding struct
  Base                      // Base is the embedded struct
  c string
}
```

Instances of Container will now have the field b as well. In the spec it's called a promoted field. We can access it just as we'd do for c:
```
co := Container{}
co.b = 1 //Or co.Base.b
co.c = "string"
fmt.Printf("co -> {b: %v, c: %v}\n", co.b, co.c)
```

## Struct Tags

>A struct is a user-defined type that contains a collection of fields. It is used to group related data to form a single unit. A Go struct can be compared to a lightweight class without the inheritance feature.

>A struct tag is additional meta data information inserted into struct fields. The meta data can be acquired through reflection. Struct tags usually provide instructions on how a struct field is encoded to or decoded from a format.

Struct tags are used in popular packages including:
1. encoding/json
2. encoding/xml
3. gopkg.in/mgo.v2/bson
4. gorm.io/gorm
5. github.com/gocarina/gocsv
6. gopkg.in/yaml.v2

```
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Id         int    `json:"id"`
    Name       string `json:"name"`
    Occupation string `json:"occupation,omitempty"`
}

func (p User) String() string {

    return fmt.Sprintf("User id=%v, name=%v, occupation=%v",
        p.Id, p.Name, p.Occupation)
}

func main() {

    user := User{Id: 1, Name: "John Doe", Occupation: "gardener"}
    res, _ := json.MarshalIndent(user, " ", "  ")

    fmt.Println(string(res))

    user2 := User{Id: 1, Name: "John Doe"}
    res2, _ := json.MarshalIndent(user2, " ", "  ")

    fmt.Println(string(res2))
}
```
>With `json:"id"` struct tag, we encode the Id field in lowercase. In addition, the `omitempty` omits the Occupation field if it is empty.

```
$ go run main.go
{
    "id": 1,
    "name": "John Doe",
    "occupation": "gardener"
    }
{
    "id": 1,
    "name": "John Doe"
}
```

>Go allows you to put some metadata on the field in a struct, and this is primarily done to support marshaling and unmarshal.



For more info, please look [here](https://zetcode.com/golang/struct-tag/).