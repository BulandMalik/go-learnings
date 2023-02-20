//line style comments. Line style comment start with two slashes and continue 
//until the end of the line

/**
* The first non comment or blank line in any Go program file must be a package 
* declaration. Every Go file belongs to a package and packages where Go uses to 
* organize code and import code from third parties.
*/
package main

import (
	"fmt"
	"net/http"
)

/*
	All Go programs start running from a function called `main` in a package called `main`
*/
func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}