package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAcla NoKTP = "2617339927873900"
	var marriedStatus Married = false
	fmt.Println(noKtpAcla)
	fmt.Println(marriedStatus)
}