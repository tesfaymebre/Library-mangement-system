package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ReadInput reads input from the user.
func ReadInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	return strings.TrimSpace(input), nil
}

// Cleaner is an interface for cleaning strings.
type Cleaner interface {
	Clean(string) string
}

// PunctuationCleaner removes punctuation and symbols from the string.
type PunctuationCleaner struct{}

func (p PunctuationCleaner) Clean(input string) string {
	input = strings.ToLower(input)
	var sb strings.Builder
	for _, r := range input {
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			sb.WriteRune(' ')
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// AlphaNumCleaner removes non-alphanumeric characters from the string.
type AlphaNumCleaner struct{}

func (a AlphaNumCleaner) Clean(input string) string {
	var sb strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(unicode.ToLower(r))
		}
	}
	return sb.String()
}

// ExtractWords splits a cleaned string into words.
func ExtractWords(cleanedString string) []string {
	return strings.Fields(cleanedString)
}

// CountWordFrequency counts the frequency of each word in a slice.
func CountWordFrequency(words []string) map[string]int {
	freqWords := make(map[string]int)
	for _, word := range words {
		freqWords[word]++
	}
	return freqWords
}

// PrintWordFrequencies prints the word frequencies in a readable format.
func PrintWordFrequencies(freqWords map[string]int) {
	fmt.Println("Word Frequencies:")
	for word, count := range freqWords {
		fmt.Printf("%s: %d\n", word, count)
	}
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

func countWordFrequencyOption() {
	sentence, err := ReadInput("Enter a sentence: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if sentence == "" {
		fmt.Println("Error: empty input. Try again.")
		return
	}

	cleaner := PunctuationCleaner{}
	cleanedString := cleaner.Clean(sentence)
	words := ExtractWords(cleanedString)

	if len(words) == 0 {
		fmt.Println("Error: the input doesn't contain valid words.")
		return
	}

	freqWords := CountWordFrequency(words)
	PrintWordFrequencies(freqWords)
}

func checkPalindromeOption() {
	input, err := ReadInput("Enter a string: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cleaner := AlphaNumCleaner{}
	cleanedInput := cleaner.Clean(input)
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

func exitOption() {
	fmt.Println("Exiting the program ...")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n***************************************")
		fmt.Println("Choose an option:")
		fmt.Println("1. Count word frequency in a sentence")
		fmt.Println("2. Check if a string is a palindrome")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice (1/2/3): ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			countWordFrequencyOption()
		case "2":
			checkPalindromeOption()
		case "3":
			exitOption()
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
