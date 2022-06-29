package main

import "fmt"

// Muestra el area del rectangulo, circulo,trapecio

func areaRectangulo(largo float32, ancho float32) float32 {
	return largo * ancho
}

func areaCirculo(pi float32, radio float32) float32 {
	return (radio * radio) * pi
}

func areaTrapecio(base1 float32, base2 float32, altura float32) float32 {
	result := ((base1 + base2) * altura) / 2
	return result
}

func main() {
	fmt.Println("El area del Rectangulo: ", areaRectangulo(2, 3))
	fmt.Println("El area del circulo: ", areaCirculo(3.14, 5))
	fmt.Println("El area del Trapecio: ", areaTrapecio(1, 1, 4))
}
