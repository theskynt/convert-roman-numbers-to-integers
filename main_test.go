package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	t.Run("should return 1 when input is I", func(t *testing.T) {
		// Arrange
		input := "I"
		want := 1

		// Act
		got := romanToInt(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 3 when input is III", func(t *testing.T) {
		// Arrange
		input := "III"
		want := 3

		// Act
		got := romanToInt(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 4 when input is IV", func(t *testing.T) {
		// Arrange
		input := "IV"
		want := 4

		// Act
		got := romanToInt(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 5 when input is V", func(t *testing.T) {
		// Arrange
		input := "V"
		want := 5

		// Act
		got := romanToInt(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 9 when input is IX", func(t *testing.T) {
		// Arrange
		input := "IX"
		want := 9

		// Act
		got := romanToInt(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 10 when input is X", func(t *testing.T) {
		// Arrange
		input := "X"
		want := 10

		// Act
		got := romanToInt(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

}