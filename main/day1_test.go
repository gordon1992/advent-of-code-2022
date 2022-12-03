package main

import (
	"testing"
)

func TestWithTestInputPart1(t *testing.T) {
	var expected = 24000
	var actual = highestCaloriesCarriedByElf("../resources/day-1-test-input")
	if actual != expected {
		t.Errorf("got %d, wanted %d", actual, expected)
	}
}

func TestWithPuzzleInputPart1(t *testing.T) {
	var want = 68923
	var calories = highestCaloriesCarriedByElf("../resources/day-1-input")
	if calories != want {
		t.Errorf("got %d, wanted %d", calories, want)
	}
}

func TestWithTestInputPart2(t *testing.T) {
	var expected = 45000
	var actual = caloriesCarriedByTopThreeElvesWithMostSnacks("../resources/day-1-test-input")
	if actual != expected {
		t.Errorf("got %d, wanted %d", actual, expected)
	}
}

func TestWithPuzzleInputPart2(t *testing.T) {
	var want = 200044
	var calories = caloriesCarriedByTopThreeElvesWithMostSnacks("../resources/day-1-input")
	if calories != want {
		t.Errorf("got %d, wanted %d", calories, want)
	}
}
