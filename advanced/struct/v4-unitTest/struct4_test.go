package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"add 1+1", 1, 1, 2},
		{"add 2+3", 2, 3, 5},
		{"add -1+1", -1, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Add(tt.a, tt.b)
			if actual != tt.expected {
				t.Errorf("Add(%d, %d): expected %d, actual %d", tt.a, tt.b, tt.expected, actual)
			}
		})
	}
}
