## Simple Web Server

We will use build in module to configure a fully running simple web server in less than 20 lines of code.

Go includes a modern standard library that solves problems that developers have today. Part of that modern standard library is a high quality HTTP server. You can use this built-in server in production. In fact the `Caddy web server` is built using it. It supports HTTP 2 and tls/ssl. 

Look at `simple-webserver.go`, its doing following.
1. It will listen on a single endpoint called `/hello`, expect a query parameter called name, and we'll return back a message hello and your name.
2. It will be up & running locally and can be access using `http://localhost:8080/hello?name=Buland`, and you'll get back the message `Hello Buland`. You can even enter in an emoji character and you'll get it back. 

## Few things about the program
1. Line style comment start with two slashes and continue until the end of the line. Next we have a package declaration.

2. The first non comment or blank line in any Go program file must be a package declaration. Every Go file belongs to a package and packages where Go uses to organize code and import code from third parties. Every file in a directory belongs to the same package, and in every Go program there is a special package called main.

3. `import packages` is the way to import from the standards library that we need in the current file. Before you use any code from another package you must import it.

4. Two important packages in the standard library: `the formatting package, which is pronounced fmt`, and `the HTTP client and server package`.

5. A multi-line comment starts with a `/*` and continues until a `*/` is found. One thing to note, you cannot nest multi-line comments. If you have a slash star inside of a multi-line comment is ignored and the first star slash that is found will be the end of a comment block.

6. Go has top-level functions and a function is declared by using the keyword func by the name of the function. The body of the function is contained between an opening brace and a closing brace. `Now all Go programs start from a function called main in the main package. This main function takes in no parameters and returns no values`. 

7. The body of the main function contains our web servers configuration. In order to register endpoints we call the handle func function. This function expects two parameters. The first is the endpoint service which is `/hello`, this is a string quotes slash a low-end quote, which is something called a `callback` that we will discuss later in detail. 

8. Functions and Go are very powerful. You can declare functions inside of functions, assign functions to variables, even pass functions to functions and return functions from functions.

