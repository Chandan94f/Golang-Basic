package main

import "fmt"

func fMap() {
	// 1. var colors := map[string]string

	//2 . colours := make(map[string]string)

	// 1 & 2 are similar, initialise now , assigne later

	// 3

	colours := map[string]string{
		"red":   "ff",
		"green": "gg",
	}

	colours["white"] = "fffff"

	// fmt.Println(colours)

	// delete(colours, "green")

	// fmt.Println(colours)

	printMap(colours)

}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
