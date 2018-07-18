package romannumber_test

import (
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
