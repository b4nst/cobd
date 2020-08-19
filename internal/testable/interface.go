package testable

type Testable interface {
	Test() error
	Error() error
	Name() string
}
