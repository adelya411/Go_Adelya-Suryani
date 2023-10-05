package main

import "fmt"

func caesar(offset int, input string) string {
	// Inisialisasi string
	result := ""

	// Loop karakter dalam input
	for _, char := range input {
		// Periksa karakter huruf kecil
		if char >= 'a' && char <= 'z' {
			// Hitung pergeseran dengan mengurangkan kode ASCII huruf 'a' (97)
			// Tambahkan offset, lalu ambil modulo 26 untuk memastikan pergeseran dalam batas a-z
			newChar := 'a' + ((char - 'a' + rune(offset)) % 26)

			// Tambahkan karakter hasil pergeseran ke string hasil
			result += string(newChar)
		} else {
			// Jika karakter bukan huruf kecil, tambahkan karakter aslinya tanpa perubahan
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
