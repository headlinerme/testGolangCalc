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
		"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
		"XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30,
		"XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40,
		"XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
		"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60,
		"LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70,
		"LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80,
		"LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
		"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100}

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
					panic("Введите корректную операцию: (+, -, *, /)")
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
			panic("Ошибка")
		}

		if romanResult > 0 {
			for k, v := range romanNumerals {
				if romanResult == v {
					result = k
				}
			}
		} else {
			panic("Ошибка")
		}

	} else {
		panic("Ошибка")
	}

	return result
}

func main() {
	var num1, operation, num2, other string
	fmt.Scanln(&num1, &operation, &num2, &other)

	if len(other) == 1 {
		panic("Ошибка")
	} else {
		fmt.Println(eval(num1, num2, operation))
	}
}
