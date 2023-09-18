package main

import "fmt"

func main() {
	var nilai int

	// Input nilai
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	//Menentukan nilai sesuai ketentuan
	if nilai < 0 || nilai > 100 {
		fmt.Println("Nilai Invalid")
	} else {
		if nilai >= 80 && nilai <= 100{
			fmt.Println("Nilai 80 - 100 : A")
		} else if nilai >= 65 {
			fmt.Println("Nilai 65 - 79 : B")
		} else if nilai >= 50 {
			fmt.Println("Nilai 50 - 64 : C")
		} else if nilai >= 35 {
			fmt.Println("Nilai 35 - 49 : D")
		} else {
			fmt.Println("Nilai 0 - 34 : E")
		}
	}
}
