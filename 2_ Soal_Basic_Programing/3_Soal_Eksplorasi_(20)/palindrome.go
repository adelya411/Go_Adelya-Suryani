package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	reader := bufio.NewReader(os.stdin)
	input, _ := reader.ReadString('\n')
	
	var result string
	for i := len(input) - 1; i >= 0; i-- {
		result = result + string(input[i])
	}

	if input == result {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}

// import (
// 	"bufio"
// 	"fmt"
// 	"strings"
// )

// func isPalindrome(input string) bool {
// 	for i := 0; i < len(input)/2; i++ {
// 		if input[i] != input[len(input)-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	var kata string

// 	//Inputkan kata
// 	fmt.Println("Apakah Palindome?")
// 	fmt.Print("Masukkan kata : ")
// 	fmt.Scan(&kata)
// 	result := isPalindrome(kata)

// 	//Menghapus spasi untuk mencetak
// 	kata = strings.ReplaceAll(kata, " ", " ")

// 	//Memeriksa
// 	if result == true {
// 		fmt.Println("Captured : ", kata)
// 		fmt.Println("Palindrome")
// 	} else {
// 		fmt.Println("Captured : ", kata)
// 		fmt.Println("Bukan Palindrome")
// 	}
// }
