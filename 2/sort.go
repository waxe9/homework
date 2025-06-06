package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func convertStringsToNumbers(stringSlice []string) ([]int64, error) {
	var numbers []int64
	for _, str := range stringSlice {
		cleanedStr := strings.TrimSpace(str)
		if cleanedStr == "" {
			continue
		}

		num, conversionErr := strconv.ParseInt(cleanedStr, 10, 64)
		if conversionErr != nil {
			log.Printf("Предупреждение: Не удалось преобразовать '%s': %v", str, conversionErr)
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func startProgram() {
	inputData, openErr := os.Open("input.txt")
	if openErr != nil {
		log.Fatalf("Ошибка открытия файла: %v", openErr)
	}
	defer inputData.Close()

	fileScanner := bufio.NewScanner(inputData)
	fileScanner.Split(bufio.ScanWords)
	var textElements []string
	for fileScanner.Scan() {
		textElements = append(textElements, fileScanner.Text())
	}

	numericValues, parseErr := convertStringsToNumbers(textElements)
	if parseErr != nil {
		log.Fatalf("Ошибка преобразования чисел: %v", parseErr)
	}

	sort.Slice(numericValues, func(firstIndex, secondIndex int) bool {
		return numericValues[firstIndex] < numericValues[secondIndex]
	})

	resultFile, createErr := os.Create("output.txt")
	if createErr != nil {
		log.Fatalf("Ошибка создания файла: %v", createErr)
	}
	defer resultFile.Close()

	fileWriter := bufio.NewWriter(resultFile)
	for _, number := range numericValues {
		fmt.Fprintf(fileWriter, "%d ", number)
	}
	fileWriter.Flush()

	fmt.Println("Результат записан в output.txt.")
}

func main() {
	startProgram()
}
