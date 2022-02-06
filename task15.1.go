package main

import (
	"fmt"
)

func main() {
	fmt.Println("Задание 1. Подсчёт чётных и нечётных чисел в массиве")

	var num [10]int
	var even int

	for i, _ := range num {
		fmt.Println("Введите число")
		fmt.Scan(&num[i])
	}

	for _, num := range num {
		if num%2 == 0 {
			even++
		}
	}

	fmt.Println("чётных —", even, " нечётных —", 10-even)
}
