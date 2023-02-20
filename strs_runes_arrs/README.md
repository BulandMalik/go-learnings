## Strings, Runes, and Arrays

## Strings
A string is a sequence of characters with a definite length used to represent text. Go strings are made up of individual bytes, usually one for each character. (Characters from other languages like Chinese are represented by more than one byte).

String literals can be created using double quotes "Hello World" or back ticks `Hello World`. The difference between these is that double quoted strings cannot contain newlines and they allow special escape sequences. For example \n gets replaced with a newline and \t gets replaced with a tab character.

Strings are “indexed” starting at 0 not 1.

```
var name string //initialize empty string of length 0
name = "Buland"
fmt.Println(name)

//Output --> Buland
```

## Runes

In the past, we only had one character set, and that was known as ASCII (American Standard Code for Information Interchange). There, we used 7 bits to represent 128 characters, including upper and lowercase English letters, digits, and a variety of punctuations and device-control characters. Due to this character limitation, the majority of the population is not able to use their custom writing systems. To solve this problem, Unicode was invented. Unicode is a superset of ASCII that contains all the characters present in today’s world writing system. It includes accents, diacritical marks, control codes like tab and carriage return, and assigns each character a standard number called “Unicode Code Point”, or in Go language, a “Rune”. The Rune type is an alias of int32.

>`Rune` in Go is a data type that stores codes that represent Unicode characters. Unicode is actually the collection of all possible characters present in the whole world. In Unicode, each of these characters is assigned a unique number called the Unicode code point. This code point is what we store in a rune data type.

To store the Unicode characters in memory, Go uses the UTF-8 encoding method. This method results in an encoded output of variable length from 1-4 Bytes. For this reason, `rune is also known as an alias for int32`, as each rune can store an integer value of at most 32-bits.

The following model shows you how to initialize a rune variable:
```var myrune rune = 'A'```

Each rune actually refers to only one Unicode code point, meaning one rune represents a single character.

Furthermore, Go does not have a char data type, so all variables initialized with a character would automatically be typecasted into int32, which, in this case, represents a rune.

Here, it might seem like rune is the char alternative for Go, but that is not actually true. In Go, strings are actually made up of sequences of bytes, not a sequence of rune.

In this regard, we can consider rune to be different from the usual char data type we see in other programming languages. Even when we need to print multiple rune in a string, we must first typecast them into rune arrays.

```
package main

import "fmt"

func main() {

    var str string = "ABCD"
    r_array := []rune(str)
    fmt.Printf("Array of rune values for A, B, C and D: %U\n", r_array)
}

//Output -> Array of rune values for A, B, C and D: [U+0041 U+0042 U+0043 U+0044]
```

```
package main

import "fmt"

func main() {
   
  // declaring and initializing a Unicode character
  emoji := '😀' 

  // %T represents the type of the variable num
   fmt.Printf("Data type of %c is %T and the rune value is %U \n", emoji, emoji, emoji);   
}

//Output -> Data type of 😀 is int32 and the rune value is U+1F600 
```

```
	s := "👋 🌍"
	s2 := s[:1]
	s3 := s[2:]
	fmt.Println(s, len(s), s2, len(s2), s3, len(s3))

    //Output ==> 👋 🌍  ___  9  ___  �  ___  1  ___  �� 🌍  ___  7
```

Strings in Go being a sequence of bytes. And in our unicode world that's often insufficient. Those source code files are always written in UTF-8 the most compact way to encode unicode characters. When use UTF-8 a character can take anywhere from one to four bytes. Due to the way that UTF-8 and codes unicode, you can tell if a character is valid or invalid. So if you take a substring that starts or ends in the middle of a character, Go can print out something to indicate the character is illegal. In the case of our emoji hello world, each emoji takes four bytes. The length or term of the Len function is the length of all of the bytes of all the characters, so it's four for each emoji and one for the space ,giving us nine. And we take a substring we're counting often bytes not characters. So ending after a single byte and our first substring gives us an illegal character in UTF-8, and starting at byte three for a second substring it is another illegal character in our second substring, followed by the space and the whole second emoji. The important lesson is that you have to be very careful if you're using the built-in len substring features and Go if your strings might contain characters that are bigger than a single byte you'll get very strange results. So Go has a special numeric type that represents a character called a rune. Just like a byte is the same as you and eight, rune is the same as an int 32. Yes strangely enough it's assigned value, and since rune is a numeric type you can assign it a number

## Arrays
An array and Go is similar to arrays in other languages, is a sequence of values all of the same type. Arrays Go are of a fixed size. Once you declare a variable you can't make your aim bigger or smaller. You clear an array by prefixing a types name with a length of the array in square brackets.

Arrays Go do have one limitation and it's a big one. The length of the array is considered part of the type, that means you can't assign a variable of type three int to a variable of type for int

```
    var a [3] int
    var b [4] int
    a[0] = 1
    a[1] = 2
    a[2] = 3

    //b = a //error as they are of different sizes (3 and 4)
```

Its better to use Slices ad slices are much more versatile and we will talk about them in depth in section four we cover more complex types.