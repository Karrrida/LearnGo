package main

import (
	"fmt"
	"math"
	//"strings"
)

func main() {

	var num float64
	fmt.Scan(&num)

	if num > 0 && num < 10000 {
		num = math.Pow(num, 2)
		fmt.Printf("%.4f", num)
	} else if num > 10000 {
		fmt.Printf("%e", num)
	} else if num <= 0 {
		fmt.Printf("число %2.2f не подходит", num)
	}

	/*
		var (
			str1, str2, result string
		)

		fmt.Scan(&str1, &str2)

		for i := 0; i < len(str1); i++ {
			var elemI = string(str1[i])

			for j := 0; j < len(str2); j++ {
				var elemJ = string(str2[j])
				if elemI == elemJ {
					result += elemI + " "
					continue
				}

			}
		}

		fmt.Println(strings.TrimSpace(result))
	*/

	/*
		var (
			x, p, y int
			//t       int = 0
		)

		fmt.Scan(&x, &p, &y)

		for i := 1; i < y; i++ {
			x += x * p / 100
			if x >= y {
				fmt.Println(i)
				break
			}

		}
	*/

	/*
		var n int

		for fmt.Scan(&n); n <= 100; fmt.Scan(&n) {
			if n < 10 {
				continue
			}

			fmt.Println(n)
		}
	*/

	/*
		var (
			n, c, d int
		)

		fmt.Scan(&n, &c, &d)

		for i := 1; i <= n; i++ {
			if i%c == 0 && i%d != 0 {
				fmt.Println(i)
				break
			}
			continue
		}
	*/

	/*
		var (
			x, a, i int
		)
		fmt.Scan(&a)
		for ; a != 0; fmt.Scan(&a) {
			if a > x {
				i = 1
				x = a
			} else if a == x {
				i++
			}
		}
		if a == 0 {
			fmt.Println(i)
		}
	*/

	/*
		var a int
		var sum int = 0

		fmt.Scan(&a)

		for i := 1; i <= a; i++ {
			var num int
			fmt.Scan(&num)
			if (num > 9 && num < 100) && num%8 == 0 {
				sum += num
			}
		}
		fmt.Println(sum)
	*/

	/*
		var (
			sum  = 0
			A, B int
		)

		fmt.Println("Input first number: ")
		fmt.Scan(&A)
		fmt.Println("Input second number: ")
		fmt.Scan(&B)

		for A <= B {
			sum += A
			A++
		}
		fmt.Println(sum)
	*/

	/*
		for i := 1; i <= 10; i++ {
			fmt.Println(i * i)
		}
	*/

	/*
		var input int
		fmt.Scan(&input)
		if input%400 == 0 || (input%4 == 0 && input%100 != 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	*/

	/*
		var a string
		fmt.Scan(&a)

		var first = string(a[0])
		var second = string(a[1])
		var third = string(a[2])

		if first == second || second == third || third == first {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	*/

	/*
		var num int

		fmt.Scan(&num)

		first_three := (num / 100000) + (num / 10000 % 10) + (num / 1000 % 10)
		second_three := (num / 100 % 10) + (num / 10 % 10) + (num % 10)

		if first_three == second_three {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	*/

}
