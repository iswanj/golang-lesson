package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"
	// colors["red"] = "#ff0000"

	// fmt.Println(colors)

	// delete(colors, "red")

	//==================

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["yellow"] = "#FFFF00"
	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
