package lesson_1

import "fmt"

func Two() {
	var arr [5]int
	fmt.Println("Введите пять чисел через пробел:")
	fmt.Scan(&arr[0], &arr[1], &arr[2], &arr[3], &arr[4])

	min := arr[0]
	max := arr[0]
	sum := 0

	for i := 0; i < 5; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
		sum += arr[i]
	}

	fmt.Println("Минимальное число:", min)
	fmt.Println("Максимальное число:", max)
	fmt.Println("Среднее арифметическое:", sum/5)
}
