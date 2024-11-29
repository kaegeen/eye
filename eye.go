package main

import (
	"fmt"
)

func main() {
	// Initial age and prescription
	startAge := 27
	endAge := 80
	leftEyePrescription := -8.0  // Left eye prescription in diopters
	rightEyePrescription := -7.0 // Right eye prescription in diopters
	deteriorationRate := 0.1     // Diopter deterioration rate per year

	// Loop from age 27 to 80
	for age := startAge; age <= endAge; age++ {
		// Calculate the visual prescription for both eyes
		leftEyeVisualAbility := leftEyePrescription + (float64(age)-float64(startAge))*deteriorationRate
		rightEyeVisualAbility := rightEyePrescription + (float64(age)-float64(startAge))*deteriorationRate

		// Print the visual ability at each age
		fmt.Printf("Age: %d, Left Eye Prescription: %.2f, Right Eye Prescription: %.2f\n",
			age, leftEyeVisualAbility, rightEyeVisualAbility)
	}
}
