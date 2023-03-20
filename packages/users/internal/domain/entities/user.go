package entities

import (
	"fmt"
	"time"

	"github.com/mateusmacedo/users/internal/domain/valueobjects"
)

type UserInterface interface {
	UpdateName(name valueobjects.Name)
	UpdateEmail(email valueobjects.EmailAddress)
	UpdatePassword(password valueobjects.Password)
	Delete()
	AddToGroup(group string)
	RemoveFromGroup(group string)
	AddRole(role string)
	RemoveRole(role string)
	StartSession(token string)
	EndSession()
	IsActive() bool
	IsDeleted() bool
	IsInGroup(group string) bool
	HasRole(role string) bool
	NextVersion(versionBy string)
}

type User struct {
	ID         *valueobjects.ID           `json:"id"`
	Name       *valueobjects.Name         `json:"name"`
	Email      *valueobjects.EmailAddress `json:"email"`
	Password   *valueobjects.Password     `json:"password,omitempty"`
	Groups     []string                   `json:"groups,omitempty"`
	Roles      []string                   `json:"roles,omitempty"`
	Token      string                     `json:"-"`
	LastActive *time.Time                 `json:"last_active"`
	CreatedAt  time.Time                  `json:"created_at"`
	UpdatedAt  time.Time                  `json:"updated_at"`
	DeletedAt  *time.Time                 `json:"deleted_at,omitempty"`
	Version    int                        `json:"version"`
	VersionBy  string                     `json:"version_by"`
}

func NewUser(id *valueobjects.ID, name *valueobjects.Name, email *valueobjects.EmailAddress, password *valueobjects.Password, groups []string, roles []string, token string, lastActive *time.Time, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time, version int, versionBy string) (*User, error) {
	return &User{
		ID:         id,
		Name:       name,
		Email:      email,
		Password:   password,
		Groups:     groups,
		Roles:      roles,
		Token:      token,
		LastActive: lastActive,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  deletedAt,
		Version:    version,
		VersionBy:  versionBy,
	}, nil
}

func CreateNew(name *valueobjects.Name, email *valueobjects.EmailAddress, password *valueobjects.Password) (*User, error) {
	groups := []string{}
	roles := []string{}
	createdAt := time.Now()
	updatedAt := time.Now()
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		Groups:    groups,
		Roles:     roles,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Version:   0,
		VersionBy: "",
	}, nil
}

func (u *User) UpdateName(name *valueobjects.Name) error {
	u.Name = name
	return nil
}

func (u *User) UpdateEmail(email *valueobjects.EmailAddress) error {
	u.Email = email
	return nil
}

func (u *User) UpdatePassword(password *valueobjects.Password) error {
	u.Password = password
	return nil
}

func (u *User) Delete() error {
	now := time.Now()
	u.DeletedAt = &now
	return nil
}

func (u *User) IsInGroup(group string) bool {
	for _, g := range u.Groups {
		if g == group {
			return true
		}
	}
	return false
}

func (u *User) AddToGroup(group string) error {
	if u.IsInGroup(group) {
		return fmt.Errorf("user %s already in group %s", u.ID, group)
	}
	u.Groups = append(u.Groups, group)
	return nil
}

func (u *User) RemoveFromGroup(group string) error {
	if !u.IsInGroup(group) {
		return fmt.Errorf("user %s not in group %s", u.ID, group)
	}
	for i, g := range u.Groups {
		if g == group {
			u.Groups = append(u.Groups[:i], u.Groups[i+1:]...)
			break
		}
	}
	return nil
}

func (u *User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u *User) AddRole(role string) error {
	if u.HasRole(role) {
		return fmt.Errorf("user %s already has role %s", u.ID, role)
	}
	u.Roles = append(u.Roles, role)
	return nil
}

func (u *User) RemoveRole(role string) error {
	if !u.HasRole(role) {
		return fmt.Errorf("user %s does not have role %s", u.ID, role)
	}
	for i, r := range u.Roles {
		if r == role {
			u.Roles = append(u.Roles[:i], u.Roles[i+1:]...)
			break
		}
	}
	return nil
}

func (u *User) StartSession(token string) error {
	if u.IsDeleted() {
		return fmt.Errorf("cannot start session for deleted user %s", u.ID)
	}

	if u.IsActive() {
		return fmt.Errorf("cannot start session for active user %s", u.ID)
	}

	u.Token = token
	now := time.Now()
	u.LastActive = &now

	return nil
}

func (u *User) EndSession() error {
	if u.IsDeleted() {
		return fmt.Errorf("cannot end session for deleted user %s", u.ID)
	}

	if !u.IsActive() {
		return fmt.Errorf("cannot end session for inactive user %s", u.ID)
	}

	u.Token = ""
	u.LastActive = nil
	return nil
}

func (u *User) IsActive() bool {
	if u.LastActive == nil {
		return false
	}

	return time.Since(*u.LastActive) < time.Hour
}

func (u *User) IsDeleted() bool {
	return u.DeletedAt != nil
}

func (u *User) NextVersion(versionBy string) {
	u.Version++
	u.VersionBy = versionBy
}
