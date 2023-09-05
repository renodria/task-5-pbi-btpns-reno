package helpers

import (
	"github.com/asaskevich/govalidator"
)

func IsEmail(email string) bool {
	if !govalidator.IsEmail(email) {
		return false
	} else {
		return true
	}
}

func MinlengthPassword(password string) bool {
	if !govalidator.StringLength(password, "6", "15") {
		return false
	} else {
		return true
	}
}

func Required(value string) bool {
	if value == "" {
		return false
	} else {
		return true
	}
}
