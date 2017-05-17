package main

import (
	"fmt"
)

var x = 0

//normal function
func addition() int {
	x++
	return x
}

//function which returns a function
func wrapper() func() int {
	z := 9
	return func() int {
		z++
		return z
	}

}

func main() {
	y := 10
	//inner function
	substract := func() int {
		y--
		return y
	}
	fmt.Println("testing closure")
	fmt.Println(substract(), " substraction example -1")
	fmt.Println(addition(), " incremented example +1")
	fmt.Println(addition(), " increment example +1")


	incrementViaWrapper := wrapper()
	fmt.Println(incrementViaWrapper())

}
