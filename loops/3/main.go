/*
 Desarrollar un programa que imprima la suma de los números impares compren- didos entre 42 y 176.
*/

package main

import "fmt"

func main() {
	var suma int

	for i := 42; i <= 176; i++ {
		if i%2 != 0 {
			suma += i
		}
	}

	fmt.Println("Suma: ", suma)
}
