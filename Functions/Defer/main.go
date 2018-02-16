package main

import (
	"fmt"
)

func main() {
	defer world()
	hello()
}

func world() {
	fmt.Println("world")
}

func hello() {
	fmt.Println("Hello")
}
