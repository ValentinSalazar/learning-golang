package main

/*Leer un numero entero y determinar si es un numero par o impar.*/

import "fmt"

func main() {
	fmt.Print("Ingrese un numero entero: ")

	var numeroEntero int

	_, error := fmt.Scanln(&numeroEntero)

	if error != nil {
		fmt.Println("Error al leer el numero entero", error)
	}

	if numeroEntero%2 == 0 {
		fmt.Println("El numero ingresado es Par")
	} else {
		fmt.Println("El numero ingresado es Impar")
	}
}
