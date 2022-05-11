package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, world.")

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	err := errors.New("invalid APIVersion value")
	panic(err)
}
