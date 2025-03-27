package OOPs

import "fmt"

// Vehicle represents a basic vehicle with common attributes
type Vehicle struct {
	Make  string
	Model string
	Year  int
}

// Start simulates starting the vehicle
func (v Vehicle) Start() {
	fmt.Printf("Starting the %s %s\n", v.Make, v.Model)
}

// DisplayInfo shows vehicle information
func (v Vehicle) DisplayInfo() {
	fmt.Printf("Make: %s, Model: %s, Year: %d\n", v.Make, v.Model, v.Year)
}
