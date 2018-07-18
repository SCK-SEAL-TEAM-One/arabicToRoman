package romannumber

func ConvertToRoman(number int) string {
	numberBudget := []int{50, 40, 10, 9, 5, 4}
	romanCharacter := map[int]string{50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}

	var romanNumber string
	for _, value := range numberBudget {
		if number >= value {
			number = number - value
			romanNumber += romanCharacter[value]
		}
	}
	for index := number; index > 0; index-- {
		romanNumber += romanCharacter[1]
	}
	return romanNumber
}
