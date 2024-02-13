package lesson_2

import "fmt"

func Easy() {
	var age int
	fmt.Println("Введите возраст")
	fmt.Scan(&age)
	ageDetector(age)
}

func ageDetector(age int) {
	if age >= 18 {
		fmt.Println("Совершеннолетний")
	} else {
		fmt.Println("Несовершеннолетний")
	}
}
