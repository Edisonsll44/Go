package main

import (
	"fmt"
)

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(2)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(2, 23, 45, 68))
	fmt.Println(calculo(5))
	fmt.Println(calculo(5, 99, 100, 101))
	fmt.Println(calculo(5, 46, 2, 45, 3567, 8))
	fmt.Println(calculo(5, 46, 2, 88, 900))
}

func uno(numero int) (z int) {
	z = numero * 10
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 1, true
	}
}

func calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Println("\n Item\n", item)
	}
	return total
}
