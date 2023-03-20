package valueobjects_test

import (
	"testing"

	"github.com/mateusmacedo/users/internal/domain/valueobjects"
	"github.com/mateusmacedo/users/internal/domain/valueobjects/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewEmailAddress(t *testing.T) {
	// Test valid email address
	email, err := valueobjects.NewEmailAddress("john.doe@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, email)
	assert.Equal(t, "john.doe@example.com", email.Value().(string))
	assert.Equal(t, "john.doe@example.com", email.String())

	// Test empty email address
	email, err = valueobjects.NewEmailAddress("")
	assert.Error(t, err)
	assert.Nil(t, email)

	// Test invalid email address format
	email, err = valueobjects.NewEmailAddress("invalid-email-format")
	assert.Error(t, err)
	assert.Nil(t, email)

	// Test invalid email address format
	email, err = valueobjects.NewEmailAddress("invalid-email-format@")
	assert.Error(t, err)
	assert.Nil(t, email)

	// Test invalid email address format
	email, err = valueobjects.NewEmailAddress("invalid-email-format@invalid-email-format.")
	assert.Error(t, err)
	assert.Nil(t, email)
}

func TestEmailAddress_Equals(t *testing.T) {
	// Test equal email addresses
	email1, _ := valueobjects.NewEmailAddress("john.doe@example.com")
	email2, _ := valueobjects.NewEmailAddress("JOHN.DOE@example.com")
	assert.True(t, email1.Equals(email2))

	// Test different email addresses
	email3, _ := valueobjects.NewEmailAddress("jane.doe@example.com")
	assert.False(t, email1.Equals(email3))

	// Test different value object
	anotherValueObject := mocks.NewMockedValueObject("dummy")
	assert.False(t, email1.Equals(anotherValueObject))
}
