package main

import (
	"fmt"
)

func Compare(a, b string) string {
	// your code here
	var compareSubs string

	for i := 0; i < len(a); i++ {
		for j := i + 1; j <= len(a); j++ {
			substring := a[i:j]
			if len(substring) > len(compareSubs) && b != "" && len(b) >= len(substring) && (substring == b[:len(substring)]) {
				compareSubs = substring
			}
		}
	}

	return compareSubs
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA
}