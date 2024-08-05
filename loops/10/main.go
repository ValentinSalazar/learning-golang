package main

/* Realizar la sucesión de Fibonacci hasta el numero N, el cual se
* ingresa por teclado. */

import "fmt"

func main() {
	var n int
	fmt.Println("Ingrese un numero natural para calcular su tribonacci: ")
	fmt.Scan(&n)
	a := 0
	b := 1
	c := 2

	if n == 0 {
		fmt.Print(a)
	} else if n == 1 {
		fmt.Print(a, " ", b)
	} else {
		fmt.Print(a, " ", b, " ", c, " ")
		for c < n {
			fib := a + b + c
			a = b
			b = c
			c = fib

			if c < n {
				fmt.Print(fib, " ")
			}
		}
	}
}
