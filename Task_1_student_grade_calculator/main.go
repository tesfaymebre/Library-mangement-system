package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Student struct to hold student details and grades
type Student struct {
	name   string
	grades []Grade
}

// Grade struct to hold subject name and grade point
type Grade struct {
	subjectName string
	gradePoint  float64
}

// containsSpecialCharacterOrNumber checks if the string contains any non-alphabetical characters
func containsSpecialCharacterOrNumber(s string) bool {
	match, _ := regexp.MatchString(`[^a-z A-Z]`, s)
	return match
}

// Method to calculate the average grade of a student
func (student Student) calculateAverageGrade() (float64, error) {
	totalGrade := 0.0

	for _, grade := range student.grades {
		totalGrade += grade.gradePoint
	}

	// Avoid division by zero if there are no grades
	if len(student.grades) == 0 {
		return 0.0, errors.New("No grades found")
	}

	return totalGrade / float64(len(student.grades)), nil
}

func newStudent(reader *bufio.Reader) *Student {
	// initialize a new student
	var student Student

	// Get the student's name
	fmt.Print("Enter your name: ")

	for {
		name, err := reader.ReadString('\n')
		name = strings.TrimSpace(name) // Remove any trailing newline characters

		if err != nil || name == "" || containsSpecialCharacterOrNumber(name) {
			fmt.Print("Invalid Name. Please Input a Valid Name.")
		} else {
			student.name = name
			break
		}
	}

	return &student
}

func populateGrade(student *Student, reader *bufio.Reader) {
	// Get the number of subjects
	var subjectCount int
	fmt.Print("Enter the number of subjects: ")
	for {
		numberInput, err := reader.ReadString('\n')
		numberInput = strings.TrimSpace(numberInput)
		subjectCount, err = strconv.Atoi(numberInput)

		if err != nil || subjectCount <= 0 {
			fmt.Println("Invalid input. Please enter a valid number.")
		} else {
			break
		}
	}

	// Loop through the subjects to get their details
	for i := 0; i < subjectCount; i++ {
		var grade Grade

		// Get the subject name
		fmt.Print("Enter the subject name: ")
		for {
			subjectName, err := reader.ReadString('\n')
			subjectName = strings.TrimSpace(subjectName)

			if err != nil || subjectName == "" || containsSpecialCharacterOrNumber(subjectName) {
				fmt.Print("Please enter valid subject name.")
			} else {
				grade.subjectName = subjectName
				break
			}
		}

		// Get the grade point
		fmt.Print("Enter the grade: ")
		for {
			gradeInput, err := reader.ReadString('\n')
			gradeInput = strings.TrimSpace(gradeInput)
			grade.gradePoint, err = strconv.ParseFloat(gradeInput, 64)

			if err != nil || grade.gradePoint < 0 || grade.gradePoint > 100 {
				fmt.Println("Invalid input. Please enter a number between 0 and 100.")
			} else {
				break
			}
		}

		student.grades = append(student.grades, grade)
	}

}

func displayStudentInfo(student *Student) {
	// Print the student details and average grade
	fmt.Println("-----------------------------------")
	fmt.Printf("Student Name: %s\n", student.name)
	fmt.Println("Grades:")
	for _, grade := range student.grades {
		fmt.Printf("Subject: %s, Grade: %.2f\n", grade.subjectName, grade.gradePoint)
	}

	averageGrade, err := student.calculateAverageGrade()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("The average grade for %s is %.2f\n", student.name, averageGrade)
	fmt.Println("-----------------------------------")

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Create a new student and no need for error checking as the newStudent function will handle it
	student := newStudent(reader)

	//populate the student's grades and no need for error checking as the populateGrade function will handle it
	populateGrade(student, reader)

	// Display the student's information
	// No need for error checking as the displayStudentInfo function will handle it
	displayStudentInfo(student)

}
