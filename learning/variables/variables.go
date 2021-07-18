package main

import "fmt"

func main() {
	var number int
	number = 15
	fmt.Println(number)

	number = 35
	fmt.Println(number)

	name := "Nacho"
	fmt.Println(name)

	var nombre string
	var numero int
	nombre, numero = "Nacho", 35
	fmt.Println(nombre, numero)

	var nombre2 string
	var numero2 string

	nombre2, numero2 = "nacho", "35"
	fmt.Println(nombre2, numero2)

	var (
		el_nombre = "Nacho"
		el_numero = 35
	)

	fmt.Println(el_nombre, el_numero)
	// scope

	var edad = 35
	fmt.Println("la edad de nacho es:", edad)

}
