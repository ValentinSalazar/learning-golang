package main

import "fmt"

/*
* First of all, what is an array? It's a collection of elements ordered sequentially
* All these elements in the array must be of the specified type. (int, string, etc)
* GoLang has some features which are different from other lenguagues.
* - Length: Go considers the "size" of the array to be part of the type of the array.
*   This means [3]int is different type from [4]int because they've got different sizes.
*   So, we cannot use a variable to specify the size of array because types must be resolved
*   at compile time, not at runtime.
* - Own functions which receive an array as parameter: We cannot create function that works
*   with arrays of any size, because they are different types.
*
* When we declare an array in GoLang with no values (just the length), all elements in array
* are declared by "Zero-Value". In case that we declared an array of Integer type:
* - var array [5]int -> [0, 0, 0, 0, 0]
*
* Matrix in GoLang
* [ [1,2,3], [4,5,6], [7,8,9] ]
*    i = 0    i = 1    i = 2
*
*   Every index "i" is an Array
 */

func main() {
	// 1st) Style declaration
	// var nameAarray [length]typeElements
	var array [100]int

	fmt.Println("First style declaration")
	// Following foor lop will be priting the elements from array and they all are "Zero-Value".
	for _, v := range array {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("Second style declaration")

	// 2nd) Style declaration (Literal Declaration)
	// nameArray := [lenth]typeElements{All or some elements}
	secondArray := [5]int{22, 37, 94, 14, 6}

	for i := 0; i < len(secondArray); i++ {
		fmt.Print(secondArray[i], " ")
	}
	fmt.Println()

	// Bi-dimensional arrays (matrix)
	var matrix [3][3]int

	matrix[1][0] = 45
	matrix[1][1] = 25
	matrix[1][2] = 30

	matrix[0][0] = 10
	matrix[0][1] = 11
	matrix[0][2] = 12

	// the index "i" belongs to "row", so i = 0 will be an array -> [10, 11, 12]
	// then, we can iterate that index because it's an array and pring the values.
	// i = 0 -> [ 10,  11,  12]
	//           j=0, j=1, j=2
	//
	// i = 1 -> [45, 25, 30]
	// i = 2 -> [0, 0, 0]
	//
	// "row" it's the array where index "j" iterates.

	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, value)
		}
	}
}
