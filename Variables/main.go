//Package main, paquete principal que permite ejecutar mi aplicacion
//de variables

package main

import "fmt"

func main() {

	var dog, cat = "🐶", "🐱"

	// Shorthand para declarar una variable
	zorro := "zorrito"

	/*Creando una constante, que tiene un valor unico*/
	e1, e2 := 2, 2
	//Multiplicando numeros
	multi := e1 * e2
	const unico = 3.14

	//Consola de impresion de variables
	fmt.Println(dog, cat, zorro, unico)
	fmt.Println(multi)
}
