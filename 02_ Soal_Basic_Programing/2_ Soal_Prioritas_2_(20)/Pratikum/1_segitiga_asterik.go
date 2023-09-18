package main

import "fmt"

func main() {
	var N int

	// Input nilai
	fmt.Print("Input: ")
	fmt.Scan(&N)

	//Looping
	fmt.Println("Output: ")
	for i := 0; i < N; i++ {

		//cetak spasi
		for j := 0; j <= N-i; j++ {
			fmt.Print(" ")
		}

		//cetak bintang
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}

		//hasil
		fmt.Println()
	}
}
