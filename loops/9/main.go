package main

/*
La sucesión de Fibonacci es una sucesión de números enteros donde cada término
se obtiene como suma de los dos anteriores, siendo los dos primeros 0 y 1.
Por lo tanto, Fib=0, 1, 1, 2, 3, 5, 8, 13, 21....
Realizar un programa que lea N e imprima los N primeros términos de esta sucesión,
como así también la suma de los mismos.
*/

import "fmt"

func main() {
	var n int
	fmt.Println("Ingrese un numero para calcular su sucesión Fibonnaci: ")
	fmt.Scan(&n)
	a := 0
	b := 1

	if n == 0 {
		fmt.Print(0)
	} else if n == 1 {
		fmt.Print(1)
	} else {
		fmt.Print(a, " ")
		fmt.Print(b, " ")
		for b < n {
			fib := a + b
			a = b
			b = fib
			fmt.Print(fib, " ")
		}
	}
}
