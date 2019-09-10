package main

import (
	"fmt"

	"github.com/kurabi1966/GoDesignPatterns/hello_world/hello"
)

func main() {
	number := 5
	pointer_to_number := &number

	fmt.Println(hello.SayHello())
	fmt.Println(pointer_to_number)
	fmt.Println(*pointer_to_number)
}
