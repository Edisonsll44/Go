package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}
func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))
	//range and close
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje1", email1)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido desde email", m1)

		case m2 := <-email2:
			fmt.Println("Email recibido desde email", m2)
		}
	}

}
