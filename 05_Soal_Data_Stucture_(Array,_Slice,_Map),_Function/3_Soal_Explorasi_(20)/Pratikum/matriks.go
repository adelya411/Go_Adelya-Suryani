package main

import (
	"fmt"
	"math"
)

func main() {
	matriks := [][]int{ //input matriks
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	//jumlah matrik dari kiri ke kanan
	diagonalkirikanan := 0
	diagonalkanankiri := 0

	//looping untuk akses dan hitung
	for i := 0; i < len(matriks); i++ {
		diagonalkirikanan += matriks[i][i]
		diagonalkanankiri += matriks[i][len(matriks)-1-i]
	}

	//selisih jumlah diagonal
	mutlak := int(math.Abs(float64(diagonalkirikanan - diagonalkanankiri)))

	fmt.Printf("Diagonal kiri ke kanan = %d\n", diagonalkirikanan)
	fmt.Printf("Diagonal kanan ke kiri = %d\n", diagonalkanankiri)
	fmt.Printf("Perbedaan mutlak = %d\n", mutlak)
}
