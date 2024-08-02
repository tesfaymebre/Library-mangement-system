package main

import (
	"testing"
)

// TestIsPunctuation tests the IsPunctuation function.
func TestIsPunctuation(t *testing.T) {
	tests := []struct {
		input rune
		want  bool
	}{
		{'!', true},
		{'a', false},
		{'1', false},
		{',', true},
		{'-', true},
	}

	for _, test := range tests {
		got := IsPunctuation(test.input)
		if got != test.want {
			t.Errorf("IsPunctuation(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

// TestCleanString tests the CleanString function.
func TestCleanString(t *testing.T) {
	input := "Hello, world! test@gmail.com"
	want := "hello  world  test gmail com"
	got := CleanString(input)
	if got != want {
		t.Errorf("CleanString(%q) = %q; want %q", input, got, want)
	}
}

// TestExtractWords tests the ExtractWords function.
func TestExtractWords(t *testing.T) {
	input := "hello world this is a test"
	want := []string{"hello", "world", "this", "is", "a", "test"}
	got := ExtractWords(input)
	if !equalSlices(got, want) {
		t.Errorf("ExtractWords(%q) = %v; want %v", input, got, want)
	}
}

// TestCountWordFrequency tests the CountWordFrequency function.
func TestCountWordFrequency(t *testing.T) {
	input := []string{"hello", "world", "hello"}
	want := map[string]int{"hello": 2, "world": 1}
	got := CountWordFrequency(input)
	if !equalMaps(got, want) {
		t.Errorf("CountWordFrequency(%v) = %v; want %v", input, got, want)
	}
}

// equalSlices checks if two string slices are equal.
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// equalMaps checks if two string-to-int maps are equal.
func equalMaps(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
