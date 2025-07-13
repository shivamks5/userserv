package errs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shivamks5/userserv/model"
)

var (
	ErrNotFound     = errors.New("user not found")
	ErrInvalidField = errors.New("invalid data field")
	ErrBadRequest   = errors.New("invalid request")
)

func CheckName(name string) bool {
	return strings.TrimSpace(name) != ""
}

func CheckEmail(email string) bool {
	return strings.TrimSpace(email) != ""
}

func CheckAge(age int) bool {
	return age > 0
}

func ValidateUser(user model.User) error {
	var err error = nil
	if !CheckName(user.Name) {
		err = fmt.Errorf("%w, name is required", ErrInvalidField)
	}
	if !CheckEmail(user.Email) {
		if err == nil {
			err = fmt.Errorf("%w, email is required", ErrInvalidField)
		} else {
			err = fmt.Errorf("%w, email is required", err)
		}
	}
	if !CheckAge(user.Age) {
		if err == nil {
			err = fmt.Errorf("%w, age must be greater than 0", ErrInvalidField)
		} else {
			err = fmt.Errorf("%w, age must be greater than 0", err)
		}
	}
	return err
}
