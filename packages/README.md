# Packages In Go

Go was designed to be a language that encourages good software engineering practices. An important part of high quality software is code reuse – embodied in the principle “Don't Repeat Yourself.”

>Go also provides another mechanism for code reuse: `packages`. 

>Go packages can be hierarchical so that means even if you have a package with some same existing package name but because hierarchy is different so it should work.

> When we import our math library we use its full name (`import "golang-samples/pkgs/math`), but inside of the `math.go` file we only use the last part of the name (`package math`).

```
import m "golang-book/chapter11/math"

func main() {
  xs := []float64{1,2,3,4}
  avg := m.Average(xs) //m is the alias
  fmt.Println(avg)
```

> Package public access is driven by how we name its functions (`Start with the Capital Letter`). You may have noticed that every function in the packages we've seen start with a capital letter. In Go if something starts with a capital letter that means other packages (and programs) are able to see it. If we had named the function average instead of Average our main program would not have been able to see it. 

>>thats why we do fmt.Println(....)

>Package names match the folders they fall in. There are ways around this, but it's a lot easier if you stay within this pattern.

>You can't have cyclic dependencies between packages. That means you can't create a situation where package a imports package B and package B imports package a. Doing so is, like unused imports, a compilation error.

>The first is that is illegal to import a package but not to use anything from that package. If you import a package but do not use it you will get a syntax error when the code is compiled.

>`go get github_path` to get that package locally

Handling many different third-party libraries is not easy (they might depend on two different versions of another third-party library etc.). Go team create a solution to this problem by allowing packages to a magic directory called `vendor`. And inside of vendor directory, our third-party dependencies that are only visible to your package and any contained packages. So if you copy the entire path from GitHub calm on down into a vendor directory, and the code and there will be used instead of the version of the code in the Go path.

Its hard to do ny yourself so need a tool called `dep`. dep is a separate download that is used to manage the vendor directory.

>`dep intit`/`$GOPATH/bin/dep init` to create files that dep uses to manage your dependencies
>> This creates 2 files
>>> `Gopkg.toml` - contains information about dependencies inside your code

>>> `Gopkg.lock` - contains the exact versions of the 3rd library you are using

>You can install dep using https://github.com/golang/dep.

>The Go team has recently announced a new project called Vigo that'll make version dependency management a core part of Go. Vigo is still in very early stages though, so the recommendation is to use depth for now.

## Package Documentation

Go has the ability to automatically generate documentation for packages we write in a similar way to the standard package documentation. In a terminal run this command:

```godoc golang-book/chapter11/math Average```

This documentation is also available in web form by running this command:

```godoc -http=":6060"```

and entering this URL into your browser: http://localhost:6060/pkg/

