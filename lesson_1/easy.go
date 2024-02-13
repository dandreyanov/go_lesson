package lesson_1

import "fmt"

func Easy() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}
