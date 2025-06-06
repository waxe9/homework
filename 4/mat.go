package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadMatrixData(filePath string) [][]float64 {
	fileHandle, openErr := os.Open(filePath)
	if openErr != nil {
		panic(openErr)
	}
	defer fileHandle.Close()

	var dataGrid [][]float64
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		if currentLine == "" {
			continue
		}

		lineElements := strings.Fields(currentLine)
		rowData := make([]float64, len(lineElements))

		for index, elementStr := range lineElements {
			elementValue, parseErr := strconv.ParseFloat(elementStr, 64)
			if parseErr != nil {
				panic(parseErr)
			}
			rowData[index] = elementValue
		}

		dataGrid = append(dataGrid, rowData)
	}

	return dataGrid
}

func saveCalculationResults(outputPath string, matrixDet, matrixTrace float64, transposedGrid [][]float64) {
	resultFile, createErr := os.Create(outputPath)
	if createErr != nil {
		panic(createErr)
	}
	defer resultFile.Close()

	fileWriter := bufio.NewWriter(resultFile)
	defer fileWriter.Flush()

	fmt.Fprintf(fileWriter, "Определитель: %.2f\n", matrixDet)
	fmt.Fprintf(fileWriter, "След: %.2f\n", matrixTrace)
	fmt.Fprintln(fileWriter, "Транспонированная матрица:")

	for _, rowValues := range transposedGrid {
		for _, cellValue := range rowValues {
			fmt.Fprintf(fileWriter, "%.2f ", cellValue)
		}
		fmt.Fprintln(fileWriter)
	}
}

func calculateMatrixDeterminant(inputGrid [][]float64) float64 {
	dimension := len(inputGrid)
	if dimension == 1 {
		return inputGrid[0][0]
	}
	if dimension == 2 {
		return inputGrid[0][0]*inputGrid[1][1] - inputGrid[0][1]*inputGrid[1][0]
	}

	totalDet := 0.0
	for currentCol := 0; currentCol < dimension; currentCol++ {
		subMatrix := make([][]float64, dimension-1)
		for i := range subMatrix {
			subMatrix[i] = make([]float64, dimension-1)
		}

		for row := 1; row < dimension; row++ {
			colIndex := 0
			for col := 0; col < dimension; col++ {
				if col == currentCol {
					continue
				}
				subMatrix[row-1][colIndex] = inputGrid[row][col]
				colIndex++
			}
		}

		coeff := 1.0
		if currentCol%2 == 1 {
			coeff = -1.0
		}
		totalDet += coeff * inputGrid[0][currentCol] * calculateMatrixDeterminant(subMatrix)
	}
	return totalDet
}

func computeMatrixTrace(inputGrid [][]float64) float64 {
	traceSum := 0.0
	for i := 0; i < len(inputGrid); i++ {
		traceSum += inputGrid[i][i]
	}
	return traceSum
}

func performMatrixTranspose(inputGrid [][]float64) [][]float64 {
	size := len(inputGrid)
	transposedResult := make([][]float64, size)
	for i := range transposedResult {
		transposedResult[i] = make([]float64, size)
	}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			transposedResult[col][row] = inputGrid[row][col]
		}
	}
	return transposedResult
}

func main() {
	matrixData := loadMatrixData("input.txt")

	matrixSize := len(matrixData)
	for _, row := range matrixData {
		if len(row) != matrixSize {
			fmt.Println("Ошибка: матрица должна быть квадратной")
			return
		}
	}

	detValue := calculateMatrixDeterminant(matrixData)
	traceValue := computeMatrixTrace(matrixData)
	transposedMatrix := performMatrixTranspose(matrixData)

	saveCalculationResults("output.txt", detValue, traceValue, transposedMatrix)

	fmt.Printf("Результаты сохранены в output.txt\n")
	fmt.Printf("Определитель: %.2f\n", detValue)
	fmt.Printf("След: %.2f\n", traceValue)
	fmt.Println("Транспонированная матрица:")
	for _, row := range transposedMatrix {
		fmt.Println(row)
	}
}
