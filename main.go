package main

import (
	"fmt"
	"math"
)

func main() {
	//Переменные и типы данных
	//Обычное объявление переменных: var variable_name variable_type var a int / var b string
	var hello string = "Hello, Go!" // variable type string, для строк можно использовать двойные либо косые кавычки
	var a int = 2023                // variable type integer
	var b float32 = 2.33
	var symbol rune = 'c' // type int32/rune можно использовать для хранения символов, для типа rune использовать одинарные кавычки

	fmt.Println(hello)
	fmt.Println(a)
	fmt.Println(symbol)                    // input 'c' output -> 99
	fmt.Println(string(symbol), "here!!!") // "c" in unicode is 99
	fmt.Println(b)

	// Можно объявить несколько переменных в одном блоке var:
	var (
		name string = "Max"
		age  int    = 32
	)
	fmt.Println(name, age)
	//Краткий способ объявления переменных (доступен только внутри функций) variable_name := value
	c := 5
	fmt.Println(c)
	//===============================================

	//Арифметические операции, ввод/вывод данных
	//Деление /
	var devideInt int = 10 / 6         //если при делении оба операнда - целые числа, то результат будет округлен до целого числа
	var devideFloat float32 = 10.0 / 6 //Для получения вещественного числа необходимо чтобы хотя бы один операнд был вещественным числом, и соответственно тип данных тоже
	fmt.Println(devideInt)
	fmt.Println(devideFloat)

	//Остаток от деления %
	var remainderDivision int = 10 % 3
	fmt.Println(remainderDivision)

	//Постфиксный инкремент/постфиксный декремент a++ a--
	var posfixInc int = 1
	posfixInc++
	var posfixDec int = 2
	posfixDec--

	fmt.Println(posfixInc, posfixDec)
	//==============================================
	//Чтение данных с консоли fmt.Scan(&a) &a - адрес на переменную куда сохранится введенное значение

	var userName string
	var userAge int

	fmt.Print("Введите имя: ")
	fmt.Scan(&userName)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&userAge)

	fmt.Println(userName, userAge)
	//=============================================
	//Условная конструкция if

	aif := 6
	bif := 7

	//Обычная конструкция if
	if aif < bif {
		fmt.Println("a less b")
	}

	// оператор if может начинаться с инструкции, которая будет выполнена перед проверкой условия
	// переменные объявленные в это блоке доступны только в области видимости текущего if
	if vif := math.Pow(5, 2); vif < 26 {
		//...
	}

	// if else if else
	if aif < bif {
		//...
	} else if aif > bif {
		//...
	} else {
		//...
	}

	//switch
	var ic int = 2
	switch ic {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough // если fallthrough указан в текущем case то следующий case выполнится не смотря на ложность условия
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unkonown Number")
	}
	//============================================
	//Циклы for

	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println("FOR cycle", sum)

	//============================================
	//Форматированный вывод
	// fmt.Printf("%modif", values)
	var aFormatOutput rune = 'Ы'
	fmt.Printf("%q", aFormatOutput)
	var floatFormatOutput float32 = 22.123
	fmt.Printf("%9.2f", floatFormatOutput)
}
