package main

import "fmt"

func main() {
	// const pi float64 = 3.14
	// const pi2 = 3.1415

	// fmt.Println("pi:", pi)
	// fmt.Println("pi2:", pi2)

	// // Declaracion de variables enteras
	// base := 12
	// var altura int = 14
	// var area int

	// fmt.Println(base, altura, area)

	// // Zero value
	// var a int
	// var b float64
	// var c string
	// var d bool

	// fmt.Println(a, b, c, d)

	// // Area cuadrado
	// const baseCuadrado = 10
	// areaCuadrado := baseCuadrado * baseCuadrado
	// fmt.Println("Area cuadrado", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)
	//Resta
	result = x - y
	fmt.Println("Resta:", result)
	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)
	//Division
	result = y / x
	fmt.Println("Division:", result)
	//Modulo
	result = y % x
	fmt.Println("Modulo", result)
	//Incremental
	x++
	fmt.Println("Incremental", x)
	//Decremental
	x--
	fmt.Println("Decremental", x)
	// Reto - calcula el area del rectangulo, del trapecio y de un circulo
	var largo int = 10
	var ancho int = 5
	areaRectangulo := largo * ancho
	fmt.Println("Area Rectangulo:", areaRectangulo)

	var base1 float32 = 5
	var base2 float32 = 4
	var altura float32 = 3
	areaTrapecio := ((base1 + base2) * altura) / 2
	fmt.Println("Area Trapecio:", areaTrapecio)

	var pi float64 = 3.14
	var radio float64 = 3
	areaCirculo := (radio * radio) * pi
	fmt.Println("Area Circulo:", areaCirculo)
}
