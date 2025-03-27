package main

import (
	"fmt"

	"./OOPs"
)

func main() {
	fmt.Println("=== Struct and Methods ===")
	car := OOPs.Vehicle{Make: "Toyota", Model: "Camry", Year: 2022}
	car.Start()
	car.DisplayInfo()
	fmt.Println()
}
