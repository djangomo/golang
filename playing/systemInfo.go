package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("getting Time")
	fmt.Println(now)
	x := now.Weekday()
	fmt.Println(x)
}
