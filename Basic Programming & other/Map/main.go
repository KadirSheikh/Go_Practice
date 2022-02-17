package main

import "fmt"

func main() {

	//var colors := map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"white": "#fff",
		"black": "#000",
		"red":   "#00f",
	}

	print(colors)

	//fmt.Println(colors)

}

func print(m map[string]string) {
	for key, value := range m {

		fmt.Println(key, value)

	}
}
