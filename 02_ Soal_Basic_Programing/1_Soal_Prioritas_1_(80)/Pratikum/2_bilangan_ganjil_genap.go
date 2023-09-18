package main

import (
	"fmt"
)

func main() {
	var bilangan int

	// Input bilangan
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&bilangan)

	// Menentukan apakah bilangan ganjil atau genap
	if bilangan%2 == 0 {
		fmt.Printf("%d adalah bilangan genap.\n", bilangan)
	} else {
		fmt.Printf("%d adalah bilangan ganjil.\n", bilangan)
	}
}
