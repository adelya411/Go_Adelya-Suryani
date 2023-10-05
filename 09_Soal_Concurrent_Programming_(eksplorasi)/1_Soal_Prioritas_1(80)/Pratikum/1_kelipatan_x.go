package main

import (
	"fmt"
	"time"
)

// Fungsi untuk mencetak angka kelipatan x
func cetakKelipatan(x int) {
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Printf("%d adalah kelipatan dari %d\n", i, x)
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var x int
	fmt.Print("Masukkan nilai x (bilangan bulat positif): ")
	_, err := fmt.Scanf("%d", &x)
	if err != nil || x <= 0 {
		fmt.Println("Masukan tidak valid. Harap masukkan bilangan bulat positif.")
		return
	}

	// Menerapkan goroutine untuk mencetak angka kelipatan x
	go cetakKelipatan(x)

	// Agar program tidak berhenti segera,
	fmt.Println("Program berjalan. Tekan Ctrl+C untuk menghentikan program.")
	for {
		time.Sleep(1 * time.Second) // Tidur selama 3 detik atau sampai program dihentikan.
	}
}
