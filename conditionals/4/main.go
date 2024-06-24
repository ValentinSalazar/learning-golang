/*
  Leer un número correspondiente a un año e imprimir un mensaje indicando
  si es bisiesto o no. Un año es bisiesto cuando es divisible por 4.
  Sin embargo, aquellos años que sean divisibles por 4 y también por 100 no son bisiestos,
  a menos que también sean divisibles por 400.
  Por ejemplo, 1900 no fue bisiesto pero sí el 2000.
*/

package main

import "fmt"

func solicitarAno() int {
	var ano int
	fmt.Print("Ingresa un año: ")

	_, error := fmt.Scan(&ano)

	if error != nil {
		fmt.Print("Hubo un problema al ingresar el año.")
	}

	return ano
}

func esBisiesto(ano int) bool {
	var anoBisiesto bool

	if ano%4 == 0 && (ano%100 != 0 || ano%400 == 0) {
		anoBisiesto = true
	} else {
		anoBisiesto = false
	}

	return anoBisiesto
}

func main() {
	anoIngresado := solicitarAno()

	esAnoBisiesto := esBisiesto(anoIngresado)

	fmt.Println(esAnoBisiesto)
}
