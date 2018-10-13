// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// Make them typeless
	const (
		width  = 25
		height = width * 2
	)

	fmt.Printf("area = %d\n", width*height)
}
