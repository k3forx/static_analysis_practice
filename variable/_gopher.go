package gopher

import (
	"errors"
	"fmt"
)

func main() {
	res, err := Translate("こんにちは")
	if err != nil {
		if IsTypeError(err) {
			return
		}
	}
	fmt.Printf("res: %+v\n", res)

	// res, err = Translate("おはよう")
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("res: %+v\n", res)
}

// func Greeting() string {
// 	return "hello"
// }

// func IsJapanese(greeting string) bool {
// 	return greeting == "こんちには"
// }

func IsTypeError(err error) bool {
	return err == errors.New("unsupported type error")
}

func Translate(greeting interface{}) (string, error) {
	v, ok := greeting.(string)
	if !ok {
		return "", errors.New("unsupported type error")
	}

	switch v {
	case "こんにちは":
		return "hello", nil
	case "おはよう":
		return "good morning", nil
	}

	return "", errors.New("unsupported string error!")
}
