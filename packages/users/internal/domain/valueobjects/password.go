package valueobjects

import (
	"errors"
	"regexp"
)

type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	if len(value) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}
	if !containsNumber(value) {
		return nil, errors.New("password must contain at least one number")
	}
	if !containsUppercase(value) {
		return nil, errors.New("password must contain at least one uppercase letter")
	}
	if !containsSpecialChar(value) {
		return nil, errors.New("password must contain at least one special character")
	}
	return &Password{
		value: value,
	}, nil
}

func (p *Password) Equals(other ValueObject) bool {
	if otherPass, ok := other.(*Password); ok {
		return p.value == otherPass.value
	}
	return false
}

func (p *Password) Value() interface{} {
	return p.value
}

func (p *Password) String() string {
	return p.value
}

func containsNumber(s string) bool {
	match, _ := regexp.MatchString("[0-9]", s)
	return match
}

func containsUppercase(s string) bool {
	match, _ := regexp.MatchString("[A-Z]", s)
	return match
}

func containsSpecialChar(s string) bool {
	match, _ := regexp.MatchString("[!@#$%^&*(),.?\":{}|<>]", s)
	return match
}
