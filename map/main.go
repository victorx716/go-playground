package main

import "fmt"

func main() {
	// colors := make(map[string]string)
	var colors = map[string]string{
		"red":   "#ff0000",
		"green": "#gffre",
	}

	// colors[10] = "dfdgfgfg"
	// delete(colors, 10)
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " has code of " + hex)
	}
}
