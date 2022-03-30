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
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------

func main() {
	count := len(os.Args) - 1

	fmt.Printf("There are %d names.\n", count)
}
