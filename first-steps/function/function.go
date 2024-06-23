package main

import "fmt"

func sumarEnteros(a int, b int) int {
	return a + b
}

func imprimirDatosPersona(nombre string, edad int) {
	fmt.Println("La edad de ", nombre, " es ", edad)
}

func main() {
	numeroA := 20

	numeroB := -17

	suma := sumarEnteros(numeroA, numeroB)

	fmt.Println("La suma es: ", suma)

	nombre := "Valentin"
	edad := 22

	imprimirDatosPersona(nombre, edad)
}
