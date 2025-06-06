package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	inputFile, fileError := os.Open("input.txt")
	if fileError != nil {
		fmt.Println("Ошибка при открытии файла:", fileError)
		return
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	var coefficientA, coefficientB, coefficientC int
	if fileScanner.Scan() {
		fmt.Sscanf(fileScanner.Text(), "%d %d %d", &coefficientA, &coefficientB, &coefficientC)
	} else if err := fileScanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	discriminant := coefficientB*coefficientB - 4*coefficientA*coefficientC

	fmt.Printf("Коэффициент A: %d, Коэффициент B: %d, Коэффициент C: %d\n", coefficientA, coefficientB, coefficientC)
	fmt.Printf("Дискриминант: %d\n", discriminant)

	if discriminant > 0 {
		root1 := (-float64(coefficientB) + math.Sqrt(float64(discriminant))) / (2 * float64(coefficientA))
		root2 := (-float64(coefficientB) - math.Sqrt(float64(discriminant))) / (2 * float64(coefficientA))
		fmt.Printf("Корни уравнения: первый корень = %.2f, второй корень = %.2f\n", root1, root2)
	} else if discriminant == 0 {
		root := -float64(coefficientB) / (2 * float64(coefficientA))
		fmt.Printf("Единственный корень уравнения: x = %.2f\n", root)
	} else {
		fmt.Println("Корней нет (дискриминант меньше нуля)")
	}
}
