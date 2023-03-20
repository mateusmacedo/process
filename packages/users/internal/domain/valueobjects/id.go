package valueobjects

type ID struct {
	id string
}

func NewID(id string) *ID {
	return &ID{
		id: id,
	}
}

func (s *ID) Equals(other ValueObject) bool {
	if otherID, ok := other.(*ID); ok {
		return s.id == otherID.id
	}
	return false
}

func (s *ID) Value() interface{} {
	return s.id
}

func (s *ID) String() string {
	return s.id
}
