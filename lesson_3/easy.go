package lesson_3

import (
	"fmt"
)

type Subject struct {
	Name       string
	Instructor string
	Hours      int
}

type School struct {
	Subjects []Subject
}

func (s *School) Add(name, instructor string, hours int) {
	subject := Subject{Name: name, Instructor: instructor, Hours: hours}
	s.Subjects = append(s.Subjects, subject)
}

func (s *School) Delete(name string) {
	for i, subject := range s.Subjects {
		if subject.Name == name {
			s.Subjects = append(s.Subjects[:i], s.Subjects[i+1:]...)
			break
		}
	}
}

func (s *School) List() {
	fmt.Println("Учебные предметы:", len(s.Subjects), "шт")
	for _, subject := range s.Subjects {
		fmt.Printf("Предмет: %s, Преподаватель: %s, Часы: %d\n", subject.Name, subject.Instructor, subject.Hours)
	}
}

func Easy() {
	var target string
	school := School{}
	fmt.Println("\n\nДобавляем предметы")
	// Добавление предметов
	school.Add("Математика", "Иванов", 40)
	school.Add("Физика", "Петров", 30)
	school.Add("История", "Сидоров", 20)

	// Вывод списка предметов
	school.List()

	// Удаление предмета
	fmt.Println("\nКакой предмет удалим?")
	fmt.Scan(&target)
	school.Delete(target)

	// Вывод списка предметов после удаления
	school.List()
}
