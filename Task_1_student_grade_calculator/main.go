package main

import (
	"fmt"
)

type Student struct {
	name   string
	result []Grade
}

type Grade struct {
	subject string
	score   float64
}

func (s *Student) CalculateAverage() float64 {
	total := 0.0

	for _, res := range s.result {
		total += res.score
	}

	return total / float64(len(s.result))

}

func main() {
	s := &Student{}

	fmt.Print("Enter your name: ")
	fmt.Scan(&s.name)

	numSubjects := 0
	fmt.Print("Hom many subjects you took: ")
	fmt.Scan(&numSubjects)

	for numSubjects <= 0 {
		fmt.Println("You entered an invalid number.")
		fmt.Print("Please enter the number of subjects you took: ")
		fmt.Scan(&numSubjects)
	}

	s.result = make([]Grade, numSubjects)

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Name of Subject %d: ", i+1)
		fmt.Scan(&s.result[i].subject)

		fmt.Printf("Enter %s's grade: ", s.result[i].subject)
		fmt.Scan(&s.result[i].score)

		for s.result[i].score < 0 || s.result[i].score > 100 {
			fmt.Println("You entered invalid grade.")
			fmt.Printf("Enter %s's grade: ", s.result[i].subject)
			fmt.Scan(&s.result[i].score)
		}
	}

	fmt.Printf("\nStudent Name: %s\n", s.name)
	fmt.Println("Subject Grades:")
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("%s: %.1f\n", s.result[i].subject, s.result[i].score)
	}

	fmt.Printf("Average: %.1f", s.CalculateAverage())

}
