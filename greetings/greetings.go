package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string,error) {
	if name == ""{
		return "",errors.New(("Name is empty"))
	}
	return fmt.Sprintf("Hello, %v Welcome!", name),nil
}