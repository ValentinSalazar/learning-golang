package utils

/*
Escribir una función para ingresar desde el teclado una serie de números
entre A y B y guardarlos en una lista.
En caso de ingresar un valor fuera de rango la función mostrará un mensaje
de error y solicitará un nuevo número. Para finalizar la carga se deberá ingresar -1.
La función recibe como parámetros los valores de A y B, y devuelve la lista cargada
(o vacía, si el usuario no ingresó nada) como valor de retorno.
Tener en cuenta que A puede ser mayor, menor o igual a B.

* Utilizar la función desde el ejercicio hasta el 6 para realizarlos.

*/

import "fmt"

func elegirParametros() (int, int) {
	var (
		parametroA int
		parametroB int
	)
	flag := true

	for flag {
		fmt.Println("Ingresa desde que numero se puede cargar una lista: ")
		fmt.Scan(&parametroA)
		fmt.Println("Ingresa hasta que numero: ")
		fmt.Scan(&parametroB)
		if parametroA > parametroB {
			fmt.Println("El numero 'desde' debe ser menor que el numero 'hasta'. ")
		} else {
			flag = false
		}
	}

	return parametroA, parametroB
}

func CargarLista() []int {
	var (
		sliceEnteros           []int
		parametroA, parametroB = elegirParametros()
	)

	flag := true

	for flag {
		var numeroIngresado int
		fmt.Println("Ingrese un numero entre", parametroA, "y", parametroB, "para cargar a la lista (-1 para finalizar): ")
		fmt.Scan(&numeroIngresado)
		if numeroIngresado == -1 {
			flag = false
		} else if (numeroIngresado < parametroA) || (numeroIngresado > parametroB) {
			fmt.Println("El debe estar dentro del parametro de", parametroA, "y", parametroB)
		} else {
			sliceEnteros = append(sliceEnteros, numeroIngresado)
		}
	}
	return sliceEnteros
}

func imprimirSlice(slice []int) {
	for _, value := range slice {
		fmt.Print(value, " ")
	}
}
