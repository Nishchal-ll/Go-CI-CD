package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "with name", input: "Alice", expected: "Hello, Alice!"},
		{name: "empty name defaults to World", input: "", expected: "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Greet(tt.input)
			if result != tt.expected {
				t.Errorf("Greet(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	HealthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := strings.TrimSpace(rec.Body.String())
	expected := `{"status": "ok"}`
	if body != expected {
		t.Errorf("expected body %q, got %q", expected, body)
	}
}

func TestRootHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?name=Antigravity", nil)
	rec := httptest.NewRecorder()

	RootHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := strings.TrimSpace(rec.Body.String())
	expected := "Hello, Antigravity!"
	if body != expected {
		t.Errorf("expected body %q, got %q", expected, body)
	}
}
