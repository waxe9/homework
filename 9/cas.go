package main

import (
	"fmt"
	"unicode"
)

// performCipher осуществляет шифрование текста методом Цезаря
func performCipher(inputText string, offset int) string {
	encryptedResult := ""

	for _, currentSymbol := range inputText {
		if unicode.IsUpper(currentSymbol) {
			// Шифрование заглавных букв русского алфавита (33 буквы)
			encryptedResult += string((int(currentSymbol)-'А'+offset)%33 + 'А')
		} else if unicode.IsLower(currentSymbol) {
			// Шифрование строчных букв русского алфавита
			encryptedResult += string((int(currentSymbol)-'а'+offset)%33 + 'а')
		} else {
			// Оставляем другие символы без изменений
			encryptedResult += string(currentSymbol)
		}
	}
	return encryptedResult
}

func runProgram() {
	var userText string
	var cipherOffset int

	fmt.Print("Введите текст для шифрования: ")
	fmt.Scanln(&userText)

	fmt.Print("Введите величину сдвига: ")
	fmt.Scanln(&cipherOffset)

	fmt.Println("Зашифрованный результат:", performCipher(userText, cipherOffset))
}

func main() {
	runProgram()
}
