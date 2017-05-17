package main

import "fmt"

func test(s int)  {
	var x int = 0
	x = s+s
	fmt.Println(x)
}

func main() {
	fmt.Println("testing")
	test(10)
}
