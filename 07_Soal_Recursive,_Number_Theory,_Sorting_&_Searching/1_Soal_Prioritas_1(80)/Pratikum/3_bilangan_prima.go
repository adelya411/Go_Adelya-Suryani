package main

import "fmt"

//Memeriksa bilangan prima
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primeX(number int) int {
	// your code here
	primeCount := 0
	currentNum := 2  //Mulai dari bilangan prima pertama

	for primeCount < number {
		if isPrime(currentNum) {
			primeCount++
		}
		if primeCount == number {
			return currentNum
		}
		currentNum++
	}

	return -1  //Jika bilangan prima tidak ditemukan sesuai urutan
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}