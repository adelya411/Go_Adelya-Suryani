package main

import (
	"fmt"
)

// Definisikan struct Car
type Car struct {
	CarType string
	FuelIn  float64
}

func main() {
	// Membuat instance Car
	myCar := Car{
		CarType: "SUV",
	}

	// Pengguna memasukkan nilai FuelIn
	fmt.Print("Jumlah bahan bakar (dalam liter): ") // 10.5
	fmt.Scan(&myCar.FuelIn)

	// Perkiraan jarak
	fuelEfficiency := 1.5 // L/mill
	estimatedDistance := myCar.FuelIn * fuelEfficiency

	// Hasil
	fmt.Printf("Car type: %s, est. distance: %.2f\n", myCar.CarType, estimatedDistance)
}
