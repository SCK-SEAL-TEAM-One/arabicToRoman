package romannumber_test

import (
	. "romannumber"
	"testing"
)

func Test_ConvertToRomanSymbol_Input_1_Should_Be_I(t *testing.T) {
	number := 1
	expectedRoman := "I"

	actual := ConvertToRomanSymbol(number)

	if expectedRoman != actual {
		t.Errorf("Expected is %s but got %s", expectedRoman, actual)
	}
}
func Test_ConvertToRomanSymbol_Input_2_Should_Be_II(t *testing.T) {
	number := 2
	expectedRoman := "II"

	actual := ConvertToRomanSymbol(number)

	if expectedRoman != actual {
		t.Errorf("Expected is %s but got %s", expectedRoman, actual)
	}
}
func Test_ConvertToRomanSymbol_Input_4_Should_Be_IV(t *testing.T) {
	number := 4
	expectedRoman := "IV"

	actual := ConvertToRomanSymbol(number)

	if expectedRoman != actual {
		t.Errorf("Expected is %s but got %s", expectedRoman, actual)
	}
}
func Test_ConvertToRomanSymbol_Input_49_Should_Be_XLIX(t *testing.T) {
	number := 49
	expectedRoman := "XLIX"

	actual := ConvertToRomanSymbol(number)

	if expectedRoman != actual {
		t.Errorf("Expected is %s but got %s", expectedRoman, actual)
	}
}
