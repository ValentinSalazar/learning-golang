package main

/*
Se desea analizar cuántos autos circulan con patente con numeración par y cuántos
con numeración impar en un día. Escribir un programa que permita in- gresar la terminación
de la patente (entre 0 y 9) hasta ingresar -1 e informe cuántos vehículos pasaron
con numeración par y cuántos con numeración impar.

*/

import "fmt"

func main() {
	var (
		cantidadPares   int
		cantidadImpares int
		flag            = true
	)

	for flag {
		var patenteIngresado int
		fmt.Println("Ingrese el numero con el que termina la patente (-1 para finalizar): ")
		fmt.Scan(&patenteIngresado)
		if patenteIngresado == -1 {
			flag = false
		} else {
			if patenteIngresado%2 == 0 {
				cantidadPares += 1
			} else {
				cantidadImpares += 1
			}
		}

	}

	fmt.Println("Cantidad de pantetes pares: ", cantidadPares)
	fmt.Println("Cantidad de pantentes impares: ", cantidadImpares)
}
