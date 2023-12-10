package main

import (
	"fmt"
	"sync"
)

func performCalculation(arr []int, index int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()         // Mengunci mutex sebelum mengakses array bersama
	defer mu.Unlock() // Membuka mutex setelah selesai mengakses array bersama

	// Melakukan perhitungan pada elemen array (contoh: menggandakan nilainya)
	arr[index] = arr[index] * 2
}

func main() {
	arraySize := 10
	arr := make([]int, arraySize)

	var wg sync.WaitGroup
	var mu sync.Mutex

	// Inisialisasi array dengan nilai acak
	for i := 0; i < arraySize; i++ {
		arr[i] = i
	}

	// Menampilkan array sebelum perhitungan
	fmt.Println("Array awal:", arr)

	for i := 0; i < arraySize; i++ {
		wg.Add(1)
		go performCalculation(arr, i, &wg, &mu)
	}

	wg.Wait()

	// Menampilkan array setelah perhitungan
	fmt.Println("Array setelah perhitungan:", arr)
}
