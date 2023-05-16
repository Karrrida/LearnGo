package main

import (
	"fmt"
)

type SuperBike struct {
	On          bool
	Ammo, Power int
}

func (superBike *SuperBike) Shoot() bool {
	if superBike.On && superBike.Ammo > 0 {
		superBike.Ammo--
		return true
	}
	return false
}

func (superBike *SuperBike) RideBike() bool {
	if superBike.On && superBike.Power > 0 {
		superBike.Power--
		return true
	}
	return false
}

func main() {

	newOne := SuperBike{true, 1, 1}
	testStruct := &newOne
	fmt.Println(testStruct)
	testStruct.On, testStruct.Ammo, testStruct.Power = true, 10, 10

	testStruct.Shoot()
	testStruct.RideBike()
	fmt.Println(testStruct)
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
