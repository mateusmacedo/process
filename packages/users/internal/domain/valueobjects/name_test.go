package valueobjects_test

import (
	"testing"

	"github.com/mateusmacedo/users/internal/domain/valueobjects"
	"github.com/mateusmacedo/users/internal/domain/valueobjects/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	// Test creating new Name
	name := valueobjects.NewName(" John Doe ")
	assert.NotNil(t, name)
	assert.Equal(t, "John Doe", name.Value().(string))
	assert.Equal(t, "John Doe", name.String())
}

func TestName_Equals(t *testing.T) {
	// Test equal Names
	name1 := valueobjects.NewName("John Doe")
	name2 := valueobjects.NewName("John Doe")
	assert.True(t, name1.Equals(name2))

	// Test different Names
	name3 := valueobjects.NewName("Jane Doe")
	assert.False(t, name1.Equals(name3))

	// Test different value object
	anotherValueObject := mocks.NewMockedValueObject("John Doe")
	assert.False(t, name1.Equals(anotherValueObject))
}
