package main

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	nameEncode string
	score int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode strings.Builder

	// your code here
	for _, char := range s.name {
		if char >= 'a' && + char <= 'z' {
			encodedChar := 'a' + ((char-'a'+3)%26)		// Pergeseran sebesar 3 karakter
			nameEncode.WriteRune(encodedChar)
		} else if char >= 'A' && char <= 'Z' {
			encodedChar := 'A' + ((char-'A'+3)%26)		// Pergeseran sebesar 3 karakter
			nameEncode.WriteRune(encodedChar)
		} else {
			nameEncode.WriteRune(char)
		}
	} 
	
	return nameEncode.String()
}

func (s *student) Decode() string {
	var nameDecode strings.Builder

  // your code here
  for _, char := range s.nameEncode {
	if char >= 'a' && char <= 'z' {
		decodedChar := 'a' + ((char-'a'+23)%26)		// Pergeseran 3 karakter
		nameDecode.WriteRune(decodedChar)
	} else if char >= 'A' && char <= 'Z'{
		decodedChar := 'A' + ((char-'A'+23)%26)		// Pergeseran 3 karakter
		nameDecode.WriteRune(decodedChar)
	} else {
		nameDecode.WriteRune(char)
	}
  }

  return nameDecode.String()
}

func main() {
	var menu int
  var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)
	
	if menu == 1 {
		fmt.Print("\nInput Student Name: ")		
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + "is : " + c.Decode())
	}
}