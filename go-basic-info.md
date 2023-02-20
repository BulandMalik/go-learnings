## Go History

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. To be honest that's a little vague so let's dive in some more. 

The Go project was created by three smart people (Ken Thompson, Rob Pike, and Robert Grishamer) at Google who weren't happy with the status quo.

Ken is probably the most famous of the three, he was originally known for co-creating a little operating system called Unix. 

Rob Pike is another old-school Unix developer from Bell Labs as well as wrote few books on Unix as well. He also partner with Ken on an operating system called plan 9 where they invented UTF-8, the standard way to represent unicode characters.

Robert greenberg doesn't have quite the same pedigree but he's a software engineer with a long history of working on various languages, including Hotspot, which makes Java's JVM so fast, and V8, Google's JavaScript engine.


In 2007 they were frustrated by how complicated it was to write applications in C++ and how long it took for them to compile. They created Go, a language with the following features. 

1. First of all Go as a simple language. Unlike many other languages the feature set is intentionally kept small. Very few new language features have been added since Go 1.0 was released in 2012.

2. Go as a procedural language with object-oriented features. Go has methods and composition and encapsulation and interfaces that you've seen in other programming languages, but it leaves out other object-oriented features that are less useful. 

3. Go is strongly typed, every variable has a tight and one some variables type is declared you can't assign a different kind of value to it.

4. Go includes a modern standard library. Because Go was created recently it's library includes features that address the needs of modern programmers, like built-in HTTP clients and servers, JSON support, and a standard SQL connection pool. 

5. Go programs compiled to a single native binary. This means you don't have to include a large runtime with your program, like you do in Java, Python, Ruby, or JavaScript. You get a single file that's easy to distribute or embed in a docker image. If you want to write your code on one platform but run it on another cross compiling is very simple just set a couple of environment variables before you build. 

6. Go programs are fast. Because the code is native you don't pay the performance penalty that you get with languages like Python, or Ruby, or JavaScript, but even though Go programs are native and fast they are still garbage collected. That means, like Java and scripting languages, you don't have to worry about managing your memory, your Go program does it for you automatically. However, unlike Java, Go is very memory efficient. You'll find that a well-written Go program will use anywhere from one half to a twentieth of the memory of a similar Java program.

7. Finally Go takes advantage of the multiple cores in virtually every computer today by having concurrency built into the language. Go style of concurrency is simpler to understand. Standard model using other languages, but it's just as powerful. Every programming language is designed to solve a certain set of problems.

8. Go is meant to be easy to learn, fast to compile, run quickly and efficiently, and make it easy to manage code on large teams.

The first few goals are pretty common for programming languages, but that last goal is interesting and kind of unique. Because Go is supposed to be easy to use with large teams, the language is intentionally simple and discourages people from writing clever code. Code compiles quickly because slow compile times cause developers to lose their flow and get distracted. There's a single standard format for Go code and forced by a tool called Go fmt, that's included as part of the standard Go development tools. Having one format eliminates all the arguments that how code to be indented in space that wastes time on so many projects. The general consensus is that the standard format isn't anyone's favorite format, but the way it eliminates arguments makes it everyone's favorite tool. 

Go has found its niche as a great language for building cloud infrastructure and command-line tools. The most popular Go projects rank among some of the most popular software projects today. The Hugo static site generator is written in Go, caddy, a web server that supports HTTP 2 and automatically downloads and maintains HTTPS support for your sites it's written and Go. Prometheus, the leading open-source metrics and monitoring solution, also written and Go. Docker, the tool to popularized use of containers to manage software development, also written and Go. And Kubernetes, the leading system for automating container management, you guessed it, it's also written and Go. 

In short, if you're working on high-performance Internet services you need to be learning Go. In our next video we're gonna look at installing your Go environment, and Visual Studio code, the IDE that we're going to use for this course.