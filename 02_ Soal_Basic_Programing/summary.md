# Basic Programming

BASIC PROGRAMMING OF GOLANG

Our Rules
- Silent Mode
- Ask Question
- Go Toilet

Time Allocation
- Explanation
- Challenge
- Review


Introduction Golang
Go adalah bahasa pemrograman tujuan umum yang baru yang membuat perangkat lunak yang mudah dibuat, sederhana, andal, dan efisien


System Programming
Go adalah bahasa yang bagus untuk menulis program tingkat rendah yang menyediakan layanan ke sistem lain, yang disebut pemrograman sistem

Application Programs        |       System Programs
          E-commerce        |       APIs
        Music Player        |       Game
   Social Media Apps        |       CLI apps


History of Golang
- Late 2007  = Design
- Mid 2008   = Start
- Nov 2009   = Released As open-source
- Mar 2012   = Version 1.0 
- Feb 2019   = Version 1.12
- Sep 2022   = Version 1.19


Developed By :
- Robert Grisemer
- Rob Pike
- Ken Thompson
- Ian Lance Taylor
- Russ Cox


Kenapa harus Go?
> Menyenangkan dan membuat Pemrograman Menjadi Menyenangkan (sederhana)
> Menggabungkan yang terbaik dari keduanya
  - Bahasa yang dikompilasi secara statis (seperti C) dan bahasa Dinamis (seperti JavaScript)
  - Go memiliki nuansa yang lebih ringan dari bahasa scipting tetapi sesuai
> Sintaks yang ringan
> Konkurensi bawaan
  - Komputasi jaringan berskala besar, seperti pencarian web Google
  - Perangkat keras multi-inti
> Sumber Terbuka
> Digunakan oleh Perusahaan Besar


Installation Golang
Editor (IDE) :
- VS Code
- GoLand
- Atom
- Vim 

Installation (Via File Archice (Mac, Linux, & Windows))
- Go to https://golang.org/dl/ dan download versi terakhir
- Ikuti instruksi instalasi


Go Workspace
-terminal
Workspace/
  bin/
  pkg/
  src/

bin : berisi eksekusi biner yang dapat dieksekusi
pkg : berisi arsip paket Go
src : berisi file sumber GO


Go Modules
Bagaimana jika kita membuat proyek di luar GOPATH?
1. Membuat proyek di luar GOPATH.
2. Ope Terminal di direktori yang telah ditetapkan ke proyek ini
3. Jalankan perintah!

-terminal
$ go mod init


Compiling Process
Source Code - Compile - Run


Compiling & Run
-code editor
package main

import "fmt"

func main() {
  fmt.Println("Hello, I am Programmer")
}

-code editor
$ go build main.go
$ ./main

OUTPUT:
Hello, I am Programmer


Command Go Terminal
Command       | Function
go run        | Jalankan program dengan ekstensi .go
go build      | Mengkompilasi file program
go install    | Seperti go build dan melanjutkan proses instalasi
go test       | Untuk pengujian dengan akhiran _test.go
go get        | mengunduh paket Go


Main Syntax
-code editor
//main.go
package main    --> main package

import "fmt"      --> Import package "fmt" for input output

func main() {         --> main function golang
  fmt.Println("Hello, I am Programmer")       --> standard output using package "fmt"
}


Package "fmt"
Output:
  fmt.Printf()
  fmt.Println()
  fmt.Sprintf()
  etc

Scanning
  fmt.Scanln()

Format Verb
  %T, %v, %s, %q, etc.


Variables & Types
Variable
--> Variabel digunakan untuk menyimpan informasi dalam program komputer, variabel menyediakan cara untuk me-labeli data dengan nama deskriptif dan memiliki tipe data (boolean, numerik, string).

Data Type in Golang
                              |--- True
|--------------- Boolean -----|
|                             |--- False
|
|                                            |--- unit8 (0 to 255)
|                                            |--- uint16 (0 to 65535)
|                                            |--- uint32 (0 to 4294967295)
|                                            |--- uint64 (0 to 18446744073709551615)
|                                            |--- int8 (-128 to 128)
|                            |--- Integer ---|
|                            |               |int16 (-32768 to 32767)
|                            |               |--- int32 (-2147483648 to 2147483647)
Data Type ------- Numeric ---|               |--- int64 (-922337036854775808 to 922337036854775807)
|                            |               |--- int (same as int32/int64 based value)
|                            |               |--- rune (same as uint8)
|                            |
|                            |             |--- float32 (IEEE-754 32-bit floating-point numbers)
|                            |--- Float ---|
|                            |             |--- float64 (IEEE-754 64-bit floating-point numbers)
|                            |
|                            |               |--- complex64 (Complex numbers with float32 real and imaginary parts)
|                            |--- Complex ---|
|                                            |--- complex128 (Complex numbers with float64 real and imaginary parts)
|
|--------------- String


Variable Declaration
Variable -----> var age int
--> type
--> value <--|
--> address -|


Type Declaration
*Long
- var <variable_name> <type_data>
  --> var  name string

- var <variable_name> <type_data> = <value>
  --> var name string = "moryku"

- var <list_variable_name> <type_data>
  --> var name, gender string

- var <list_variable_name> <type_data> = <value>
  --> var name, gender string = "moryku", "L"

*Short
- <variable_name> := <value>
  --> name := "moryku"


Zero Values :
- booleans --> false
- floats ----> 0.0
- integers --> 0
- strings ---> ""


const
--> type
--> ~address~x
--> value


Single constants -----> const    pi float64 = 3.14
                          |      |     |      |
                          v      v     v      v
                       Declare  name  type  value

Multiple constans ----> const (     |-------------> declares a float64 constant with a value of 3.14
                            Pi float64 = 3.14
                            Pi2       ------------> declares the same float64 constant with a value of 3.14
                            Age = 10
                        )    |---------------------------> declares an int constant int with a value of 10


Example <\> Variable Declaration
-code editor
package main

import "fmt"

func main() {
  //declaration boolean with assignment
  var isLogin bool = false
  fmt.Println(isLogin)

  //declaration numeric integer
  var prime int = 7
  fmt.Println(prime)

  //declaration decimal number
  var decimal float64 = 45.6
  fmt.Println(decimal)

  //declaration string
  var hello string = "Hello World!"
  fmt.Println(hello)

  //declaration const
  const pi = 3.14159265358979323846 
  fmt.Println(pi)
}


Operation
1 + 2 -> Operand
| |----> Operator
-------> Operand

Arithmetic Operators
+ Addition
- Subtraction
/ Division
* Multiplication
% Modulo
++ Increment
-- Decrement


Example <\> Using Arithmetic Operator
-code editor
pacakge main

import "fmt"
func main(){
  //number operation
  a := 10
  t := 15
  L := (a * t) / 2
  fmt.Println(L)

  //string operation
  helloworld := "hello" + " " + "world"
  fmt.Println(helloworld)
}


Operator In Golang
- Arithmetic
  --> ex : +, -, *, /, %, ++, --

- Comparison
  --> ex : ==, !=, >, <, >=, <=

- Logical
  --> ex : &&, ||, !

- Bitwise
  --> ex : &, |, ^, <<, >>

- Assigment
  --> ex : =, +=, -=, *=, /=, %=, <<=, >>=. &=, ^=, |=

- Miscellaneous
  --> ex : &, * (Pointer)


Example :
-code editor
package main
import "fmt"

func main(){
  a := 1 * 2    // Arithmetic
  b := a > 0    // Relational & Assignment
  c := a >> 1   // Bitwise
  d := &c       // Miscellaneous
  fmt.Println(a, b, c, d)
}

command line
2 true 1 0xc000084000


Control Structures Branching & Looping
Control Structures
Hanya ada beberapa struktur kontrol di Go. untuk percabangan kita menggunakan if dan switch, untuk menulis loop kita menggunakan kata kunci for.

If Statement With Comparison
-code editor
var myAge = 17

if myAge == 5 {
  fmt.Println("Ypu too young")
}
if (myAge == 17){
  fmt.Println("So Sweet")
}
if myAge >= 17 && myAge <= 30 {
  fmt.Println("My Age is between 17 and 30")
}
if dadAge := 9; dadAge < myAge {
  fmt.Println(dadAge)
}


      |
      |
      v
Belah ketupat (condition) -------
      |                         |
 True |                   False |
      |                         |
      v                         |
Persegi (block if)              |
      |                         |
      |<------------------------|
      |
      v


If Else Statement
-code editor
number := 4
if number%2 == 0 {
  fmt.Println("Genap")
} else {
  fmt.Println("Ganjil")
}


      |
      |
      v
Belah ketupat (condition) -------
      |                         |
 True |                   False |
      |                         |
      v                         v
Persegi (block if)     Persegi(block else)
      |                         |
      |<------------------------|
      |
      v


If, Else If, Else
-code editor
hour := 15
if hour < 12 {
  fmt.Println("Selamat Pagi")
} else if hour < 18 {
  fmt.Println("Selamat Sore")
} else {
  fmt.Println("Selamat Malam")
}


      |
      |
      v                       False  
Belah ketupat (condition 1) ---------> Belah ketupat(condition 2) -------
      |                         |                                       |
 True |                    True |                                 False |
      |                         |                                       |
      v                         v                                       v
Persegi (block if)     Persegi (block else if)              Persegi (block else)
      |                         |                                       |
      |<------------------------|                                       |
      |<----------------------------------------------------------------|
      v


Nested If Statement
-code editor
var v1 int = 400
var v2 int = 700

if v1 == 400 {
  if v2 == 700 {
    fmt.Printlf("Value of v1 is 400 and v2 is 700\n");
  }
}


Switch <\>
switch expression {
  case value 1 :
    statement(s)
  case value 2 :
    statement(s)
  default :
    statement(s)
}

-code editor
var today int = 2
switch today {
  case 1 :
    fmt.Printf("Today is Monday")
  case 2 :
    fmt.Printf("Today is Tuesday")
  default :
    fmt.Printf("Invalid Date")
}

command line
Today is Tuesday


Switch Without Expression
switch {
  case expression == value 1 :
    statement(s)
  case expression == value 2 :
    statement(s)
  default :
    statement(s)
}

-code editor
var today int = 2
switch today {
  case today == 1 :
    fmt.Printf("Today is Monday")
  case today == 2 :
    fmt.Printf("Today is Tuesday")
  default :
    fmt.Printf("Invalid Date")
}

command line
Today is Tuesday


Fallthrough >_
Dalam GO, kontrol keluar dari pernyataan switch segera setelah sebuah case dieksekusi. Pernyataan fallthrough digunakan untuk mentransfer kontrol ke pernyataan pertama dari case yang ada segera setelah case yang telah dieksekusi.

-code editor
value := 42
switch value {
case 100 :
  fmt.Println(100)
  fallthrough
case 42 :
  fmt.Println(42)
  fallthrough
case 1 :
  fmt.Println(1)
  fallthrough
default :
  fmt.Println("default")
}

command line
42
1
default


Loops ><
Statement which allows code to be repeatelly executed.

for init; condition; post {
  //run command till condition is true 
}

-code editor
package main
import "fmt"
func main() {
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }
}

command line
0
1
2
3
4


While ><
For is Go's "while"

Pada saat itu Anda dapat menghilangkan titik koma: C sementara dieja dalam Go.

-code editor
package main

import "fmt"

func main() {
  sum := 1
  for sum < 10 {
    sum += sum
  }
  fmt.Println(sum)
}

command line
16


Continue & Break ><
-code editor
package main
import "fmt"
func main() {
  for i := 0; i < 5; i++ {
    if i == 1 {
      continue
    }
    if i > 3 {
      break
    }
    fmt.Println(i)
  }
}

command line
0
2
3


Loops Over String ><
command line
character H starts at byte position 0
character e starts at byte position 1
character l starts at byte position 2
character l starts at byte position 3
character o starts at byte position 4
-----
H
e
l
l
o

-code editor
package main

import (
  "fmt"
)

func main() {
  sentence := "Hello"

  for i := 0; i < len(sentence); i++ {
    fmt.Println(string(sentence[i]))
  }

  fmt.Println("-----")

  for pos, char := range sentence {
    fmt.Printf("character %c starts at byte position %d\n", char, pos)
  }
}


Advance Looping (1) ><
-code editor 
package main

import "fmt"

func main() {
  N := 5
  for i := 0; i < N; i++ {
    for j := 0; j < N; j++ {
      fmt.Print("*")
    }
    fmt.Println() 
  }
}

-command line
*****
*****
*****
*****
*****


Advance Looping (2) ><
-code editor 
package main

import "fmt"

func main() {
  N := 5
  for i := 0; i < N; i++ {
    for j := 0; j <= i; j++ {
      fmt.Print("*")
    }
    fmt.Println() 
  }
}

-command line
*
**
***
****
*****
