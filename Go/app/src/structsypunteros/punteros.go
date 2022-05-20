package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

//sobre escribir el string
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de disco y es de marca %s", myPc.ram, myPc.disk, myPc.brand)
}
func main() {
	a := 50
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = 1000
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	myPc.duplicateRam()
	fmt.Println(myPc)

	myPc.duplicateRam()
	fmt.Println(myPc)
}
