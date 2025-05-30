package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Функция сортировки массива методом пузырька
func arrangeElements(elements []int64) {
	count := len(elements)
	for step := 0; step < count; step++ {
		changed := false
		for position := 0; position < count-1; position++ {
			if elements[position] > elements[position+1] {
				elements[position], elements[position+1] = elements[position+1], elements[position]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}

func execute() {
	// Открытие файла для чтения
	fileReader, openErr := os.Open("input.txt")
	if openErr != nil {
		fmt.Println("Ошибка открытия файла:", openErr)
		os.Exit(1)
	}
	defer fileReader.Close() // Гарантированное закрытие файла

	// Чтение данных из файла
	rawData := make([]byte, 64)
	_, readErr := fileReader.Read(rawData)
	if readErr != nil && readErr != io.EOF {
		fmt.Println("Ошибка чтения файла:", readErr)
		os.Exit(1)
	}

	// Парсинг чисел из файла
	var digitSequence []int64
	for _, part1 := range strings.Split(string(rawData), " ") {
		for _, part2 := range strings.Split(part1, "\r") {
			for _, valueStr := range strings.Split(part2, "\x00") {
				if valueStr == "" {
					continue
				}
				convertedNum, parseErr := strconv.ParseInt(valueStr, 10, 64)
				if parseErr != nil {
					fmt.Printf("Ошибка преобразования '%s': %v\n", valueStr, parseErr)
					continue
				}
				digitSequence = append(digitSequence, convertedNum)
			}
		}
	}

	// Сортировка чисел
	arrangeElements(digitSequence)

	// Создание файла для результатов
	resultFile, createErr := os.Create("output.txt")
	if createErr != nil {
		fmt.Println("Ошибка создания файла:", createErr)
		os.Exit(1)
	}
	defer resultFile.Close()

	// Формирование строки результата
	var formattedResult strings.Builder
	for _, val := range digitSequence {
		fmt.Fprintf(&formattedResult, "%d ", val)
	}

	// Запись результатов в файл
	if _, writeErr := resultFile.WriteString(formattedResult.String()); writeErr != nil {
		fmt.Println("Ошибка записи в файл:", writeErr)
		os.Exit(1)
	}

	// Вывод результатов в консоль
	fmt.Println("Отсортированные числа:", digitSequence)
}

func main() {
	execute()
}
