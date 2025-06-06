package main

import (
	"fmt"
	"unicode"
)

func cipherTransform(inputStr string) string {
	transformedStr := ""

	for _, currentRune := range inputStr {
		if unicode.IsUpper(currentRune) {
			transformedStr += string(('Z' - (int(currentRune) - 'A')))
		} else if unicode.IsLower(currentRune) {
			transformedStr += string(('z' - (int(currentRune) - 'a')))
		} else {
			transformedStr += string(currentRune)
		}
	}
	return transformedStr
}

func executeProgram() {
	var userInput string

	fmt.Print("Введите текст: ")
	fmt.Scanln(&userInput)
	fmt.Println("Результат: ", cipherTransform(userInput))
}

func main() {
	executeProgram()
}
