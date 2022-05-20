package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	miNombreLento("Edison Sánchez")
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
