package gopher

import (
	"errors"
	"fmt"
)

func main() {
	if _, err := Translate("こんにちは"); err != nil {
		if IsTypeError(err) {
			fmt.Println(err)
		}
	}

	if _, err := Translate("こんにちは"); err != nil {
		fmt.Println(err)
	}
}

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
	}

	return "", errors.New("unsupported string error!")
}
