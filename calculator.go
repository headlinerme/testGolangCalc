package main

import (
	"fmt"
	"slices"
	"strconv"
)

func eval(num1, num2, operation string) string {

	var calculation = []string{num1, num2, operation}

	arabicNumerals := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	romanNumerals := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20}

	result := ""

	if slices.Contains(arabicNumerals, calculation[0]) && slices.Contains(arabicNumerals, calculation[1]) {

		var convertedNums []int

		for i, v := range calculation {

			if i != 2 {
				v, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				convertedNums = append(convertedNums, v)
				continue

			} else {
				if i == 2 && v == "+" {
					result = strconv.Itoa(convertedNums[0] + convertedNums[1])

				} else if i == 2 && v == "-" {
					result = strconv.Itoa(convertedNums[0] - convertedNums[1])

				} else if i == 2 && v == "*" {
					result = strconv.Itoa(convertedNums[0] * convertedNums[1])

				} else if i == 2 && v == "/" {
					result = strconv.Itoa(convertedNums[0] / convertedNums[1])

				} else {
					return "Введите корректную операцию: (+, -, *, /)"
				}
			}
		}

	} else if romanNumerals[calculation[0]] > 0 && romanNumerals[calculation[1]] > 0 {
		romanResult := 0

		if calculation[2] == "+" {
			romanResult += romanNumerals[calculation[0]] + romanNumerals[calculation[1]]

		} else if calculation[2] == "-" {
			romanResult += romanNumerals[calculation[0]] - romanNumerals[calculation[1]]

		} else if calculation[2] == "*" {
			romanResult += romanNumerals[calculation[0]] * romanNumerals[calculation[1]]

		} else if calculation[2] == "/" {
			romanResult += romanNumerals[calculation[0]] / romanNumerals[calculation[1]]

		} else {
			return "Ошибка"
		}

		if romanResult > 0 {
			for k, v := range romanNumerals {
				if romanResult == v {
					result = k
				}
			}
		} else {
			return "Ошибка"
		}

	} else {
		return "Ошибка"
	}

	return result
}

func main() {
	var num1, operation, num2 string
	fmt.Scanln(&num1, &operation, &num2)

	fmt.Println(eval(num1, num2, operation))
}
