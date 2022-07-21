package a

import "fmt"

func f() {
	isOdd := "a" == "a"
	isOdd = IsOdd(3)
	if isOdd {
		fmt.Println("Odd number")
	} else {
		fmt.Println("Even number")
	}
}

func g() {
	if isOdd := IsOdd(1); isOdd {
		fmt.Println("Odd number")
	} else {
		fmt.Println("Even number")
	}
}

func IsOdd(num int) bool {
	return num%2 != 0
}
