package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	// your code here
	for x := -10000; x <= 10000; x++ {
		for y := -10000; y <= 10000; y++ {
			if x != y {
				z := a - x - y
				if z != x && z != y && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Printf("%d %d %d\n", x, y, z)
					return
				}
			}
		}
	}
	fmt.Println("No solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
