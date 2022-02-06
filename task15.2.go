package main

import (
	"fmt"
)

func reverse(massif [10]int) (reverse [10]int) {
	a := 9
	for _, r := range massif {
		reverse[a] = r
		a--
	}
	return
}

func main() {
	fmt.Println("Задание 2. Функция, реверсирующая массив")

	var num [10]int

	for i, _ := range num {
		fmt.Println("Введите число")
		fmt.Scan(&num[i])
	}

	num = reverse(num)

	for i, _ := range num {
		if i != 9 {
			fmt.Print(num[i], ", ")
		} else {
			fmt.Print(num[i], ".")
		}
	}
}
