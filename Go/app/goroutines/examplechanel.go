package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}
func main() {
	c := make(chan string, 1)
	fmt.Println("Hola")
	go say("bye", c)
	result := <-c

	fmt.Println(result)

}
