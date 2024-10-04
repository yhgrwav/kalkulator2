package calc

import (
	"errors"
	"strconv"
	"strings"
)

// Calculate выполняет операцию на основе введенного выражения
func Calculate(expression string) (string, error) {
	expression = strings.TrimSpace(expression)

	if !strings.HasPrefix(expression, "\"") {
		return "", errors.New("первый операнд должен быть строкой в кавычках")
	}

	strEndIndex := strings.Index(expression[1:], "\"") + 1
	if strEndIndex == 0 {
		return "", errors.New("ошибка: неверный формат строки")
	}

	str1 := expression[1:strEndIndex]
	if len(str1) > 10 {
		return "", errors.New("первый операнд должен быть не более 10 символов")
	}

	remainingExpr := strings.TrimSpace(expression[strEndIndex+1:])

	if len(remainingExpr) < 1 {
		return "", errors.New("ошибка: отсутствует операция")
	}

	operation := string(remainingExpr[0])
	remainingExpr = strings.TrimSpace(remainingExpr[1:])

	str2OrNumber := strings.TrimSpace(remainingExpr)

	switch operation {
	case "+":
		if !strings.HasPrefix(str2OrNumber, "\"") || !strings.HasSuffix(str2OrNumber, "\"") {
			return "", errors.New("второй операнд должен быть строкой для операции сложения")
		}
		str2 := str2OrNumber[1 : len(str2OrNumber)-1]
		return handleLength(str1 + str2), nil

	case "-":
		if !strings.HasPrefix(str2OrNumber, "\"") || !strings.HasSuffix(str2OrNumber, "\"") {
			return "", errors.New("второй операнд должен быть строкой для операции вычитания")
		}
		str2 := str2OrNumber[1 : len(str2OrNumber)-1]
		return handleLength(strings.Replace(str1, str2, "", 1)), nil

	case "*":
		num, err := strconv.Atoi(str2OrNumber)
		if err != nil || num < 1 || num > 10 {
			return "", errors.New("второй операнд должен быть числом от 1 до 10")
		}
		return handleLength(strings.Repeat(str1, num)), nil

	case "/":
		num, err := strconv.Atoi(str2OrNumber)
		if err != nil || num < 1 || num > 10 {
			return "", errors.New("второй операнд должен быть числом от 1 до 10")
		}
		if len(str1)/num == 0 {
			return "", errors.New("результат деления слишком короткий")
		}
		return handleLength(str1[:len(str1)/num]), nil

	default:
		return "", errors.New("неподдерживаемая операция")
	}
}

// Обрезает строку, если её длина больше 40 символов
func handleLength(result string) string {
	if len(result) > 40 {
		return result[:40] + "..."
	}
	return result
}
