/*
  Realizar un programa para ingresar desde el teclado un conjunto de números
  y mostrar por pantalla el menor y el mayor de ellos.
  Finalizar la lectura de datos con un valor -1.
*/

package main

import "fmt"

func solicitarNumero() int {
	var elemento int
	fmt.Print("Ingrese un numero entero: ")
	_, error := fmt.Scan(&elemento)

	if error != nil {
		fmt.Print("Hubo un error al ingresar el elemento.")
	}
	return elemento
}

func main() {
	var (
		numeroMaximo int
		numeroMinimo int
		flag         = true
		i            = 0
	)

	for flag {
		elemento := solicitarNumero()

		if elemento == -1 {
			flag = false
		}
		if i == 0 {
			numeroMinimo, numeroMinimo = elemento, elemento
		} else {
			if elemento > numeroMaximo {
				numeroMaximo = elemento
			} else if elemento < numeroMinimo {
				numeroMinimo = elemento
			}
		}

		i++
	}

	fmt.Println("Numero maximo: ", numeroMaximo)
	fmt.Println("Numero minimo: ", numeroMinimo)
}
