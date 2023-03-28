package FizzBuzz

import (
	"testing"
)

func TestFizzBuzzReturnsNumberForNumberNotDivisibleByThreeOrFive(t *testing.T) {
	scenarios := []struct {
		number   int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{8, "8"},
	}

	for _, example := range scenarios {
		response := FizzBuzz(example.number)

		if response != example.expected {
			t.Errorf("FizzBuzz expected: %q, got: %q", example.expected, response)
		}
	}

}

func TestFizzBuzzReturnsFizzForNumberDivisibleByThree(t *testing.T) {
	scenarios := []struct {
		number   int
		expected string
	}{
		{3, "Fizz"},
		{6, "Fizz"},
	}

	for _, example := range scenarios {
		response := FizzBuzz(example.number)

		if response != example.expected {
			t.Errorf("FizzBuzz expected: %q, got: %q", example.expected, response)
		}
	}

}

func TestFizzBuzzReturnsBuzzForNumberDivisibleByFive(t *testing.T) {
	scenarios := []struct {
		number   int
		expected string
	}{
		{5, "Buzz"},
		{10, "Buzz"},
	}

	for _, example := range scenarios {
		response := FizzBuzz(example.number)

		if response != example.expected {
			t.Errorf("FizzBuzz expected: %q, got: %q", example.expected, response)
		}
	}

}

func TestFizzBuzzReturnsFizzBuzzForNumberDivisibleByFifteen(t *testing.T) {
	scenarios := []struct {
		number   int
		expected string
	}{
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
	}

	for _, example := range scenarios {
		response := FizzBuzz(example.number)

		if response != example.expected {
			t.Errorf("FizzBuzz expected: %q, got: %q", example.expected, response)
		}
	}

}
