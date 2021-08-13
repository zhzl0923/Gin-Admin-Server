package err

import (
	"fmt"
	"sync"
)

var ErrContainer = map[int]*Error{}
var ErrMutex = &sync.Mutex{}
var ErrUnknown = NewError(100000, 500, "Internal server error")

type Error struct {
	code       int
	httpStatus int
	msg        string
	details    []string
}

func NewError(code, httpStatus int, msg string) *Error {
	err := &Error{code: code, httpStatus: httpStatus, msg: msg, details: []string{}}
	Register(err)
	return err
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() int {
	return e.code
}

func (c *Error) HttpStatus() int {
	return c.httpStatus
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	e.details = append(e.details, details...)
	return e
}

func ParseErr(e error) *Error {
	if e == nil {
		return ErrUnknown
	}

	if v, ok := e.(*Error); ok {
		if err, ok := ErrContainer[v.code]; ok {
			return err
		}
	}

	return ErrUnknown
}

func Register(err *Error) {
	ErrMutex.Lock()
	defer ErrMutex.Unlock()

	if _, ok := ErrContainer[err.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", err.Code()))
	}

	ErrContainer[err.Code()] = err
}
