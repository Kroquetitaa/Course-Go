package main

import "fmt"

func normalFunction(message string) {
	fmt.Println("Hola mundo")
	fmt.Println("Hola mundo 2")
	fmt.Println("Hola mundo 3")
}

func tripleArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c int, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")
	value := returnValue(2)
	fmt.Println("Value: ", value)

	value, value2 := doubleReturn(3)
	fmt.Println(value, value2)
}
