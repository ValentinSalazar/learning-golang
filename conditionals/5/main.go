package main

/*
Diseñar un programa que calcule y muestre el sueldo neto de un empleado en base a su
sueldo básico y su antigüedad en años.
Si es soltero se le incrementa el sueldo en 5% del salario bruto por cada año de antigüedad,
mientras que si es casado se le incrementa el sueldo en 7% del bruto por cada año de antigüedad.
También se le realizan los siguientes descuentos: Jubilación: 11%, Obra Social: 3%, Sindicato: 3%

# Sueldo Basico: El ingresado por el usuario.
# Sueldo Bruto: Sueldo basico con aumentos aplicados.
# Sueldo Neto: Sueldo bruto con descuentos aplicados.

*/

import "fmt"

const (
	soltero    = 0.05
	casado     = 0.07
	jubilacion = 0.11
	obraSocial = 0.03
	sindicato  = 0.03
)

func solicitarDatos() (float32, int, int) {
	var sueldo float32
	fmt.Print("Ingrese su sueldo: ")
	_, error := fmt.Scan(&sueldo)
	if error != nil {
		fmt.Println("Hubo un error en el ingreso del sueldo.")
	}

	var antiguedad int
	fmt.Print("Ingrese su antiguedad en años: ")
	_, error = fmt.Scan(&antiguedad)
	if error != nil {
		fmt.Println("Hubo un error en el ingreso de la antiguedad.")
	}

	var estadoCivil int
	fmt.Print("Ingrese su estado civil (1 solero; 2 casado): ")
	_, error = fmt.Scan(&estadoCivil)

	if error != nil {
		fmt.Print("Hubo un error en el ingreso del estado civil.")
	}

	return sueldo, antiguedad, estadoCivil
}

func calcularSueldoBruto(sueldoBasico float32, antiguedad int, estadoCivil int) float32 {
	var sueldoBruto float32

	if estadoCivil == 1 {
		sueldoBruto = sueldoBasico + ((sueldoBasico * soltero) * float32(antiguedad))
	} else {
		sueldoBruto = sueldoBasico + ((sueldoBasico * casado) * float32(antiguedad))
	}

	return sueldoBruto
}

func calcularSueldoNeto(sueldoBruto float32) float32 {
	descuentos := (sueldoBruto * jubilacion) + (sueldoBruto * obraSocial) + (sueldoBruto * sindicato)

	return sueldoBruto - descuentos
}

func main() {
	sueldoIngresado, antiguedadIngresada, estadoCivilIngresado := solicitarDatos()
	sueldoBruto := calcularSueldoBruto(sueldoIngresado, antiguedadIngresada, estadoCivilIngresado)
	sueldoNeto := calcularSueldoNeto(sueldoBruto)

	fmt.Println("El sueldo bruto es: ", sueldoBruto)
	fmt.Println("El sueldo neto es: ", sueldoNeto)
}
