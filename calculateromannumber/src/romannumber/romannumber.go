package romannumber

func ConvertToRomanNumber(number int) string {
	valueRomanNumber := []int{100, 90, 50, 40, 30, 10, 9, 5, 4, 1}
	roman := make(map[int]string)
	roman[1] = "I"
	roman[4] = "IV"
	roman[5] = "V"
	roman[9] = "IX"
	roman[10] = "X"
	roman[30] = "XXX"
	roman[40] = "XL"
	roman[50] = "L"
	roman[90] = "XC"
	roman[100] = "C"
	romanNumber := ""

	for _, arabicNumber := range valueRomanNumber {
		for number >= arabicNumber {
			number -= arabicNumber
			romanNumber += roman[arabicNumber]
		}
	}

	return romanNumber
}
