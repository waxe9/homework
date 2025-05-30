package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	// Открываем файл с названием "input.txt" для чтения
	inputFile, fileError := os.Open("input.txt")
	if fileError != nil {
		fmt.Println("Ошибка при открытии файла:", fileError)
		return
	}
	// Закрываем файл после завершения работы с ним
	defer inputFile.Close()

	// Создаем сканер для построчного чтения файла
	fileScanner := bufio.NewScanner(inputFile)
	var coefficientA, coefficientB, coefficientC int
	// Читаем первую строку из файла
	if fileScanner.Scan() {
		// Парсим три числа из строки в переменные coefficientA, coefficientB, coefficientC
		fmt.Sscanf(fileScanner.Text(), "%d %d %d", &coefficientA, &coefficientB, &coefficientC)
	} else if err := fileScanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Вычисляем дискриминант квадратного уравнения по формуле: B² - 4AC
	discriminant := coefficientB*coefficientB - 4*coefficientA*coefficientC

	// Выводим введенные коэффициенты и вычисленный дискриминант
	fmt.Printf("Коэффициент A: %d, Коэффициент B: %d, Коэффициент C: %d\n", coefficientA, coefficientB, coefficientC)
	fmt.Printf("Дискриминант: %d\n", discriminant)

	// Анализируем значение дискриминанта для определения количества корней
	if discriminant > 0 {
		// Если дискриминант положительный - два действительных корня
		root1 := (-float64(coefficientB) + math.Sqrt(float64(discriminant))) / (2 * float64(coefficientA))
		root2 := (-float64(coefficientB) - math.Sqrt(float64(discriminant))) / (2 * float64(coefficientA))
		fmt.Printf("Корни уравнения: первый корень = %.2f, второй корень = %.2f\n", root1, root2)
	} else if discriminant == 0 {
		// Если дискриминант равен нулю - один действительный корень
		root := -float64(coefficientB) / (2 * float64(coefficientA))
		fmt.Printf("Единственный корень уравнения: x = %.2f\n", root)
	} else {
		// Если дискриминант отрицательный - действительных корней нет
		fmt.Println("Корней нет (дискриминант меньше нуля)")
	}
}
