package main

/*
El factorial de un número entero N mayor que cero se define como el producto
de todos los enteros X tales que 0 < X <= N.
Desarrollar un programa para cal- cular el factorial de un número dado.
Deberán rechazarse las entradas inválidas (menores que 0).
*/

import "fmt"

func main() {
	fmt.Println("Ingrese un numero entero N para calcular su factorial: ")
	var numeroIngresado int
	fmt.Scan(&numeroIngresado)
	contador := 0
	if numeroIngresado < 0 {
		fmt.Println("Solo se puede calcular el factorial de un numero mayor a 0.")
	} else if numeroIngresado == 0 {
		contador = 1
	} else {
		contador = 1
		for i := numeroIngresado; i > 1; i-- {
			contador *= i
		}
	}

	if contador > 0 {
		fmt.Println("El factorial de", numeroIngresado, "es:", contador)
	}
}
