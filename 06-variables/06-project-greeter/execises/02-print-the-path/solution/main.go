// Copyright © 2022 Daniel Tinivella
// Learn Go Programming Course Bootcamp
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more information  : https://tinix.github.io
// In-person training  : https://www.linkedin.com/in/tinivella/
// Follow me on twitter: https://twitter.com/tinix

package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print the Path
//
//  Print the path of the running program by getting it
//  from `os.Args` variable.
//
// HINT
//  Use `go build` to build your program.
//  Then run it using the compiled executable program file.
//
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//  myprogram
// ---------------------------------------------------------

func main() {
	fmt.Println(os.Args[0])
}
