package main

import "testing"

func TestGetMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "with name", input: "Nishchal", expected: "Hello, Nishchal!"},
		{name: "empty name defaults to World", input: "", expected: "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetMessage(tt.input)
			if result != tt.expected {
				t.Errorf("GetMessage(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
