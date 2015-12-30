package main

import (
	"fmt"
)

func main() {
	port := 8080
	protocol := "tcp"
	fmt.Println("Listening on port ", port, " Protocol:", protocol)
	wait := make(chan int)
	ser := NewServer(port, protocol)
	go ser.Listen()

	//Fool the main saying someone would call :)
	<-wait

}
