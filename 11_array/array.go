package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Acla"
	names[1] = "Putra"
	names[2] = "Nurdin"
	names[3] = "Hamzah"

	var values = [3]int{
		90,
		29,
		40,
	}
	
	fmt.Println(names)
	fmt.Println(names[0]) 
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
	fmt.Println(cap(values))
}
