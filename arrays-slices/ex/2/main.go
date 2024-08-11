package main

/*
  Calcular la suma de los números de la lista.
*/

import (
	"arrays-slices/utils"
	"fmt"
)

func calcularSumaSlice(s []int) int {
	sumaTotal := 0

	for _, valor := range s {
		sumaTotal += valor
	}

	return sumaTotal
}

func main() {
	sliceCargado := utils.CargarLista()

	sumaNumerosSlice := calcularSumaSlice(sliceCargado)

	fmt.Println("La suma es:", sumaNumerosSlice)
}
