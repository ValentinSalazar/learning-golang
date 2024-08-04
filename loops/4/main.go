package main

/*

Desarrollar un programa que imprima los números naturales comprendidos entre 1 y N.
El valor de N se ingresa desde el teclado.

*/

import "fmt"

func main() {
	var N int

	fmt.Println("Ingrese un numero natural N: ")

	_, error := fmt.Scan(&N)

	if error != nil {
		fmt.Println("Hubo un error al ingresar un dato.")
	}

	for i := 0; i <= N; i++ {
		fmt.Println(i)
	}
}
