package main

import (
	"fmt"
	"strings"
	"sync"
)

func countFrequency(text string, ch chan map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()

	frequencyMap := make(map[rune]int)
	for _, char := range text {
		if 'a' <= char && char <= 'z' {
			frequencyMap[char]++
		}
	}
	ch <- frequencyMap
}

func main() {
	text := "Aku bermain bersama mereka seharian penuh" // Teks yang ingin Anda hitung frekuensinya
	numWorkers := 4                                      // Jumlah worker goroutine yang ingin Anda gunakan

	text = strings.ToLower(text)
	segmentSize := len(text) / numWorkers

	var wg sync.WaitGroup
	resultCh := make(chan map[rune]int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		start := i * segmentSize
		end := start + segmentSize
		if i == numWorkers-1 {
			end = len(text)
		}
		go countFrequency(text[start:end], resultCh, &wg)
	}

	wg.Wait()
	close(resultCh)

	frequencyMap := make(map[rune]int)

	for partialMap := range resultCh {
		for char, count := range partialMap {
			frequencyMap[char] += count
		}
	}

	fmt.Println("Frekuensi huruf dalam teks:")
	for char, count := range frequencyMap {
		fmt.Printf("%c: %d\n", char, count)
	}
}
