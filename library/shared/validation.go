package shared

import (
	"regexp"

	"github.com/google/uuid"
)

const (
	numericOnly       = "^[0-9]*$"
	alphabetWithSpace = "^[a-zA-Z ]*$"
	email             = "^[A-Z0-9._%+-]+@[A-Z0-9.-]+\\.[A-Z]{2,6}$"
)

func ValidAlphabet(param string) bool {
	regex, _ := regexp.Compile(alphabetWithSpace)
	isMatch := regex.MatchString(param)

	if isMatch {
		return true
	} else {
		return false
	}
}

func ValidNumeric(param string) bool {
	regex, _ := regexp.Compile(numericOnly)
	isMatch := regex.MatchString(param)

	if isMatch {
		return true
	} else {
		return false
	}
}

func ValidEmail(param string) bool {
	regex, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	isMatch := regex.MatchString(param)

	if isMatch {
		return true
	} else {
		return false
	}
}

func ValidateUUID(id string) bool {
	_, err := uuid.Parse(id)
	if err != nil {
		return false
	} else {
		return true
	}
}
