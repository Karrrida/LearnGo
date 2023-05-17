package main

import "fmt"

//type SuperBike struct {
//	On          bool
//	Ammo, Power int
//}
//
//func (s *SuperBike) Shoot() bool {
//	if s.On && s.Ammo > 0 {
//		s.Ammo--
//		return true
//	}
//	return false
//}
//
//func (s *SuperBike) RideBike() bool {
//	if s.On && s.Power > 0 {
//		s.Power--
//		return true
//	}
//	return false
//}

func main() {
	//newOne := SuperBike{true, 1, 1}
	//testStruct := &newOne
	//fmt.Println(testStruct)
	//testStruct.On, testStruct.Ammo, testStruct.Power = true, 10, 10
	//testStruct.Shoot()
	//testStruct.RideBike()
	//fmt.Println(testStruct)
	//inp, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//isLegalString := "Wrong"
	//sr := []rune(inp)
	//
	//if unicode.IsUpper(sr[0]) && string(sr[len(sr)-3]) == "." {
	//	isLegalString = "Right"
	//}
	//fmt.Println(isLegalString)

	//var inpStr string
	//
	//fmt.Scan(&inpStr)
	//
	//sr2 := []rune(inpStr)
	//
	//if string(sr2[0]) == string(sr2[len(sr2)-1]) {
	//	fmt.Println("Палиндром")
	//} else {
	//	fmt.Println("Нет")
	//}

	var str string

	fmt.Scan(&str)
	//var resStr string
	rs := []rune(str)
	var resRs []rune
	for idx, val := range rs {
		if idx%2 != 0 {
			resRs = append(resRs, val)
		}
	}
	fmt.Print(string(resRs))

}

//func sumInt(a ...int) (lenInt, sum int) {
//
//	for idx := range a {
//		lenInt++
//		sum += a[idx]
//	}
//
//	return
//}

//func vote(x int, y int, z int) int {
//	var result int = 0
//
//	if x+y+z > 1 {
//		result = 1
//	}
//
//	return result
//}
