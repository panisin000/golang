package main

import (
	"fmt"
	"log"
	"github.com/panisin000/golang"
)

func main() {
	message,err:= greetings.Hello("Name")
	if err!= nil{
		log.Fatal(err)
	}
    fmt.Println(message)
}