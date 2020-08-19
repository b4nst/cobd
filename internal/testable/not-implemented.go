package testable

type NotImplemented struct {
	err error
}

func (ni *NotImplemented) Error() error {
	return ni.err
}

func (ni *NotImplemented) Name() string {
	return "Not implemented"
}

func (ni *NotImplemented) Test() error {
	return ni.Error()
}
