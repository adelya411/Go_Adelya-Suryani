package main

import (
	"fmt"
)

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			result += romans[i]
		}
	}
	return result
}

func main() {
	inputs := []int{4, 9, 23, 2021, 1646}

	for _, num := range inputs {
		roman := intToRoman(num)
		fmt.Printf("Input: %d\nOutput: %s\n\n", num, roman)
	}
}
