package valueobjects_test

import (
	"testing"

	"github.com/mateusmacedo/users/internal/domain/valueobjects"
	"github.com/mateusmacedo/users/internal/domain/valueobjects/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	// Test valid password
	pass, err := valueobjects.NewPassword("Passw0rd!")
	assert.NoError(t, err)
	assert.NotNil(t, pass)
	assert.Equal(t, "Passw0rd!", pass.Value().(string))
	assert.Equal(t, "Passw0rd!", pass.String())

	// Test password too short
	pass, err = valueobjects.NewPassword("123")
	assert.Error(t, err)
	assert.Nil(t, pass)

	// Test password without number
	pass, err = valueobjects.NewPassword("Password!")
	assert.Error(t, err)
	assert.Nil(t, pass)

	// Test password without uppercase letter
	pass, err = valueobjects.NewPassword("password1!")
	assert.Error(t, err)
	assert.Nil(t, pass)

	// Test password without special character
	pass, err = valueobjects.NewPassword("Password1")
	assert.Error(t, err)
	assert.Nil(t, pass)
}

func TestPassword_Equals(t *testing.T) {
	// Test equal passwords
	pass1, _ := valueobjects.NewPassword("Passw0rd!")
	pass2, _ := valueobjects.NewPassword("Passw0rd!")
	assert.True(t, pass1.Equals(pass2))

	// Test different passwords
	pass3, _ := valueobjects.NewPassword("AnotherPassw0rd!")
	assert.False(t, pass1.Equals(pass3))

	// Test different value object
	anotherValueObject := mocks.NewMockedValueObject("Passw0rd!")
	assert.False(t, pass1.Equals(anotherValueObject))
}
