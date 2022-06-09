package gopher

import "fmt"

func main() {
	// test
	greeting := Greeting()
	fmt.Println(greeting)
	fmt.Println(IsJapanese(greeting))

	greeting2 := Greeting()
	fmt.Println(greeting2)
}

func Greeting() string {
	return "hello"
}

func IsJapanese(greeting string) bool {
	return greeting == "こんちには"
}
