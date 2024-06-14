package factory

import (
	"errors"
	"fmt"
)

var (
	ErrCancel 	= errors.New("error because cancel by client or customer")
	ErrDeadline = errors.New("error deadline")
	ErrSteal	= errors.New("error item has steal by courier")
	ErrOverLoad = errors.New("item has overload")
)

type CourierError struct {
	Err     error
	Message string
}

func (e *CourierError) Error() string {
	return fmt.Sprintf("%v: %s", e.Err, e.Message)
}

func NewCourierError(err error, msg string) *CourierError {
	return &CourierError{
		Err:     err,
		Message: msg,
	}
}