package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	canal0 := make(chan time.Duration)
	go miNombreLento("Edison Sánchez Llerena", canal0)
	msg0 := <-canal0
	fmt.Println(msg0)

	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")
	msg := <-canal1
	fmt.Println(msg)

}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 10000000; i++ {

	}
	final := time.Now()
	canal1 <- final.Sub(inicio)
}

func miNombreLento(nombre string, canal0 chan time.Duration) {
	inicio := time.Now()
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(10000 * time.Millisecond)
		fmt.Printf(letra)
	}
	final := time.Now()
	canal0 <- final.Sub(inicio)
}
