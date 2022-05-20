package main

import "fmt"

var calculos func(int, int) int = func(i1 int, i2 int) int { return i1 + i2 }

func main() {
	fmt.Println("Suma: ", calculos(5, 7))
	calculos = func(i1 int, i2 int) int {
		return i1 - i2
	}
	fmt.Println("Resta: ", calculos(9, 2))

	calculos = func(i1 int, i2 int) int {
		return i1 * i2
	}
	fmt.Println("Multiplicación: ", calculos(4, 9))

	calculos = func(i1, i2 int) int {
		return i1 / i2
	}
	fmt.Println("División:", calculos(100, 10))

	operaciones()
}

func operaciones() {
	resultado := func() int {
		var a int = 123
		var b int = 123
		return a + b
	}
	fmt.Println(resultado())

	tablaDel := 2
	miTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}

}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
