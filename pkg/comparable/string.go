package comparable

type String struct {
	Value string
}

func (ref *String) Less(other Interface) bool {
	return ref.Value < other.(*String).Value
}

func (ref *String) Greater(other Interface) bool {
	return ref.Value > other.(*String).Value
}

func (ref *String) Equal(other Interface) bool {
	return ref.Value == other.(*String).Value
}
