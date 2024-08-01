package main

import (
	"fmt"
)

type Student struct {
	name     string
	subjects []string
	grades   []float64
}

func (student *Student) CalculateAverage() float64 {
	total := 0.0

	for _, grade := range student.grades {
		total += grade
	}

	return total / float64(len(student.grades))

}

func main() {
	student := &Student{}

	fmt.Print("Enter your name: ")
	fmt.Scan(&student.name)

	numSubjects := 0
	fmt.Print("Hom many subjects you took: ")
	fmt.Scan(&numSubjects)

	for numSubjects <= 0 {
		fmt.Println("You entered an invalid number.")
		fmt.Print("Please enter the number of subjects you took: ")
		fmt.Scan(&numSubjects)
	}

	student.subjects = make([]string, numSubjects)
	student.grades = make([]float64, numSubjects)

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Name of Subject %d: ", i+1)
		fmt.Scan(&student.subjects[i])

		fmt.Printf("Enter %s's grade: ", student.subjects[i])
		fmt.Scan(&student.grades[i])

		for student.grades[i] < 0 || student.grades[i] > 100 {
			fmt.Println("You entered invalid grade.")
			fmt.Printf("Enter %s's grade: ", student.subjects[i])
			fmt.Scan(&student.grades[i])
		}
	}

	fmt.Printf("\nStudent Name: %s\n", student.name)
	fmt.Println("Subject Grades:")
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("%s: %.1f\n", student.subjects[i], student.grades[i])
	}

	fmt.Printf("Average: %.1f", student.CalculateAverage())

}
