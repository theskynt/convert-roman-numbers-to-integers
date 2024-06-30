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

	t.Run("should return 8 when input is VIII", func(t *testing.T) {
		// Arrange
		input := "VIII"
		want := 8

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

	t.Run("should return 15 when input is XV", func(t *testing.T) {
		// Arrange
		input := "XV"
		want := 15

		// Act
		got := romanToInt(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 18 when input is XVIII", func(t *testing.T) {
		// Arrange
		input := "XVIII"
		want := 18

		// Act
		got := romanToInt(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

	t.Run("should return 19 when input is XIX", func(t *testing.T) {
		// Arrange
		input := "XIX"
		want := 19

		// Act
		got := romanToInt(input)

		// Assert
		if got != want {
			t.Errorf("want %v , but got %v", want, got)
		}
	})

}
