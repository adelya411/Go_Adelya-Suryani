package main

import (
	"fmt"
)

func generateMultiplesOfThree(max int, ch chan int) {
	for i := 3; i <= max; i += 3 {
		ch <- i
	}
	close(ch)
}

func printMultiples(ch chan int) {
	for num := range ch {
		fmt.Printf("%d adalah kelipatan 3\n", num)
	}
}

func main() {
	max := 30 // Ganti dengan nilai maksimum yang Anda inginkan
	ch := make(chan int, 5) // Membuat channel dengan buffer size 5

	go generateMultiplesOfThree(max, ch)
	go printMultiples(ch)

	// Menunggu kedua goroutine selesai
	select {}
}
