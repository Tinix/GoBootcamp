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

/*
 ---------------------------------------------------------
 EXERCISE: Greet 5 People

  Greet 5 people this time.

  Please do not copy paste from the previous exercise!

 RESTRICTION
  This time do not use variables.

 INPUT
  bilbo balbo bungo gandalf legolas

 EXPECTED OUTPUT
  There are 5 people!
  Hello great bilbo !
  Hello great balbo !
  Hello great bungo !
  Hello great gandalf !
  Hello great legolas !
  Nice to meet you all.
 ---------------------------------------------------------
*/

func main() {
	fmt.Println(os.Args[0])

	count := len(os.Args) - 1

	fmt.Printf("There are %d people\n", count)

	fmt.Println("Hello", os.Args[1], "!")
	fmt.Println("Hello", os.Args[2], "!")
	fmt.Println("Hello", os.Args[3], "!")
	fmt.Println("Hello", os.Args[4], "!")
	fmt.Println("Hello", os.Args[5], "!")
	fmt.Println("Nice to meet you all.")
}
