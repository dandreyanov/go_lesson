package lesson_2

import (
	"errors"
	"fmt"
)

func Hard() {
	people := map[string]int{
		"Иван":  18,
		"Петр":  15,
		"Анна":  20,
		"Антон": 17,
		"Елена": 45,
		"Вова":  3,
	}

	fmt.Println("Список людей:", people)

	findChild(people)

}

func findChild(people map[string]int) {
	var underage []string
	var err error

	for name, age := range people {
		if age < 18 {
			underage = append(underage, name)
		}
	}

	if len(underage) > 0 {
		err = errors.New("Ошибка: есть несовершеннолетние")
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Все совершеннолетние")
	}

	if len(underage) > 0 {
		fmt.Println("Имена несовершеннолетних:", underage)
	}
}
