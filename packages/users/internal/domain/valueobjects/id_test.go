package valueobjects_test

import (
	"testing"

	"github.com/mateusmacedo/users/internal/domain/valueobjects"
	"github.com/mateusmacedo/users/internal/domain/valueobjects/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	// Test creating new ID
	id := valueobjects.NewID("123")
	assert.NotNil(t, id)
	assert.Equal(t, "123", id.Value().(string))
	assert.Equal(t, "123", id.String())
}

func TestID_Equals(t *testing.T) {
	// Test equal IDs
	id1 := valueobjects.NewID("123")
	id2 := valueobjects.NewID("123")
	assert.True(t, id1.Equals(id2))

	// Test different IDs
	id3 := valueobjects.NewID("456")
	assert.False(t, id1.Equals(id3))

	// Test different value object
	anotherValueObject := mocks.NewMockedValueObject("123")
	assert.False(t, id1.Equals(anotherValueObject))
}
