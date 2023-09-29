package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	// your code here
	temp := make(map[string]bool) 		//inisialissi map
	
	result := []string{}  //slice

	merge := append(arrayA, arrayB...)

	for _, value := range merge {
		if !temp[value]{
			temp[value] = true

			result = append(result, value)
		}
	}
	
	return result
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}