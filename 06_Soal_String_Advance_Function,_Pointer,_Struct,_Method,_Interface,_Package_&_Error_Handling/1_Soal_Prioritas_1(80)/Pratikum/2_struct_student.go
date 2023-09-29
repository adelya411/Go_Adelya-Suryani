package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	return 0
}

func (s Student) Min() (min int, name string) {
	return 0, ""
}

func (s Student) Max() (max int, name string) {
	return 99, ""
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input  " + string(i) + " Student’s Name :  ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input  " + name + " Score :  ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is ", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is :  "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is :  "+nameMin+" (", scoreMin, ")")
}