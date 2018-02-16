package main

import (
	"fmt"
)

func main() {
	greet("Jane", "Brown")
	greet("Bob", "Green")
}

func greet(name, lastname string) {
	fmt.Println(name, lastname)
}
