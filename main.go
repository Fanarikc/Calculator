package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите два числа (от 1 до 10 или I до X) и операцию (например, 5 + 3 или X + V): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разбиваем ввод на числа и операцию
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Некорректный ввод")
		return
	}

	// Проверяем, являются ли числа римскими числами или арабскими числами
	isRimNum1 := isRim(parts[0])
	isRimNum2 := isRim(parts[2])

	// Проверяем, если одно число римское, а другое - арабское, выводим ошибку
	if isRimNum1 && !isRimNum2 || !isRimNum1 && isRimNum2 {
		fmt.Println("Ошибка: нельзя выполнять операции смешиванием римских и арабских чисел")
		return
	}

	// Преобразуем числа в int или в римские числа
	var num1, num2 interface{}
	var err error

	if isRimNum1 {
		num1, err = rimToArab(parts[0])
		if err != nil {
			fmt.Println("Ошибка при преобразовании первого числа:", err)
			return
		}

		num2, err = rimToArab(parts[2])
		if err != nil {
			fmt.Println("Ошибка при преобразовании второго числа:", err)
			return
		}
	} else {
		num1, err = strconv.Atoi(parts[0])
		if err != nil || num1.(int) < 1 || num1.(int) > 10 {
			fmt.Println("Первое число должно быть от 1 до 10")
			return
		}

		num2, err = strconv.Atoi(parts[2])
		if err != nil || num2.(int) < 1 || num2.(int) > 10 {
			fmt.Println("Второе число должно быть от 1 до 10")
			return
		}
	}

	// Выполняем операцию в зависимости от введенного символа
	var result interface{}
	switch parts[1] {
	case "+":
		result = add(num1, num2, isRimNum1)
	case "-":
		result = subtract(num1, num2, isRimNum1)
	case "*":
		result = multiply(num1, num2, isRimNum1)
	case "/":
		if num2.(int) == 0 {
			fmt.Println("Делить на ноль нельзя")
			return
		}
		result = divide(num1, num2, isRimNum1)
	default:
		fmt.Println("Ошибка")
		return
	}

	// Проверяем результат для римских чисел
	if isRimNum1 {
		if result.(int) <= 0 {
			fmt.Println("Ошибка: результат не может быть отрицательным или равным нулю для римских чисел")
			return
		}
	}

	// Выводим результат в соответствующей форме (римские числа или арабские числа)
	if isRimNum1 {
		fmt.Println("Результат:", arabToRim(result.(int)))
	} else {
		fmt.Println("Результат:", result)
	}
}

// Функция проверки, является ли строка римским числом
func isRim(s string) bool {
	rimNum := map[string]bool{
		"I": true, "II": true, "III": true, "IV": true, "V": true,
		"VI": true, "VII": true, "VIII": true, "IX": true, "X": true,
	}

	_, ok := rimNum[s]
	return ok
}

// Функция преобразования римского числа в арабское число
func rimToArab(s string) (int, error) {
	rimNum := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	val, ok := rimNum[s]
	if !ok {
		return 0, fmt.Errorf("некорректное римское число")
	}

	return val, nil
}

// Функция преобразования арабского числа в римское число
func arabToRim(num int) string {
	rimNum := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
		10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX",
		20: "XX", 21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX",
		30: "XXX", 31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV", 36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX",
		40: "XL", 41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV", 46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX",
		50: "L", 51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX",
		60: "LX", 61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX",
		70: "LXX", 71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX",
		80: "LXXX", 81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX",
		90: "XC", 91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
	}

	return rimNum[num]
}

// Функция сложения
func add(a, b interface{}, isRimNum bool) interface{} {
	if isRimNum {
		return a.(int) + b.(int)
	} else {
		return a.(int) + b.(int)
	}
}

// Функция вычитания
func subtract(a, b interface{}, isRimNum bool) interface{} {
	if isRimNum {
		return a.(int) - b.(int)
	} else {
		return a.(int) - b.(int)
	}
}

// Функция умножения
func multiply(a, b interface{}, isRimNum bool) interface{} {
	if isRimNum {
		return a.(int) * b.(int)
	} else {
		return a.(int) * b.(int)
	}
}

// Функция деления
func divide(a, b interface{}, isRimNum bool) interface{} {
	if isRimNum {
		return a.(int) / b.(int)
	} else {
		return a.(int) / b.(int)
	}
}
