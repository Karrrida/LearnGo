package main

import "fmt"

func main() {

	var N int

	fmt.Scan(&N)

	var fibPrev, fibNext = 0, 1
	n := 1

	for fibNext <= N {
		if fibNext == N {
			fmt.Println(n)
			break
		}

		fibPrev, fibNext = fibNext, fibPrev+fibNext
		n += 1
		if fibNext == N {
			fmt.Println(n)
			break
		}

	}

	if N > fibPrev && N < fibNext {
		fmt.Println("-1")
	}
}

//var N string
//var target string
//fmt.Scan(&N, &target)
//
//for i := 0; i < len(N); i++ {
//	elem := string(rune(N[i]))
//
//	if elem != target {
//		fmt.Print(elem)
//	}
//
//}

//var N int
//var i float64 = 0
//fmt.Scan(&N)
//
//for ; ; i++ {
//	var pow float64 = i
//	var elem float64 = 2
//	elemInPow := int(math.Pow(elem, pow))
//
//	if elemInPow <= N {
//		fmt.Print(elemInPow, " ")
//	} else if elemInPow > N {
//		break
//	}
//}

/*
	var n int

	fmt.Scan(&n)

	switch {
	case (n > 4 && n <= 20) || (n%10 > 4 && n%10 < 9) || n%10 == 0:
		fmt.Println(n, "korov")
	case n == 1 || n%10 == 1:
		fmt.Println(n, "korova")
	case (n > 1 && n <= 4) || (n%10 > 1 && n%10 < 5):
		fmt.Println(n, "korovy")
	}
*/

//var a, b int
//
//fmt.Scan(&a, &b)
//
//for i := b; i >= a; i-- {
//	if i%7 == 0 {
//		fmt.Println(i)
//		break
//	} else if i == a && i%7 != 0 {
//		fmt.Println("NO")
//	}
//}
// var n int
// fmt.Scan(&n)

// digitRoot := 1 + ((n - 1) % 9) // dr(n) = 1 + ((n - 1) % 9) - формула вычисления цифрового корня

// fmt.Println(digitRoot)

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

/*
	//Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.
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
