package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ReadInput reads a line of input from the user
func ReadInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	return strings.TrimSpace(input), nil
}

// IsAlphaNum checks if a rune is alphanumeric
func IsAlphaNum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// CleanString removes non-alphanumeric characters and converts to lowercase
func CleanString(input string) string {
	var sb strings.Builder
	for _, r := range input {
		if IsAlphaNum(r) {
			sb.WriteRune(unicode.ToLower(r))
		}
	}
	return sb.String()
}

// IsPalindrome checks if a given string is a palindrome
func IsPalindrome(word string) bool {
	size := len(word)
	for i := 0; i < size/2; i++ {
		if word[i] != word[size-i-1] {
			return false
		}
	}
	return true
}

func main() {
	input, err := ReadInput("Enter string: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cleanedInput := CleanString(input)
	if len(cleanedInput) == 0 {
		fmt.Println("Inserted invalid input. Try again.")
		return
	}

	if IsPalindrome(cleanedInput) {
		fmt.Printf("The string '%s' is a palindrome\n", input)
	} else {
		fmt.Printf("The string '%s' is not a palindrome\n", input)
	}
}
