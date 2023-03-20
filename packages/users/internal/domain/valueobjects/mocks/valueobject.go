package mocks

import (
	"github.com/mateusmacedo/users/internal/domain/valueobjects"
)

type MockedValueObject struct {
	value string
}

func NewMockedValueObject(value string) valueobjects.ValueObject {
	return MockedValueObject{value: value}
}

func (v MockedValueObject) Value() interface{} {
	return v.value
}

func (v MockedValueObject) String() string {
	return v.value
}

func (v MockedValueObject) Equals(other valueobjects.ValueObject) bool {
	return true
}
