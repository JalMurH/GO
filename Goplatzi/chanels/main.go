package main

import (
	"fmt"
)

//con el simbolo <- se esta diciendo que en esta funcion el canal esta establecido como solo entrada
//para decir que el canal sera de solo salida c<- chan string y cambiar el interior de la funcion
// text = <- c para asignarle a text el contenido del canal
// func say(text string, c chan<- string)  {
// 	c <- text
// }
// func main() {
// 	c := make(chan string, 1)

// 	fmt.Println("hello")

// 	go say("bye", c) {

//		}()
//	}
func message(text string, c chan string) {
	c <- text

}

func main() {
	c := make(chan string, 2)
	c <- "message 1"
	c <- "message 2"

	fmt.Println(len(c), cap(c))
	//cierra el canal
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	//select
	email := make(chan string)
	email2 := make(chan string)

	go message("message1", email)
	go message("message2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("Email recibido de canal email", m1)

		case m2 := <-email2:
			fmt.Println("Email recibido de canal email2", m2)
		}
	}

}
