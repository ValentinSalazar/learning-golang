/*
  Ingresar las notas de los dos parciales de un alumno e indicar
  si promocionó, aprobó o si debe recuperar.
  Informar un error si el valor de alguna nota no está entre 0 y 10.
· Se promociona cuando las notas de ambos parciales son mayores o iguales a 7.
· Se aprueba cuando las notas de ambos parciales son mayores o iguales a 4.
· Se debe recuperar cuando al menos una de las dos notas es menor a 4.
*/

package main

import (
	"fmt"
)

func main() {
	var (
		primerNota  int
		segundaNota int
	)

	fmt.Println("Ingresa las notas: ")

	_, error := fmt.Scanln(&primerNota)

	if error != nil {
		fmt.Print("Hubo un error en el ingreso del primer dato.")
	}
	_, error = fmt.Scanln(&segundaNota)

	if error != nil {
		fmt.Print("Hubo un error en el ingreso del segundo dato.")
	}

	notaFinal := (primerNota + segundaNota) / 2

	if notaFinal >= 7 {
		fmt.Println("Promociona")
	} else if (notaFinal >= 4) && (notaFinal < 7) {
		fmt.Println("Aprueba")
	} else {
		fmt.Println("Desaprueba")
	}
}
