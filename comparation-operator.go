package main

import "fmt"

func main() {
	var name1 = "Acla"
	var name2 = "Yovi"

	var result = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}