package main

import "fmt"

func main() {
	var name, major, email string = "Tasneem", "cs", "tasneem@gmail.com"
	var id, yearOfBirth int = 12323555, 2005
	var GPA float64 = 3.70

	fmt.Println("\nStudent Information \n---------------------")
	fmt.Println("Name:", name)
	fmt.Println("Major:", major)
	fmt.Println("ID:", id)
	fmt.Println("Year of Birth:", yearOfBirth)
	fmt.Println("GPA:", GPA)
	fmt.Println("Email:", email, "\n")
}
