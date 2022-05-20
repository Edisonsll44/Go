package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	ejemploPanic()
	ejemploDefer()

}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor 1")
	}
}
func ejemploDefer() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)
	defer f.Close()

	if err != nil {
		fmt.Println("error abriendo el archivo")
		os.Exit(1)
	}
}
