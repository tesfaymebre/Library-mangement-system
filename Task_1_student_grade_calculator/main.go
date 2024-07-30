package main

import (
	"fmt"
)

func main() {
	name := ""
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	numSubjects := 0
	fmt.Print("Hom many subjects you took: ")
	fmt.Scan(&numSubjects)

	for numSubjects <= 0 {
		fmt.Println("You entered an invalid number.")
		fmt.Print("Please enter the number of subjects you took: ")
		fmt.Scan(&numSubjects)
	}

	subjects := make([]string, numSubjects)
	grades := make([]float64, numSubjects)
	total := 0.0

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Name of Subject %d: ", i+1)
		fmt.Scan(&subjects[i])

		fmt.Printf("Enter %s's grade: ", subjects[i])
		fmt.Scan(&grades[i])

		for grades[i] < 0 || grades[i] > 100 {
			fmt.Println("You entered invalid grade.")
			fmt.Printf("Enter %s's grade: ", subjects[i])
			fmt.Scan(&grades[i])
		}

		total += grades[i]
	}

	fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("Subject Grades:")
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("%s: %.2f\n", subjects[i], grades[i])
	}

	fmt.Printf("Average: %.2f", total/float64(numSubjects))

}
