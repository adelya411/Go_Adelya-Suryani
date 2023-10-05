package main

import (
	"fmt"
	"strconv"
)

func hitungBiner(num int) int {
	sum := 0
	for num > 0 {
		if num%2 == 1 {
			sum++
		}
		num /= 2
	}
	return sum
}

func representasiBiner(n int) []string {
	result := make([]string, n+1)
	
	for i := 0; i <= n; i++ {
		result[i] = strconv.FormatInt(int64(i), 2)
	}
	return result
}

func main() {
	n1 := 2
	fmt.Printf("Input: n = %d\n", n1)
	fmt.Printf("Output: %v\n", representasiBiner(n1))

	n2 := 3
	fmt.Printf("Input: n = %d\n", n2)
	fmt.Printf("Output: %v\n", representasiBiner(n2))
}
