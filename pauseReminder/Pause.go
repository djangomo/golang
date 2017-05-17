package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time for a Pause")

	getSystemTime()
}


func getSystemTime(){
	var t = time.Now()
	p := fmt.Println
	p(t.Format(time.RFC3339))

}
