package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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
	fileReader, openErr := os.Open("input.txt")
	if openErr != nil {
		fmt.Println("Ошибка открытия файла:", openErr)
		os.Exit(1)
	}
	defer fileReader.Close()

	rawData := make([]byte, 64)
	_, readErr := fileReader.Read(rawData)
	if readErr != nil && readErr != io.EOF {
		fmt.Println("Ошибка чтения файла:", readErr)
		os.Exit(1)
	}

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

	arrangeElements(digitSequence)

	resultFile, createErr := os.Create("output.txt")
	if createErr != nil {
		fmt.Println("Ошибка создания файла:", createErr)
		os.Exit(1)
	}
	defer resultFile.Close()

	var formattedResult strings.Builder
	for _, val := range digitSequence {
		fmt.Fprintf(&formattedResult, "%d ", val)
	}

	if _, writeErr := resultFile.WriteString(formattedResult.String()); writeErr != nil {
		fmt.Println("Ошибка записи в файл:", writeErr)
		os.Exit(1)
	}

	fmt.Println("Отсортированные числа:", digitSequence)
}

func main() {
	execute()
}
