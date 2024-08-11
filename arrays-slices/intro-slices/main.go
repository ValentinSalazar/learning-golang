package main

import "fmt"

/*
  The behaviour of a Slice is similar to an Array.
  The main difference is that with an Array we must to declare its length,
  whereas with a Slice, we don't.

  Slices are very useful because we can grow slices as needed.
  The length in this case, it's not party of the type.
  So, we can create a single function which receive slices of any size.

*/

func main() {
	// Literal declaration
	slice := []int{8, 7, 4, 2}

	for _, value := range slice {
		fmt.Print(value, " ")
	}
}
