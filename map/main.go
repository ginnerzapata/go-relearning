package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#00FF00",
	}

	people := map[string]person{
		"id1": {
			name: "ginner",
			age:  36,
		},
	}

	fmt.Println(colors, people)
}
