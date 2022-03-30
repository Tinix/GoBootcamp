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
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go tinivella
//
//    tinivella
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi Tinivella
//    How are you?
// ---------------------------------------------------------

func main() {
	fmt.Println(os.Args[0])

	// Bonus solution
	fmt.Println("Hello", os.Args[1])
	fmt.Println("How are you?")
}
