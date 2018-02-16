package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(name, lastname string) string {
	return fmt.Sprint(name, lastname)
}
