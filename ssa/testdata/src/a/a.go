package a

import "fmt"

func f() {
	b1 := IsApple("orange")
	b1 = IsOdd(3)
	if b1 {
		fmt.Println("Odd number")
	} else {
		fmt.Println("Even number")
	}
}

func g() {
	if b1 := IsOdd(1); b1 {
		fmt.Println("Odd number")
	} else {
		fmt.Println("Even number")
	}
}

func IsOdd(num int) bool {
	return num%2 != 0
}

func IsApple(str string) bool {
	return str == "apple"
}
