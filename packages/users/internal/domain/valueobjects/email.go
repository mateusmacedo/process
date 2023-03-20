package valueobjects

import (
	"errors"
	"regexp"
	"strings"
)

type EmailAddress struct {
	value string
}

func NewEmailAddress(value string) (*EmailAddress, error) {
	if len(strings.TrimSpace(value)) == 0 {
		return nil, errors.New("email address cannot be empty")
	}
	if !isValidEmail(value) {
		return nil, errors.New("invalid email address format")
	}
	return &EmailAddress{
		value: strings.ToLower(strings.TrimSpace(value)),
	}, nil
}

func (e *EmailAddress) Equals(other ValueObject) bool {
	if otherEmail, ok := other.(*EmailAddress); ok {
		return e.value == otherEmail.value
	}
	return false
}

func (e *EmailAddress) Value() interface{} {
	return e.value
}

func (e *EmailAddress) String() string {
	return e.value
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
	return emailRegex.MatchString(email)
}
