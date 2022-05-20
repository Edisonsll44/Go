package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool { return h.vivo }

func humanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

func estoyVivo(sv SerVivo) bool {
	return sv.estaVivo()
}

func ejemploListaInterfaces() []interface{} {
	return []interface{}{"Holaaaa", true, 12, 245, 12.4}
}

func main() {
	Eddy := new(hombre)
	Eddy.esHombre = true
	Eddy.vivo = true
	humanoRespirando(Eddy)

	Evelin := new(mujer)
	humanoRespirando(Evelin)

	println("Resultado si esta vivo o no: ", estoyVivo(Eddy))

	arrayInterface := ejemploListaInterfaces()
	fmt.Println(arrayInterface)
}
