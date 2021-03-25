package main

import (
	"fmt"

	"creating-a-module/greetings"
)

func main() {
	message := greetings.Hello("Jonathan")
	fmt.Println(message)
}
