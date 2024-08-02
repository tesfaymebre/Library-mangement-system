package main

import (
	"testing"
)

func TestIsAlphaNum(t *testing.T) {
	tests := []struct {
		input rune
		want  bool
	}{
		{'a', true},
		{'A', true},
		{'1', true},
		{' ', false},
		{'#', false},
	}

	for _, test := range tests {
		got := IsAlphaNum(test.input)

		if got != test.want {
			t.Errorf("IsAlphaNum(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestCleanString(t *testing.T) {
	tests := []struct {
		input, want string
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
		got := CleanString(test.input)
		if got != test.want {
			t.Errorf("CleanString(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal, Panama!", true},
		{"racecar", true},
		{"hello", false},
		{"Madam In Eden, I'm Adam", true},
		{"12321", true},
		{"12345", false},
		{"a", true},
		{"ab", false},
	}

	for _, test := range tests {
		cleanedInput := CleanString(test.input)
		got := IsPalindrome(cleanedInput)
		if got != test.want {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}
