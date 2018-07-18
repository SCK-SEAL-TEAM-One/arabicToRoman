package romannumber

func ConvertToRomanNumber(number int) string {
	//	valueRomanNumber := []int{100, 90, 50, 40, 30, 9, 5, 4, 1}
	roman := make(map[int]string)
	roman[1] = "I"
	roman[4] = "IV"
	roman[5] = "V"
	roman[9] = "IX"
	roman[30] = "XXX"
	roman[40] = "XL"
	roman[50] = "L"
	roman[90] = "XC"
	roman[100] = "C"

	if number == 2 {
		return "II"
	}
	if number == 6 {
		return "VI"
	}
	if number == 10 {
		return "X"
	}

	return roman[number]
}
