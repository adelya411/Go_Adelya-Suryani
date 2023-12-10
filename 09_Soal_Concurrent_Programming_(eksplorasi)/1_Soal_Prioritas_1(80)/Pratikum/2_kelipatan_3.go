package main

import (
	"fmt"
)

func printMultiplesOfThree(max int, ch chan int) {
	for i := 3; i <= max; i += 3 {
		ch <- i					//  channel yang digunakan untuk mengirim bilangan kelipatan 3
	}
	close(ch)
}

func main() {
	var max int
	fmt.Print("Masukkan nilai maksimum: ")
	_, err := fmt.Scan(&max)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ch := make(chan int)

	go printMultiplesOfThree(max, ch)		// gorountie mencetak kelipatan 3

	for num := range ch {
		fmt.Printf("%d adalah kelipatan 3\n", num)
	}
}
