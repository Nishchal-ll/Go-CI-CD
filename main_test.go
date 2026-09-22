package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "with custom name", input: "Nishchal", expected: "Hello, Nishchal!"},
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

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := strings.TrimSpace(rec.Body.String())
	expected := "Hello, Nishchal!"
	if body != expected {
		t.Errorf("expected body %q, got %q", expected, body)
	}
}
