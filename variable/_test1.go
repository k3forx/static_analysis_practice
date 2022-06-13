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
