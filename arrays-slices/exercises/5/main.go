package main

/* Desarrollar una función que reciba la lista como parámetro
* y devuelva una nueva lista con los mismos elementos de la primera,
* pero en orden inverso.
* Por ejemplo, si la función recibe [5, 7, 1] debe devolver [1, 7, 5].
 */

import (
	"arrays-slices/utils"
	"fmt"
)

func invertirOrdenLista(s []int) []int {
	var sliceAux []int

	for i := len(s) - 1; i >= 0; i-- {
		sliceAux = append(sliceAux, s[i])
	}

	return sliceAux
}

func main() {
	sliceCargado := utils.CargarLista()

	utils.ImprimirSlice(sliceCargado)

	nuevoSliceInvertido := invertirOrdenLista(sliceCargado)

	fmt.Println()

	utils.ImprimirSlice(nuevoSliceInvertido)
}
