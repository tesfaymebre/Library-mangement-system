package main

import (
	"testing"
)

// TestPunctuationCleaner tests the Clean method of PunctuationCleaner.
func TestPunctuationCleaner(t *testing.T) {
	cleaner := PunctuationCleaner{}
	input := "Hello, world! test@gmail.com"
	expected := "hello  world  test gmail com"
	result := cleaner.Clean(input)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestAlphaNumCleaner tests the Clean method of AlphaNumCleaner.
func TestAlphaNumCleaner(t *testing.T) {
	cleaner := AlphaNumCleaner{}
	tests := []struct {
		input, expected string
	}{
		{"A man, a plan, a canal, Panama!", "amanaplanacanalpanama"},
		{"racecar", "racecar"},
		{"hello", "hello"},
		{"Madam In Eden, I'm Adam", "madaminedenimadam"},
		{"12321", "12321"},
		{"12345", "12345"},
		{"", ""},
		{"a", "a"},
		{"ab", "ab"},
	}

	for _, test := range tests {
		got := cleaner.Clean(test.input)
		if got != test.expected {
			t.Errorf("CleanString(%q) = %v; want %v", test.input, got, test.expected)
		}
	}
}

// TestExtractWords tests the ExtractWords function.
func TestExtractWords(t *testing.T) {
	input := "hello world this is a test"
	expected := []string{"hello", "world", "this", "is", "a", "test"}
	result := ExtractWords(input)
	if len(result) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], result[i])
		}
	}
}

// TestCountWordFrequency tests the CountWordFrequency function.
func TestCountWordFrequency(t *testing.T) {
	words := []string{"hello", "world", "hello", "test"}
	expected := map[string]int{"hello": 2, "world": 1, "test": 1}
	result := CountWordFrequency(words)
	for word, count := range expected {
		if result[word] != count {
			t.Errorf("expected %s: %d, got %d", word, count, result[word])
		}
	}
}

// TestIsPalindrome tests the IsPalindrome function.
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal, Panama!", true},
		{"racecar", true},
		{"hello", false},
		{"Madam In Eden, I'm Adam", true},
		{"12321", true},
		{"12345", false},
	}

	cleaner := AlphaNumCleaner{}
	for _, test := range tests {
		cleanedInput := cleaner.Clean(test.input)
		result := IsPalindrome(cleanedInput)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
