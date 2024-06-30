/*
Leer tres números correspondientes al día, mes y año de una fecha e imprimir un mensaje
indicando si la fecha es válida o no
*/

package main

import "fmt"

func solicitarDatos() (int, int, int) {
	var (
		dia int
		mes int
		ano int
	)

	fmt.Println("Ingrese un dia: ")

	_, error := fmt.Scan(&dia)

	if error != nil {
		fmt.Println("Hubo un error en el ingreso del dia.")
	}

	fmt.Println("Ingrese un mes: ")

	_, error = fmt.Scan(&mes)
	if error != nil {
		fmt.Println("Hubo un error en el ingreso del mes.")
	}

	fmt.Println("Ingrese un año: ")

	_, error = fmt.Scan(&ano)

	if error != nil {
		fmt.Println("Hubo un error en el ingreso del año.")
	}

	return dia, mes, ano
}

func esBisiesto(ano int) bool {
	// Validar si el año es bisiesto, ya que esto ayuda a perfeccionar la certeza de validar la fecha ingresada.

	var esAnoBisiesto bool

	if ano%4 == 0 && (ano%100 != 0 || ano%400 == 0) {
		esAnoBisiesto = true
	} else {
		esAnoBisiesto = false
	}

	return esAnoBisiesto
}

func esFechaValida(dia int, mes int, ano int) bool {
	var esValida bool

	if (mes == 1 || mes == 3 || mes == 5 || mes == 7 || mes == 8 || mes == 10 || mes == 12) && (dia >= 1 && dia <= 31) && (ano >= 0 && ano <= 2024) {
		esValida = true
	} else if (mes == 4 || mes == 6 || mes == 9 || mes == 11) && (dia >= 1 && dia <= 30) && (ano >= 0 && ano <= 2024) {
		esValida = true
	} else if (mes == 2) && esBisiesto(ano) && (dia >= 1 && dia <= 29) {
		esValida = true
	} else {
		esValida = false
	}

	return esValida
}

func main() {
	diaIngresado, mesIngresado, anoIngresado := solicitarDatos()

	esValida := esFechaValida(diaIngresado, mesIngresado, anoIngresado)

	if esValida {
		fmt.Println("La fecha ingresada es valida.")
	} else {
		fmt.Print("La fecha ingresada es invalida.")
	}
}
