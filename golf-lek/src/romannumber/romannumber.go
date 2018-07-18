package romannumber

func ConvertToRomanSymbol(number int) string {
	numberBudget := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanCharacter := map[int]string{100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}

	var roman string
	for _, num := range numberBudget {
		for number >= num {
			number -= num
			roman += romanCharacter[num]
		}
	}
	return roman
}
