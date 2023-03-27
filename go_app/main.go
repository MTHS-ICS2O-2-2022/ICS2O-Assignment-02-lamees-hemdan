// Created By: Lamees Hemdan
// Created On: March 2023
//
// This file contains the main function for the go_app application.

package main

import "fmt"

func main() {
	var base1, base2, height float64
	fmt.Println("Enter the base1 of the trapezoid: ")
	fmt.Scan(&base1)
	fmt.Println("Enter the base2 of the trapezoid: ")
	fmt.Scan(&base2)
	fmt.Println("Enter the height of the trapezoid: ")
	fmt.Scan(&height)
	area := (base1 + base2) * height / 2
	fmt.Println("The area of the trapezoid is: ", area)

	fmt.Println("\nDone.")
}
