package main

import (
	mypk "est14/0.0-fundamentals/mypackage"
	"fmt"
	"time"
)

type usuarioDto struct {
	mypk.Usuario
}

func main() {
	nuevoUsuario := new(usuarioDto)
	nuevoUsuario.AltaUsuario(1, "Edison", time.Now(), true)
	fmt.Println(nuevoUsuario)

	var myCar mypk.CarPublic
	myCar.Brand = "Ferrary"
	fmt.Println(myCar.Brand)
}
