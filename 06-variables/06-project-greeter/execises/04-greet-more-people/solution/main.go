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
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	fmt.Println(os.Args[0])

	count := len(os.Args)

	fmt.Printf("There are %d people\n", count)

	fmt.Println("Hello ", os.Args[1], "!")
	fmt.Println("Hello ", os.Args[2], "!")
	fmt.Println("Hello ", os.Args[3], "!")
	fmt.Println("Nice to meet you all.")
	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
}
