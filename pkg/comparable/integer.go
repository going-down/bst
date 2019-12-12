package comparable

type Integer struct {
	Value int
}

func (ref *Integer) Less(other Interface) bool {
	return ref.Value < other.(*Integer).Value
}

func (ref *Integer) Greater(other Interface) bool {
	return ref.Value > other.(*Integer).Value
}

func MakeInteger(value int) Interface {
	return &Integer{Value: value}
}
