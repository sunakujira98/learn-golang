package main

import "fmt"

func main() {
	// ==== 49. Whats a Map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "4bf745",
	}

	// ==== 50. Manipulating Map
	// var colors map[string]string

	// colors := make(map[int]string)

	// colors[10] = "#ffffff"

	// delete(colors, 10)

	// fmt.Println(colors)

	// Iterating over map
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
