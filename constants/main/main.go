package main

import "fmt"

func main()  {

	//iota

	const (
	 _  = iota
	 KB = 1 << (iota * 10)
	 MB = 1 << (iota * 10)
	 GB = 1 << (iota * 10)
	)

	fmt.Printf("%d - ", KB)
	fmt.Printf("%b - ", KB)
	fmt.Printf("%d - ", MB)
	fmt.Printf("%b - ", MB)
	fmt.Printf("%d - ", GB)
	fmt.Printf("%b - ", GB)


}
