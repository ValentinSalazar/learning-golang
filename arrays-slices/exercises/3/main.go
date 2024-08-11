package main

/* Determinar si la lista es capicúa (palíndromo).
* Una lista capicúa se lee de igual modo de izquierda a derecha y de derecha a izquierda.
* Por ejemplo, [2, 7, 7, 2] es capicúa, mientras que [2, 7, 5, 2] no lo es.
 */

import (
	"arrays-slices/utils"
	"fmt"
)

func esSlicePalindromo(s []int) bool {
	esPalindromo := true

	i := 0
	j := len(s) - 1

	for i <= len(s)/2 {
		if s[i] != s[j] {
			esPalindromo = false
		}
		i++
		j--
	}

	return esPalindromo
}

func main() {
	sliceCargado := utils.CargarLista()

	if esSlicePalindromo(sliceCargado) {
		fmt.Println("Es un slice palindromo")
	} else {
		fmt.Println("No es un slice palindromo.")
	}
}
