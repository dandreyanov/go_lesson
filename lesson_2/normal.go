package lesson_2

import "fmt"

func Normal() {
	var arr [5]int
	fmt.Println("Введите пять возрастов	через пробел:")
	fmt.Scan(&arr[0], &arr[1], &arr[2], &arr[3], &arr[4])
	checkAdult(arr)
}

func checkAdult(arr [5]int) {
	for _, age := range arr {
		if age >= 18 {
			fmt.Println("Есть совершеннолетний")
			return
		}
	}
	fmt.Println("Нет совершеннолетних")

}
