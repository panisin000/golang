package main

import (
	"fmt"
	"log"

	"github.com/panisin000/golang/greetings"
)

func main() {
	message,err:= greetings.Hello("Name")
	if err!= nil{
		log.Fatal(err)
	}
    
	greetings.randomFormat()
	fmt.Println(message)
}