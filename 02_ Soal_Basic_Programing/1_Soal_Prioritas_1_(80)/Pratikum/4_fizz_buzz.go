package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {

		if i%3 == 0 {	// operator % => nyari sisa
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(" ", i)
		}
	}
}
