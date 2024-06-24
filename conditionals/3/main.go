package main

/*
  Una editorial determina el precio de un libro según la cantidad de
  páginas que contiene.
  El costo básico del libro es de $500, más $3,20 por página
  con encuadernación rústica.
  Si el número de páginas supera las 300 la encuadernación debe ser en tela,
  lo que incrementa el costo en $200.
  Además, si el número de páginas sobrepasa las 600 se hace necesario
  un procedimiento especial de en- cuadernación que incrementa
  el costo en otros $336.
  Desarrollar un programa que calcule el costo de un libro
  dado el número de páginas.
*/

import "fmt"

const (
	costoBasicoLibro            = 500.0
	costoPagina                 = 3.20
	costoEncuadernacionTela     = 200.0
	costoEncuadernacionEspecial = 336
)

func determinarCostoLibro(paginas int) float32 {
	var costoFinal float32

	if paginas <= 300 {
		costoFinal = (float32(paginas) * costoPagina) + costoBasicoLibro
	} else if paginas > 300 && paginas <= 600 {
		costoFinal = (float32(paginas) * costoPagina) + costoBasicoLibro + costoEncuadernacionTela
	} else if paginas > 600 {
		costoFinal = (float32(paginas) * costoPagina) + costoBasicoLibro + costoEncuadernacionTela + costoEncuadernacionEspecial
	}

	return costoFinal
}

func main() {
	var cantidadPaginas int
	fmt.Print("Ingresa la cantidad de paginas del libro: ")

	_, error := fmt.Scan(&cantidadPaginas)

	if error != nil {
		fmt.Println("Hubo un error en el ingreso de las paginas.")
	}

	fmt.Println("Costo del libro con", cantidadPaginas, "paginas es: ", "$", determinarCostoLibro(int(cantidadPaginas)))
}
