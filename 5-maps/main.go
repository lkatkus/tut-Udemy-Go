package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	colors["white"] = "#ffffff"

	fmt.Println(colors)
	fmt.Println(colors["red"])
	fmt.Println(colors["green"])
	fmt.Println(colors["blue"])
	fmt.Println(colors["white"])

	delete(colors, "white")
	fmt.Println(colors)

	var otherColors map[string]string
	fmt.Println(otherColors)

	yetAnotherColors := make(map[string]string)
	fmt.Println(yetAnotherColors)

	goodColor := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	printMap(goodColor)
}

func printMap(m map[string]string) {
	fmt.Println("------")

	for key, value := range m {
		fmt.Println("Hex code for", key, "is", value)
	}

	fmt.Println("------")
}
