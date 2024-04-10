package main

import "testing"

func Test_struct4(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"aaaa"},
		{"bbbb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			struct4()
		})
	}
}
