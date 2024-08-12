package main

/* Escribir una función para contar cuántas veces aparece un valor
* dentro de la lista. La función recibe como parámetros la lista
* y el valor a buscar, y devuelve un número entero.
* */

import (
	"arrays-slices/utils"
	"fmt"
)

func imprimirLista(s []int) {
	for _, valor := range s {
		fmt.Print(valor, " ")
	}

	fmt.Println()
}

func aparicionesDeNumero(s []int, n int) int {
	cantidadApariciones := 0

	for _, valor := range s {
		if valor == n {
			cantidadApariciones++
		}
	}

	return cantidadApariciones
}

func main() {
	sliceCargado := utils.CargarLista()

	imprimirLista(sliceCargado)

	fmt.Println("Ingrese un numero a buscar en la lista: ")

	var numero int

	fmt.Scan(&numero)

	apariciones := aparicionesDeNumero(sliceCargado, numero)

	fmt.Println(numero, "aparece", apariciones, "veces en la lista.")
}
