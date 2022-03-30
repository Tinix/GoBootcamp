// Copyright © 2022 Daniel Tinivella
// Learn Go Programming Course Bootcamp
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more information  : https://tinix.github.io
// In-person training  : https://www.linkedin.com/in/tinivella/
// Follow me on twitter: https://twitter.com/tinix

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #2
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------

func main() {

	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}
