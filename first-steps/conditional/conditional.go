package main

import "fmt"

func main() {
	var numeroA int = 50

	var numeroB int = 75

	if numeroA > numeroB {
		fmt.Println("Numero A es mayor que numero B")
	} else if numeroA < numeroB {
		fmt.Println("Numero B es mayor que numero A")
	} else {
		fmt.Println("Los numeros son iguales")
	}

	numeroC := 17
	numeroD := 19

	if numeroC > numeroD {
		fmt.Println(numeroC, "es mayor a: ", numeroD)
	} else if numeroC < numeroD {
		fmt.Println(numeroD, "es mayor a: ", numeroC)
	} else {
		fmt.Println("Los numeros son iguales.")
	}
}
