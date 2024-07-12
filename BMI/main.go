package main

import (
	"fmt"
)

var weight int
var height int

func mohasebe(weight, height int) float64 {
	heightInMeters := float64(height) / 100
	bmi := float64(weight) / (heightInMeters * heightInMeters)
	return bmi
}

func main() {
	fmt.Println("Enter your weight (kg):")
	fmt.Scanf("%d", &weight)

	fmt.Println("Enter your height (cm):")
	fmt.Scanf("%d", &height)

	bmi := mohasebe(weight, height)

	if bmi < 18.5 {
		fmt.Println("Loss")
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Println("Good")
	} else if bmi >= 25 {
		fmt.Println("Over")
	}
}
