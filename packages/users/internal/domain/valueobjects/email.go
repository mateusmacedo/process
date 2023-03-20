package valueobjects

import (
	"errors"
	"strings"
	"unicode/utf8"
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
	if utf8.RuneCountInString(email) > 254 {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	if len(parts[0]) > 64 {
		return false
	}
	if len(parts[1]) > 255 || len(parts[1]) < 3 {
		return false
	}
	return true
}
