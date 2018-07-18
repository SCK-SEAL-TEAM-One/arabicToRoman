package romannumber

import (
	"testing"
)

func Test_ConvertToRomanNumber_Input_Number_1_Should_Be_I(t *testing.T) {
	number := 1
	expectedRomanNumber := "I"

	actualRomanNumber := ConvertToRomanNumber(number)
	if expectedRomanNumber != actualRomanNumber {
		t.Errorf("expected %s but got %s", expectedRomanNumber, actualRomanNumber)
	}
}

func Test_ConvertToRomanNumber_Input_Number_2_Should_Be_II(t *testing.T) {
	number := 2
	expectedRomanNumber := "II"

	actualRomanNumber := ConvertToRomanNumber(number)
	if expectedRomanNumber != actualRomanNumber {
		t.Errorf("expected %s but got %s", expectedRomanNumber, actualRomanNumber)
	}
}
func Test_ConvertToRomanNumber_Input_Number_4_Should_Be_IV(t *testing.T) {
	number := 4
	expectedRomanNumber := "IV"

	actualRomanNumber := ConvertToRomanNumber(number)
	if expectedRomanNumber != actualRomanNumber {
		t.Errorf("expected %s but got %s", expectedRomanNumber, actualRomanNumber)
	}
}
func Test_ConvertToRomanNumber_Input_Number_5_Should_Be_V(t *testing.T) {
	number := 5
	expectedRomanNumber := "V"

	actualRomanNumber := ConvertToRomanNumber(number)
	if expectedRomanNumber != actualRomanNumber {
		t.Errorf("expected %s but got %s", expectedRomanNumber, actualRomanNumber)
	}
}

func Test_ConvertToRomanNumber_Input_Number_6_Should_Be_VI(t *testing.T) {
	number := 6
	expectedRomanNumber := "VI"

	actualRomanNumber := ConvertToRomanNumber(number)
	if expectedRomanNumber != actualRomanNumber {
		t.Errorf("expected %s but got %s", expectedRomanNumber, actualRomanNumber)
	}
}

func Test_ConvertToRomanNumberr_Input_Number_9_Should_Be_IX(t *testing.T) {
	number := 9
	expectedRomanNumber := "IX"

	actualRomanNumber := ConvertToRomanNumber(number)
	if expectedRomanNumber != actualRomanNumber {
		t.Errorf("expected %s but got %s", expectedRomanNumber, actualRomanNumber)
	}
}

func Test_ConvertToRomanNumber_Input_Number_10_Should_Be_X(t *testing.T) {
	number := 10
	expectedRomanNumber := "X"

	actualRomanNumber := ConvertToRomanNumber(number)
	if expectedRomanNumber != actualRomanNumber {
		t.Errorf("expected %s but got %s", expectedRomanNumber, actualRomanNumber)
	}
}
