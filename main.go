package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
	Major string
	GPA   float64
}

func displayStudent(student Student) {
	fmt.Println("ID:", student.ID)
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Major:", student.Major)
	fmt.Println("GPA:", student.GPA)
}

func isPassed(gpa float64) bool {
	if gpa >= 2.0 {
		return true
	}

	return false
}

func addStudent(students []Student, student Student) []Student {
	students = append(students, student)
	return students
}

func findStudent(students []Student, id int) (Student, error) {
	for i := 0; i < len(students); i++ {
		if students[i].ID == id {
			return students[i], nil
		}
	}

	return Student{}, fmt.Errorf("student not found")
}

func removeStudent(students []Student, id int) ([]Student, error) {
	for i := 0; i < len(students); i++ {
		if students[i].ID == id {
			students = append(students[:i], students[i+1:]...)
			return students, nil
		}
	}

	return students, fmt.Errorf("student not found")
}

func main() {

	students := []Student{
		{
			ID:    1,
			Name:  "Tasneem",
			Age:   21,
			Major: "Computer Science",
			GPA:   3.7,
		},
		{
			ID:    2,
			Name:  "Ahmad",
			Age:   22,
			Major: "Computer Science",
			GPA:   3.20,
		},
	}

	newStudent := Student{
		ID:    3,
		Name:  "Sara",
		Age:   20,
		Major: "Computer Science",
		GPA:   3.50,
	}

	students = addStudent(students, newStudent)

	student, err := findStudent(students, 3)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	displayStudent(student)

	if isPassed(student.GPA) {
		fmt.Println("Status: Passed")
	} else {
		fmt.Println("Status: Failed")
	}

	students, err = removeStudent(students, 3)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Student removed successfully")
}



