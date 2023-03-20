package valueobjects

type ValueObject interface {
	Equals(other ValueObject) bool
	Value() interface{}
	String() string
}
