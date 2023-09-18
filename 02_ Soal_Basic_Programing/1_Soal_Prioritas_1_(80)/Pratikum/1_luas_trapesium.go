package main

import "fmt"

func main() {
	var a, b, t float32

	//Input panjang sisi atas trapesium
	fmt.Print("Masukkan panjang sisi atas trapesium = ")
	fmt.Scanln(&a)

	//Input panjang sisi bawah trapesium
	fmt.Print("Masukkan panjang sisi bawah trapesium = ")
	fmt.Scanln(&b)

	//Input tinggi trapesium
	fmt.Print("Masukkan tinggi trapesium = ")
	fmt.Scanln(&t)

	//Hitung luas trapesium
	L := 0.5 * t * (a + b)
	fmt.Print("Luas Trapesium = ", L)
}
