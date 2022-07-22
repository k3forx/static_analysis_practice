package a

import (
	"errors"
)

func f() error {
	if _, err := errFunc(); err != nil {
		if errCheck(err) {
			return err
		}
	}
	return nil
}

func errFunc() (int, error) {
	return 1, errors.New("error")
}

func errCheck(err error) bool {
	return err.Error() == "error"
}
