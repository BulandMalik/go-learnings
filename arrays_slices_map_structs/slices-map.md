# Arrays
An array is a numbered sequence of elements of a single type with a fixed length. In Go they look like this:

`var x [5]int`

x is an example of an array which is composed of 5 ints. Try running the following program:

```
package main
import "fmt"
func main() {
  var x [5]int
  x[4] = 100
  fmt.Println(x)
}
```
You should see this: `[0 0 0 0 100]`

`x[4] = 100` should be read `set the 5th element of the array x to 100`. It might seem strange that x[4] represents the 5th element instead of the 4th but like strings, arrays are indexed starting from 0. Arrays are accessed in a similar way. 

A single `_ ` (underscore) is used to tell the compiler that we don't need this. (In this case we don't need the iterator variable)
```
var total float64 = 0
for _, value := range x {
  total += value
}
fmt.Println(total / float64(len(x)))
```

Go also provides a shorter syntax for creating arrays:
`x := [5]float64{ 98, 93, 77, 82, 83 }`

```
x := [5]float64{
  98,
  93,
  77,
  82,
  83,
}
```

Notice the extra trailing , after 83. This is required by Go and it allows us to easily remove an element from the array by commenting out the line:

x := [4]float64{
  98,
  93,
  77,
  82,
  // 83,
}

# Slices

>A slice is a segment of an array. Like arrays slices are indexable and have a length.

>Unlike arrays, slice  length is allowed to change.

`var x []float64` - The only difference between this and an array is the missing length between the brackets. In this case x has been created with a length of 0.

If you want to create a slice you should use the built-in make function: 

`x := make([]float64, 5)` - This creates a slice that is associated with an underlying float64 array of length 5. Slices are always associated with some array, and although they can never be longer than the array, they can be smaller. 

The make function also allows a 3rd parameter:

`x := make([]float64, 5, 10)` - 10 represents the capacity of the underlying array which the slice points to.
![](../images/slice-example1.png)

Another way to create slices is to use the [low : high] expression:
```
arr := [5]float64{1,2,3,4,5}
x := arr[0:5]
```
low is the index of where to start the slice and high is the index where to end it (but not including the index itself). For example while arr[0:5] returns [1,2,3,4,5], arr[1:4] returns [2,3,4].

For convenience we are also allowed to omit low, high or even both low and high. arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as arr[0:5] and arr[:] is the same as arr[0:len(arr)].

1. arr[0:] <==> arr[0:len(arr)]
2. arr[:5] <==> arr[0:5]
3. arr[:] <==> arr[0:len(arr)]

## Slice Functions
Go includes two built-in functions to assist with slices: append and copy. Here is an example of append.

```
func main() {
  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5) //append creates a new slice
  fmt.Println(slice1, slice2)
}
//Output
//slice1: [1,2,3]
//slice2: [1,2,3,4,5]
```

```
func main() {
  slice1 := []int{1,2,3}
  slice2 := make([]int, 2)
  copy(slice2, slice1) //Contents of slice1 are copied into slice2
  fmt.Println(slice1, slice2)
}
//Output
//slice1: [1,2,3]
//slice2: [1,2] //contents of slice1 are copied into slice2, but since slice2 has room for only two elements only the first two elements of slice1 are copied.
```

# Map

A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table or a dictionary, maps are used to look up a value by its associated key.

>`var x map[string]int`
>>The map type is represented by the keyword map, followed by the key type in brackets and finally the value type. If you were to read this out loud you would say “x is a map of strings to ints.” 

>var x map[string]int
>>It will not initialize the map so its nil hence if trying to access it, will get panic (run time error). As a general rule, never write to a `nil` map.

>x := make(map[string]int)
>>make will create as well as initialize the map

>maps are not sequential. We have x[1], and with an array that would imply there must be an x[0], but maps don't have this requirement.

>You cannot count on the values in that map being in the same order every time it's intentionally random. It's not the order of the values wooden sword it or some other sorted value.

>You remove items from the map as well using a delete function, it's another built-in function.

>This makes a map initialize 2 with 0 value. So like slices and pointers the 0 value for a map is nil, because, just like slices, maps are reference types

>We can also delete items from a map using the built-in delete function: `delete(x, 1)`

>Accessing an element of a map can return two values instead of just one. The first value is the result of the lookup, the second tells us whether or not the lookup was successful.
>>`name, ok := elements["Un"] `

```
if name, ok := elements["Un"]; ok {
  fmt.Println(name, ok)
}
```


