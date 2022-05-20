package main

import "fmt"

var result int

func main() {
	result = operacionesMiddleware(sumar)(10, 2)
	fmt.Println(result)
	result = operacionesMiddleware(restar)(100, 99)
	fmt.Println(result)
	result = operacionesMiddleware(multiplicar)(20000, 3000)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Ingreso a la operacion")
		return f(a, b)
	}
}
