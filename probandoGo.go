package main

import "fmt"

func main() {
	num := 2

	switch num {
	case 1:
		fmt.Println("Uno")
		fallthrough
	case 2:
		fmt.Println("Dos")
		fallthrough
	case 3:
		fmt.Println("Tres")
	default:
		fmt.Println("Número desconocido")
	}
}
