package main

//Setup

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/gopherjs/gopherjs/js"
)

//for benchmarking so everything uses the same
const benchmark = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@#!$&*()_-+="

//max pw-chars
const maxPwChars = 10

//for version 1
var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@#!$&*()_-+=")

//for version 2
const upperlower = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//for version 3
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)
const lowerchars = "abcdefghijklmnopqrstuvwxyz@#!$&*()_-+="

//for version 4
var src = rand.NewSource(time.Now().UnixNano())

const upperchars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ@#!$&*()_-+="

func startingPoint() {
	var input string
	fmt.Println("Wich password do you want?")
	fmt.Println("1 - 10 digits - all chars used")
	fmt.Println("2 - 10 digits - only uppercase and lowercase")
	fmt.Println("3 - 10 digits - with lowercase and symbols")
	fmt.Println("4 - 10 digits - uppercase and symbols")
	fmt.Println("5 - exit program")
	fmt.Println("Enter your choice: ")
	fmt.Scan(&input)

	switch input {
	case "1":
		fmt.Println("Option 1 selected")
		fmt.Println(generateFullPW(maxPwChars))
		startingPoint()
		break
	case "2":
		fmt.Println("Option 2 selected")
		fmt.Println(generateUpperLowerPW(maxPwChars))
		startingPoint()
		break
	case "3":
		fmt.Println("Option 3 selected")
		fmt.Println(generateLowerChars(maxPwChars))
		startingPoint()
		break
	case "4":
		fmt.Println("Option 4 selected")
		fmt.Println(generateUpperChar(maxPwChars))
		startingPoint()
		break
	case "5":
		fmt.Println("Option 5 selected")
		break
	default:
		fmt.Println("nothing selected")
		startingPoint()
	}
}

//standard version
func generateFullPW(n int) string {

	pw := make([]rune, n)
	for i := range pw {

		pw[i] = chars[rand.Intn(len(chars))]
	}
	return string(pw)
}

//version with constants & remainder
func generateUpperLowerPW(n int) string {

	pw := make([]byte, n)
	for i := range pw {

		pw[i] = upperlower[rand.Int63()%int64(len(upperlower))]
	}
	return string(pw)
}

//version with constants and masking
func generateLowerChars(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {

		if remain == 0 {

			cache, remain = rand.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(lowerchars) {

			b[i] = lowerchars[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

//version with improved src
func generateUpperChar(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(upperchars) {
			b[i] = upperchars[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func main() {
	js.Global.Call("alert", "Hello, Javascript")
	startingPoint() }
