package main

import (
	"fmt"
)

func main() {

	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a+b > c && a+c > b && b+c > a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}

	/*
		Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.
			var k, h, m int

			fmt.Scan(&k)
			h = k / 3600
			m = (k % 3600) / 60
			fmt.Printf("It is %d hours %d minutes.", h, m)

	*/
	/*
		Дано трехзначное число. Найдите сумму его цифр.
		var num int
		fmt.Scan(&num)

		num1 := num / 100 % 10
		num2 := num / 10 % 10
		num3 := num % 10

		fmt.Println(num1 + num2 + num3)

	*/

}
