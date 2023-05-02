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