// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE
//  Fix the code.
//
// HINTS
//   maximum of int8  can be 127
//   maximum of int16 can be
//
// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

func main() {
	// DO NOT TOUCH THIS VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(int8(max) + min)
}
