package lesson_1

import "fmt"

func Two() {
	var num1, num2 int
	fmt.Print("Введите первое число: ")
	num1 = inputNumber()
	fmt.Print("Введите второе число: ")
	num2 = inputNumber()
	switch {
	case num1 > num2:
		fmt.Println("Первое число больше")
	case num2 > num1:
		fmt.Println("Второе число больше")
	case num1 == num2:
		fmt.Println("Числа равны")
	}
	result := float32(num1+num2) / 2
	fmt.Println("Среднее арифметическое = ", result)
}

func inputNumber() int {
	var num int
	fmt.Scan(&num)
	return num
}
