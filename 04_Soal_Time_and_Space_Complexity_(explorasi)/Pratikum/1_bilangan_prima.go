package main

import "fmt"

func primeNumber(number int) bool {    //Bilangan prima n % 1 == 0, n % n == 0

//   fmt.Print("Input: ")
//   fmt.Scanln(&number)

   if number == 2 {
   //   fmt.Println("Output : Bilangan Prima")
      return true
   }

   if number % 2 == 0 || number <= 1 {
   //   fmt.Println("Output : Bilangan Bukan Prima")
      return false
   }

   for i := 3; i * i <= number; i += 2 {  //Dibagi dari 2 sampai number yang diinputkan
      if number % i == 0 {
   //      fmt.Println("Output : Bilangan Bukan Prima")
         return false
      }
   }
   //fmt.Println("Output : Bilangan Prima")
   return true
}


func main() {
   fmt.Println(primeNumber(1000000007)) // true
   fmt.Println(primeNumber(13))         // true
   fmt.Println(primeNumber(17))         // true
   fmt.Println(primeNumber(20))         // false
   fmt.Println(primeNumber(35))         // false
}