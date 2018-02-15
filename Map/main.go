package main

import (
	"fmt"
)

func main() {

	// colours := make(map[string]string) //	empty Map

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	// colours["white"] = "#ffffff"
	// colours["black"] = "#000000"

	// delete(colours, "black")

	colours["yellow"] = "#ffff00"

	printMap(colours)
	fmt.Print(colours)
}

func printMap(c map[string]string) {
	for eachColor, hex := range c {
		fmt.Println("Hex Code for given ", eachColor, "  is: ", hex)
	}
}
