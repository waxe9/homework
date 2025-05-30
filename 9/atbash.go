package main

import (
	"fmt"
	"unicode"
)

// cipherTransform выполняет шифрование входящей строки методом Атбаш
// Атбаш - это метод шифрования, заменяющий буквы на симметричные в алфавите
// Например: A -> Z, B -> Y, ..., Z -> A (аналогично для строчных букв)
func cipherTransform(inputStr string) string {
	transformedStr := "" // Здесь будет храниться результат шифрования

	// Перебираем каждый символ входной строки
	for _, currentRune := range inputStr {
		// Обрабатываем заглавные буквы
		if unicode.IsUpper(currentRune) {
			// Формула для замены: 'Z' - (текущий символ - 'A')
			// Например: для 'B' (66): 'Z'(90) - (66 - 65) = 89 → 'Y'
			transformedStr += string(('Z' - (int(currentRune) - 'A')))
		} else if unicode.IsLower(currentRune) {
			// Аналогичная обработка для строчных букв
			transformedStr += string(('z' - (int(currentRune) - 'a')))
		} else {
			// Все остальные символы (не буквы) добавляем без изменений
			transformedStr += string(currentRune)
		}
	}
	return transformedStr
}

// executeProgram содержит основную логику работы программы
func executeProgram() {
	var userInput string // Переменная для хранения ввода пользователя

	// Запрашиваем у пользователя текст для шифрования
	fmt.Print("Введите текст: ")
	fmt.Scanln(&userInput) // Считываем ввод

	// Выводим результат шифрования
	fmt.Println("Результат: ", cipherTransform(userInput))
}

// Точка входа в программу
func main() {
	executeProgram()
}
