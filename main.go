package main

import (
	"awesomeProject3/calc"
	"awesomeProject3/input"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Введите выражение (например, \"hello\" + \"world\"): ")

	userInput, err := input.ReadInput()
	if err != nil {
		log.Fatal("Ошибка ввода:", err)
	}

	result, err := calc.Calculate(userInput)
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	fmt.Println("Результат:", result)
}
