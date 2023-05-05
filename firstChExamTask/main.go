package main

import "fmt"

func main() {

	/*
			//Найдите количество минимальных элементов в последовательности.


				var N, min, a, count int

				fmt.Scan(&N)
				min = a

				for i := 0; i < N; i++ {
					fmt.Scan(&a)
					if i == 0 || min > a {
						min = a
						count = 1
					} else if min == a {
						count++
					}
				}

		fmt.Println(count)

			var inp, count int
			fmt.Scan(&inp)
			var arr = make([]int, inp, inp)

			for i := 0; i < inp; i++ {
				fmt.Scan(&arr[i])
			}

			temp := arr[0]
			for i := 0; i < len(arr); i++ {
				elem := arr[i]

				if temp > elem {
					temp = elem
					count = 1
				} else if temp == elem {
					count++
				}

			}
			fmt.Println(count)
	*/

	//Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.
	var k, h, m int

	fmt.Scan(&k)
	h = k / 3600
	m = (k % 3600) / 60
	fmt.Printf("It is %d hours %d minutes.", h, m)

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
