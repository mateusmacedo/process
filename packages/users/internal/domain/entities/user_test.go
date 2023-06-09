package entities_test

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/mateusmacedo/users/internal/domain/entities"
	"github.com/mateusmacedo/users/internal/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateNew(t *testing.T) {
	wayback := time.Date(2023, time.May, 12, 1, 1, 1, 4, time.UTC)
	patch := monkey.Patch(time.Now, func() time.Time { return wayback })
	defer patch.Unpatch()

	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@email.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	mockTime := time.Now()

	user, err := entities.CreateNew(name, email, password)

	assert.NoError(t, err)
	assert.Empty(t, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, password, user.Password)
	assert.Empty(t, user.Groups)
	assert.Empty(t, user.Roles)
	assert.Empty(t, user.Token)
	assert.Nil(t, user.LastActive)
	assert.Equal(t, mockTime, user.UpdatedAt)
	assert.Nil(t, user.DeletedAt)
	assert.Equal(t, 0, user.Version)
	assert.Empty(t, user.VersionBy)
}

func TestNewUser(t *testing.T) {
	id := valueobjects.NewID("1234567890")
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@email.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	groups := []string{"admin", "users"}
	roles := []string{"admin"}
	token := "abc123"
	lastActive := time.Now()
	createdAt := time.Now().Add(-24 * time.Hour)
	updatedAt := time.Now().Add(-12 * time.Hour)
	deletedAt := time.Now().Add(-1 * time.Hour)
	version := 1
	versionBy := "John"

	expectedUser := &entities.User{
		ID:         id,
		Name:       name,
		Email:      email,
		Password:   password,
		Groups:     groups,
		Roles:      roles,
		Token:      token,
		LastActive: &lastActive,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  &deletedAt,
		Version:    version,
		VersionBy:  versionBy,
	}

	user, err := entities.NewUser(id, name, email, password, groups, roles, token, &lastActive, createdAt, updatedAt, &deletedAt, version, versionBy)

	require.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUser_UpdateName(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	user := &entities.User{
		Name: name,
	}

	nameToUpdate := valueobjects.NewName("John Doe Updated")
	err := user.UpdateName(nameToUpdate)

	assert.NoError(t, err)
	assert.Equal(t, nameToUpdate, user.Name)
}

func TestUser_UpdateEmail(t *testing.T) {
	emailAddress, _ := valueobjects.NewEmailAddress("johndoe@email.com")
	user := &entities.User{
		Email: emailAddress,
	}

	emailAddressToUpdate, _ := valueobjects.NewEmailAddress("johndoeupdated@email.com")
	err := user.UpdateEmail(emailAddressToUpdate)

	assert.NoError(t, err)
	assert.Equal(t, emailAddressToUpdate, user.Email)

}

func TestUser_UpdatePassword(t *testing.T) {
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Password: password,
	}

	passwordToUpdate, _ := valueobjects.NewPassword("Pass*1234")
	err := user.UpdatePassword(passwordToUpdate)

	assert.NoError(t, err)
	assert.Equal(t, passwordToUpdate, user.Password)
}

func TestDelete(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	assert.False(t, user.IsDeleted())

	err := user.Delete()

	assert.NoError(t, err)
	assert.True(t, user.IsDeleted())
	assert.NotNil(t, user.DeletedAt)
	assert.WithinDuration(t, time.Now(), *user.DeletedAt, time.Second)
}

func TestIsInGroup(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	assert.False(t, user.IsInGroup("group1"))
}

func TestAddToGroup(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	err := user.AddToGroup("group1")
	assert.NoError(t, err)
	assert.True(t, user.IsInGroup("group1"))

	err = user.AddToGroup("group1")
	assert.Error(t, err)
}

func TestRemoveFromGroup(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	err := user.AddToGroup("group1")
	assert.NoError(t, err)
	assert.True(t, user.IsInGroup("group1"))

	err = user.RemoveFromGroup("group1")
	assert.NoError(t, err)
	assert.False(t, user.IsInGroup("group1"))

	err = user.RemoveFromGroup("group1")
	assert.Error(t, err)
}

func TestAddRole(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	err := user.AddRole("admin")
	assert.NoError(t, err)
	assert.True(t, user.HasRole("admin"))

	err = user.AddRole("admin")
	assert.Error(t, err)
	assert.True(t, user.HasRole("admin"))
}

func TestRemoveRole(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	err := user.AddRole("admin")
	assert.NoError(t, err)
	assert.True(t, user.HasRole("admin"))

	err = user.RemoveRole("admin")
	assert.NoError(t, err)
	assert.False(t, user.HasRole("admin"))

	err = user.RemoveRole("admin")
	assert.Error(t, err)
}

func TestHasRole(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	err := user.AddRole("admin")
	assert.NoError(t, err)
	assert.True(t, user.HasRole("admin"))

	err = user.RemoveRole("admin")
	assert.NoError(t, err)
	assert.False(t, user.HasRole("admin"))
}

func TestStartSession(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	err := user.StartSession("token123")
	assert.NoError(t, err)
	assert.NotEmpty(t, user.Token)
	assert.True(t, user.IsActive())

	err = user.StartSession("token123")
	assert.Error(t, err)

	user.Delete()
	err = user.StartSession("token456")
	assert.Error(t, err)
}

func TestEndSession(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	err := user.StartSession("token123")
	assert.NoError(t, err)
	assert.NotEmpty(t, user.Token)
	assert.True(t, user.IsActive())

	err = user.EndSession()
	assert.NoError(t, err)
	assert.Empty(t, user.Token)
	assert.False(t, user.IsActive())

	err = user.EndSession()
	assert.Error(t, err)

	_ = user.StartSession("token456")
	user.Delete()
	err = user.EndSession()
	assert.Error(t, err)
}

func TestIsActive(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	assert.False(t, user.IsActive())

	now := time.Now()
	monkey.Patch(time.Since, func(since time.Time) time.Duration {
		return time.Minute
	})
	defer monkey.Unpatch(time.Since)

	user.LastActive = &now
	assert.True(t, user.IsActive())

	monkey.Patch(time.Since, func(since time.Time) time.Duration {
		return time.Hour
	})
	defer monkey.Unpatch(time.Since)

	assert.False(t, user.IsActive())
}

func TestIsDeleted(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	assert.False(t, user.IsDeleted())

	now := time.Now()
	user.DeletedAt = &now
	assert.True(t, user.IsDeleted())
}

func TestNextVersion(t *testing.T) {
	name := valueobjects.NewName("John Doe")
	email, _ := valueobjects.NewEmailAddress("johndoe@example.com")
	password, _ := valueobjects.NewPassword("Pass*123")
	user := &entities.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   0,
		VersionBy: "",
	}

	user.NextVersion("admin")
	assert.Equal(t, 1, user.Version)
	assert.Equal(t, "admin", user.VersionBy)

	user.NextVersion("editor")
	assert.Equal(t, 2, user.Version)
	assert.Equal(t, "editor", user.VersionBy)
}
