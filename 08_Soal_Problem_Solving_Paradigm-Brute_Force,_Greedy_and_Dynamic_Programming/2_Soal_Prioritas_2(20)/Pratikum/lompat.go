package main

import "fmt"

func Frog(jumps []int) int {
	// your code here
	n := len(jumps)
	if n <= 1 {
		return 0
	}

	cost := make([]int, n)
	cost[0] = 0

	for i := 1; i < n; i++ {
		diff1 := jumps[i] - jumps[i-1]
		cost1 := cost[i-1] + abs(diff1)

		if i >= 2 {
			diff2 := jumps[i] - jumps[i-2]
			cost2 := cost[i-2] + abs(diff2)
			cost[i] = min(cost1, cost2)
		} else {
			cost[i] = cost1
		}
	}

	return cost[n-1]
}

//Mencari nilai minimum diantara 2 angka
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//menghitung nilai mutlak dari suatu angka
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
