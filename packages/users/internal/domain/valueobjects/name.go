package valueobjects

import "strings"

type Name struct {
	name string
}

func NewName(name string) *Name {
	return &Name{
		name: strings.TrimSpace(name),
	}
}

func (p *Name) Equals(other ValueObject) bool {
	if otherName, ok := other.(*Name); ok {
		return p.name == otherName.name
	}
	return false
}

func (p *Name) Value() interface{} {
	return p.name
}

func (p *Name) String() string {
	return p.name
}
