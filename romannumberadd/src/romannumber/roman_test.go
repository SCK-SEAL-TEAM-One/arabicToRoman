package romannumber_test

import "testing"

func Test_RomanNumberAdd_Input_FirstNumber_I_SecondNumber_I_Should_Be_II(t *testing.T) {
	firstnumber := "I"
	secondnumber := "I"
	expectedromannumberadd := "II"

	actualroman := RomaNumberAdd(firstnumber, secondnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %s but got %s", expectedromannumberadd, actualroman)
	}

}
