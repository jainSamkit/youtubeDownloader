package utils

import (
	"errors"
)

//ResponsePipe is a generic response container.
type ResponsePipe struct {
	Success bool
	Err     error
	Info    string
}

//SetError sets the error field
func (r *ResponsePipe) SetError(err string) {

	e := errors.New(err)
	r.Err = e
}
