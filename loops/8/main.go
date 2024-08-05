package main

/*
Realizar un programa que lea un número natural H e imprima un mensaje
indicandoindicando si H es primo o no.
Se dice que un número es primo cuando sólo es divisible por
sí mismo y por la unidad.
*/

import "fmt"

func esPrimo(h int) bool {
	contadorDivisiones := 1
	for i := h; i > 1; i-- {
		if h%i == 0 {
			contadorDivisiones += 1
		}
	}

	return contadorDivisiones == 2
}

func main() {
	var n int

	fmt.Println("Ingrese un numero natural para verificar si es primo: ")
	fmt.Scan(&n)

	esNumeroPrimo := esPrimo(n)

	if esNumeroPrimo {
		fmt.Println(n, "es numero primo.")
	} else {
		fmt.Println(n, "no es numero primo.")
	}
}
