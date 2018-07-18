package romannumber_test

import (
	. "romannumber"
	"testing"
)

func Test_ConvertToRoman_Input_1_Should_Be_I(t *testing.T) {
	number := 1
	expected := "I"

	actual := ConvertToRoman(number)
	if expected != actual {
		t.Errorf("expect %s but got %s", expected, actual)
	}
}
