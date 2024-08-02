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

// IsPunctuation checks if a rune is a punctuation or symbol.
func IsPunctuation(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSymbol(r)
}

// CleanString removes punctuation and symbols from the string.
func CleanString(input string) string {
	input = strings.ToLower(input)
	var sb strings.Builder
	for _, r := range input {
		if IsPunctuation(r) {
			sb.WriteRune(' ')
		} else {
			sb.WriteRune(r)
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

func main() {
	sentence, err := ReadInput("Enter a sentence: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if sentence == "" {
		fmt.Println("Error: empty input. try again.")
		return
	}

	cleanedString := CleanString(sentence)
	words := ExtractWords(cleanedString)

	if len(words) == 0 {
		fmt.Println("Error: the input doesn't contain valid words.")
		return
	}

	freqWords := CountWordFrequency(words)
	PrintWordFrequencies(freqWords)
}
