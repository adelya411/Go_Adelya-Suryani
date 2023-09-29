package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	// your code here
	result := make(map[string]int) // menghitung kemunculan angka

	for _, nomor := range angka { //hitung muncul angka yang diinput
		nomorStr := string(nomor)
		result[nomorStr]++
	}

	//Slice nyimpan angka munculsekali
	var nomorUnik []int

	for _, nomor := range angka {
		nomorStr := string(nomor)

		if result[nomorStr] == 1 {
			nomorInt := int(nomor - '0')
			nomorUnik = append(nomorUnik, nomorInt)
		}
	}

	return nomorUnik
}

func main() {
	// Test cases
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
