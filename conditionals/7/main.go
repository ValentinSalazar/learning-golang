package main

import (
	"fmt"
	"math/rand"
)

// Variable available only where they are needed.

func main() {
	if numero := rand.Intn(100); numero%2 == 0 {
		fmt.Println(numero, "Es par")
	} else {
		fmt.Println(numero, "Es impar")
	}

	// fmt.Println(numero) -> "numero" is undefined so it's not going to compile.
	// "numero" is defined only in the "if-scope"
}
