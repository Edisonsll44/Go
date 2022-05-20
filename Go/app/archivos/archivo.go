package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	crearArchivo()
	grabarArchivo()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Error al leer archivo.txt")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Error al leer archivo.txt")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Println("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}
func crearArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintf(archivo, "Esta es una nueva linea")
	archivo.Close()
}
func grabarArchivo() {
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\n Esta es la segunda linea") == false {
		fmt.Println("Error en la segunda linea")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un Error", err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un Error")
		return false
	}
	return true
}
