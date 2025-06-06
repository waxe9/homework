package main

import (
	"fmt"
	"unicode"
)

func performCipher(inputText string, offset int) string {
	encryptedResult := ""

	for _, currentSymbol := range inputText {
		if unicode.IsUpper(currentSymbol) {
			encryptedResult += string((int(currentSymbol)-'А'+offset)%33 + 'А')
		} else if unicode.IsLower(currentSymbol) {
			encryptedResult += string((int(currentSymbol)-'а'+offset)%33 + 'а')
		} else {
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
