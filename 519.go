package main

import (
	"fmt"
)

func main() {

	noreturn()

}

func noreturn() {
	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Printf("Thank you\n")
		default:
			fmt.Println("Hello world 2")
		}
	}()
	panic("Hello world!")
}
