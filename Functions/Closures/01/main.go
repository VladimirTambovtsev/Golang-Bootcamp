package main

import (
	"fmt"
)

func makeGreet() func() string {
	return func() string {
		return "Hello world"
	}
}

func main() {
	greet := makeGreet()
	fmt.Println(greet())
	fmt.Println(greet)
}
