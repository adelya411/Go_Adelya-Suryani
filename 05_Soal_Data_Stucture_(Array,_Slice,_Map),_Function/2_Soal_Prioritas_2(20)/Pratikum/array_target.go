package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// your code here
	lihat := make(map[int]int)		//buat peta lacak angka

	for i, num := range arr {
		result := target - num			//Hitung selisih target

		if index, value := lihat[result]; value {
			return []int{index, i}		//Periksa selisih
		}
		lihat[num] = i	//simpan index angka
	}
	return []int{-1, -1}		//jika tidak ditemukan pasangan
}

func main() {
	// Test cases
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}