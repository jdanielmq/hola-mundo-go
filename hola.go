package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaracion de constante
const PI float32 = 3.14

const (
	x = 100
	y = 0b1010 // valor binario
	z = 0o12   // ocatal
	w = 0xFF   //hexadecimal
)

const (
	Domiengo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("hola mundo")
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())

	//Declaracion de variables
	var firstName, lastName string
	var age int

	firstName = "Juan"
	lastName = "Daniel"
	age = 44

	fmt.Println("primer nombre : ", firstName)
	fmt.Println("Segundo nombre : ", lastName)
	fmt.Println("años : ", age)
	fmt.Println("###############################")

	//otro ejemplo como imprimir todas las variables en una sola linea
	fmt.Println(firstName, lastName, age)
	fmt.Println("###############################")

	// otra forma de declaracion de variables

	var (
		nombre, apellido string
		annos            int
	)
	nombre = "Juan Daniel"
	apellido = "Muñoz Queupul"
	annos = 44
	fmt.Println(nombre, apellido, annos)
	fmt.Println("###############################")

	// como declarar variables con datos sin espicificar el tipo de dato
	var (
		nombres   = "Juan Daniel"
		apellidos = "Muñoz Queupul"
		years     = 44
	)
	fmt.Println(nombres, apellidos, years)
	fmt.Println("###############################")

	var primerNombre, lastPrimerNombre, yearsPrimer = "Juan", "daniel", 44

	fmt.Println(primerNombre, lastPrimerNombre, yearsPrimer)
	fmt.Println("###############################")

	//declaracion de variable SOLO dentro de la funcion, esto solo sirve dentro de la funcion no para exterior
	primer, segundo := "Prueba Uno", "prueba dos"
	tercero := 30

	fmt.Println(primer, segundo, tercero)
	fmt.Println("###############################")

	//imprimir las varibales contantes
	fmt.Println(PI, x, y, z, w)
	fmt.Println("###############################")

	fmt.Println(Domiengo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)
	fmt.Println("###############################")

	// colocar un alias a un string.
	fullname := "juan Daniel \t(alias \"tonijua\")\n"

	fmt.Println(fullname)
	fmt.Println("###############################")

	// tipo de datos

	var a byte = 'a'
	fmt.Println(a)
	fmt.Println("###############################")

	s := "hola"
	fmt.Println(s[0])
	fmt.Println("###############################")

	var r rune = '❤'
	fmt.Println(r)
	fmt.Println("###############################")

	// vamos imprimir datos con el printF
	name := "Juan Daniel"
	edad := 44

	fmt.Printf("hola, me llamo %s y tengo %d años. \n", name, edad)
	fmt.Println("###############################")

	// vamos asignar el string concadenado o formteado, el ejercicio anterior y la asignamos en una variable para despues imprimir por consola
	greeting := fmt.Sprintf("hola, me llamo %s y tengo %d años.", name, edad)
	fmt.Println(" desplegando la variable: ", greeting)
	fmt.Println("###############################")

	//vamos ingresar los datos desde el teclado y solo escanea un valor
	var nameNew string
	var edadNew int

	fmt.Print("Ingrese el nombre completo: ")
	fmt.Scanln(&nameNew)
	fmt.Print("Ingrese la edad: ")
	fmt.Scanln(&edadNew)

	fmt.Printf("el nombre completo es %s y tiene una edad de : %d \n", nameNew, edadNew)
	fmt.Println("###############################")

}
