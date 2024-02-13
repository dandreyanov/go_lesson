package lesson_1

import (
	"fmt"
	"math"
)

func Hard() {
	var vertices [][]float64
	vertices = make([][]float64, 4)

	for i := 0; i < 4; i++ {
		vertices[i] = make([]float64, 2)
		fmt.Println("Введите координаты ", i+1, "вершины:")
		vertices[i][0], vertices[i][1] = inputCoordinate()
	}

	fmt.Println("Массив координат:")
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(vertices[i][j], " ")
		}
		fmt.Println()
	}

	// Вызовите функцию для вычисления площади
	area := calculateArea(vertices[0][0], vertices[0][1], vertices[1][0], vertices[1][1], vertices[2][0], vertices[2][1], vertices[3][0], vertices[3][1])

	// Выведите результат
	fmt.Println("Площадь фигуры:", area)
}

func inputCoordinate() (float64, float64) {
	var x, y float64
	fmt.Println("Введите X:")
	fmt.Scan(&x)
	fmt.Println("Введите Y:")
	fmt.Scan(&y)
	return x, y
}

// Функция для вычисления площади фигуры
func calculateArea(x1, y1, x2, y2, x3, y3, x4, y4 float64) float64 {
	// Вычисляем длины сторон
	side1 := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	side2 := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	side3 := math.Sqrt(math.Pow(x4-x3, 2) + math.Pow(y4-y3, 2))
	side4 := math.Sqrt(math.Pow(x1-x4, 2) + math.Pow(y1-y4, 2))

	// Проверка на прямоугольник
	if (side1 == side3 && side2 == side4) || (side1 == side2 && side3 == side4) {
		// Вычисление площади прямоугольника
		width := math.Min(side1, side2)
		height := math.Min(side3, side4)
		fmt.Println("Это прямоугольник")
		return width * height
	}

	// Проверка на треугольник
	if (x1 == x2 && y1 == y2) || (x1 == x3 && y1 == y3) || (x1 == x4 && y1 == y4) {
		// Вычисление площади треугольника
		base := abs(x2 - x1)
		height := abs(y2 - y1)
		fmt.Println("Это треугольник")
		return 0.5 * base * height
	}

	// Если фигура не является ни прямоугольником, ни треугольником, возвращаем 0
	return 0
}

// Функция для вычисления абсолютного значения
func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}
