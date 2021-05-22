package main

import "fmt"

func main() {
	names := [...]string{
		"Acla",
		"Putra",
		"Nurdin",
		"Hamzah",
		"Cringe",
	}
	slice1 := names[2:4]
	sliceall := names[:]

	slice2 := append(slice1, "Cool abiez")

	fmt.Println(slice1)
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])
	fmt.Println(sliceall)

	fmt.Println(cap(slice1))
	fmt.Println(cap(sliceall))
	fmt.Println(slice2)

	// Newslice
	newSlice := make([]string, 3, 5)

	newSlice[0] = "Acla"
	newSlice[1] = "Putra"
	newSlice[2] = "Anjay"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)
}