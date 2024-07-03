package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman   string
		want    int
	}{
		{"I", 1},
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"X", 10},
		{"XV", 15},
		{"XXX", 30},
		{"L", 50},
		{"LXXX", 80},
		{"XC", 90},
		{"C", 100},
	}

	for _, tt := range tests {
		t.Run(tt.roman, func(t *testing.T) {
			got := romanToInt(tt.roman)
			if got != tt.want {
				t.Errorf("want %v , but got %v", tt.want, got)
			}
		})
	}
}
