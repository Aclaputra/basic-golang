package main

import "fmt"

func main() {

	personData := map[string]string{
		"name": "Acla",
		"address": "Tangerang",
	}

	fmt.Println(personData)
	fmt.Println(personData["name"])
	fmt.Println(personData["address"])
}