package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "$4bf745",
	}
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	printMap(colors)

	var colors1 map[string]string
	colors1 = make(map[string]string)
	colors1["white"] = "#ffffff"
	fmt.Println(colors1)

	colors2 := make(map[int]string)
	colors2[10] = "Hi"
	fmt.Println(colors2)
	delete(colors2, 10)
	fmt.Println(colors2)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, " : ", hex)
	}
}
