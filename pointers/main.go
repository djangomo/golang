package main

import "fmt"

func zero(z *int) {
	*z = 10
}

func main() {
	x := 5
	//pass over the ram-address and change it back with pointers
	zero(&x)
	fmt.Println(x)

}
