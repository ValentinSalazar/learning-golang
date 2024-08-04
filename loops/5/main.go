package main

/* Leer 10 números enteros e imprimir su promedio, el mayor valor leído y en qué posición se encontraba.
* Si se ingresó más de una vez sólo debe informar la pri- mera. */

import "fmt"

func main() {
	var (
		numeroMayor      int
		posicionDelMayor int
		contador         int
	)

	fmt.Println("Ingrese 10 numeros enteros")

	for i := 0; i < 10; i++ {
		var numeroIngresado int
		fmt.Scan(&numeroIngresado)
		contador += numeroIngresado
		if i == 0 {
			numeroMayor = numeroIngresado
			posicionDelMayor = i
		} else {
			if numeroIngresado > numeroMayor {
				numeroMayor = numeroIngresado
				posicionDelMayor = i
			}
		}
	}

	fmt.Println("Promedio: ", contador/10)
	fmt.Println("Numero mayor ingresaod: ", numeroMayor)
	fmt.Println("Posicion del Mayor: ", posicionDelMayor)
}
