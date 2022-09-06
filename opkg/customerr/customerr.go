package customerr

type CustomErr struct {
	message string
}

func NewCustomerErr(message string) *CustomErr {
	return &CustomErr{message: message}
}

func (e CustomErr) Error() string {
	return e.message
}

var _ error = (*CustomErr)(nil)
