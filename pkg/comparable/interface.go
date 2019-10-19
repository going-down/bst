package comparable

type Interface interface {
	Less(other Interface) bool
	Greater(other Interface) bool
	Equal(other Interface) bool
}
