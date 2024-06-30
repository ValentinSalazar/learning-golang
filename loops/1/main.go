/*

  Realizar un programa para ingresar desde el teclado un conjunto de números.
  Al finalizar mostrar por pantalla el primer y último valor ingresado.
  Finalizar la lec- tura con -1.

*/

package main

import "fmt"

func solicitarDatos() int {
	var elementoIngresado int

	fmt.Print("Ingrese un numero entero (-1 para finalizar): ")

	_, error := fmt.Scan(&elementoIngresado)

	if error != nil {
		fmt.Print("Hubo un error al ingresar el numero entero.")
	}

	return elementoIngresado
}

func main() {
	var (
		primerElemento int
		ultimoElemento int
	)
	i := 1

	flag := true

	for flag {
		elemento := solicitarDatos()
		if elemento == -1 {
			flag = false
		} else if i == 1 {
			primerElemento = elemento
		} else {
			ultimoElemento = elemento
		}

		i++
	}

	fmt.Println("Primer elemento: ", primerElemento)
	fmt.Println("Ultimo elemento: ", ultimoElemento)
}
