package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("name is empty")
    }
    return fmt.Sprintf(randomFormat(), name), nil
}
func Hellos(names []string)  (map[string]string,error) {
	// messages:=map[string]string{
	// 	"ABC":"Hegbr",
	// }
	messages:=make(map[string]string) //map ว่าง
	for _,name:=range names {
	 message,err:=Hello(name)
	 if err!= nil{
		return nil,err
	 }
	 messages[name]=message
	}
	return messages,nil
}
    
   
func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
        "Hello, %v. Welcome!",
        "Great to see you, %v",
        "Hail, %v! Well met!",
    }
    return formats[rand.Intn(len(formats))]
}