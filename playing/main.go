package main

import (
	"fmt"
)

var x = "testing"

func textRight(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, i+1)
	}
}

func nothingToDoHere() string {
	y := "testing Stuffs"

	fmt.Println(y)
	return y
}

func main() {
	textRight(x)
	fmt.Println(nothingToDoHere())
}
