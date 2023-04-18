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

	//===========================================
	//Массивы и срезы

	//Массивы
	//Объявление массива
	//var aArray [3]int

	var aArray [3]int = [3]int{1, 2, 3}
	//bArray := [3]int{1, 2, 3}
	//cArray := [...]int{1, 2, 3} // ... многоточие вместо указания длины массива, Go возьмет длину массива при инициализации
	//dArray := [3]int{1: 12} // int{1: 12} позволяет указанному индексу присвоить значение по умолчанию
	fmt.Println(aArray)

	// Сравнение массивов: Можно сравнить массивы только если они одного типа (массивы одинаковой длины и содержат элементы одного типа)
	// a := [3]int{1, 2, 3}
	// b := [3]int{1, 2, 3}
	// c := [3]int{3, 2, 1}
	// fmt.Println(a == b) // true
	// fmt.Println(a == c) // false

	//Итерация массива
	//Можно использовать стандартный вариант for
	//Так же есть вариант с range

	var iArray = [5]int{1, 2, 3, 4, 5}

	for idx, elem := range iArray {
		fmt.Printf("Elem with index %d: %d\n", idx, elem)
	}
	//Range возвращает два объекта: индекс элемента и копию значения элемента
	//Можно опустить любой из этих объектов если указать символ "_", также если мы хотис использовать только индекс то указание элемента можно не указывать вовсе
	for idx := range iArray { //Здесь мы не указываем временную переменную для элемента массива, и получаем только индекс
		fmt.Println(iArray[idx])
	}

	for _, elem := range iArray { //Здесь используем только элемент, опуская индекс
		fmt.Println(elem)
	}
	/*
		Необходимо запомнить, что в качестве второго значения range
		возвращает копию элемента массива, это может быть важно,
		 если в цикле мы хотим изменить массив.
		 В этом случае мы должны обращаться к элементам массива по индексу:
		for idx := range a {
			a[idx] = 100
		}
	*/

	//Срезы
	var iSlice = []int{1: 12} // Создание среза с явным указанием значения по индексу
	var iSliceEmpty []int     // Создание пустого среза
	fmt.Println(iSlice, len(iSlice), cap(iSlice))
	fmt.Println(iSliceEmpty, len(iSliceEmpty), cap(iSliceEmpty))

	var sliceMake = make([]int, 10, 10) //Создание среза через метод make([]T, length, capacity)
	fmt.Println(sliceMake)

	//Оператор среза s[i:j] создает из последовательности s новый срещу который содержит элементы посл-ти s c i по j-1
	//Если не указан i то по умолчание берется значение 0, если не указан j то берется последний элемент последовательности
	sliceInitialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := sliceInitialUsers[2:6]
	fmt.Println(users1)

	//Стандартные методы работы со срезами
	//append(slice []Type, elems...Type) []Type
	sliceAppend := []int{1, 2, 3}
	sliceAppend = append(sliceAppend, 4, 5)
	//baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//baseSlice := baseArray[5:8]

	//pointer := fmt.Sprintf("%p", baseSlice)

	removeArray := []int{1, 2, 3, 4, 5}
	removeArray = append(removeArray[0:2], removeArray[3:]...) //"удаление" элемента массива, на самом деле создание
	fmt.Println(removeArray)

	/*
		func copy(dst, src []Type) int
		Copy принимает срез-назначение и срез источник, а возвращает число скопированных элементов:

		a := []int{1, 2, 3}
		b := make([]int, 3, 3)
		n := copy(b, a)
	*/

}
