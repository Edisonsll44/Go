package main

import "fmt"

func main() {
	variante2()
	variante3()
	slice := varianteSlice()
	nuevoSlice := []int{100, 200, 300, 400}
	slice = append(slice, nuevoSlice...)
	fmt.Println(slice)

}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[2:4]
	fmt.Println(porcion)
}

func variante3() {
	//asigna una cantidad de memoria, para un arreglo que inica con 5 elementos y luego puede subir a 20
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

//agregar un elemento a un slice
func varianteSlice() []int {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	return nums
}
