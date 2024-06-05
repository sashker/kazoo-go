package kazooapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type genericKazooError struct {
	Data      interface{} `json:"data"`
	Error     string      `json:"error"`
	Message   string      `json:"message"`
	Status    string      `json:"status"`
	Timestamp string      `json:"timestamp"`
	Version   string      `json:"version"`
	Node      string      `json:"node"`
	RequestID string      `json:"request_id"`
	AuthToken string      `json:"auth_token"`
}

type Error interface {
	// Satisfy generic error interface
	error
	// Returns the original error code
	Code() string
	// Returns the error details message.
	Message() string
	// Returns the original error if one was set.  Nil is returned if not set.
	Unwrap() error
}

// An Error wraps lower level errors with code, message and an original error.
// The underlying concrete error type may also satisfy other interfaces which
// can be to used to obtain more specific information about the error.
//
func NewError(code, message string, origErr error) Error {
	return newBaseError(code, message, origErr)
}

// SprintError returns a string of the formatted error code.
//
// Both extra and origErr are optional.  If they are included their lines
// will be added, but if they are not included their lines will be ignored.
func SprintError(code, message, extra string, origErr error) string {
	msg := fmt.Sprintf("%s: %s", code, message)
	if extra != "" {
		msg = fmt.Sprintf("%s\n\t%s", msg, extra)
	}
	if origErr != nil {
		msg = fmt.Sprintf("%s\ncaused by: %s", msg, origErr)
	}
	return msg
}

type baseError struct {
	code    string
	message string
	err     error
}

func newBaseError(code, message string, err error) *baseError {
	b := &baseError{
		code:    code,
		message: message,
		err:     err,
	}

	return b
}

//Error implemens Error interface
func (be baseError) Error() string {
	return SprintError(be.code, be.message, "", nil)
}

// String returns the string representation of the error.
// Alias for Error to satisfy the stringer interface.
func (be baseError) String() string {
	return be.Error()
}

// Code returns the short phrase depicting the classification of the error.
func (be baseError) Code() string {
	return be.code
}

// Message returns the error details message.
func (be baseError) Message() string {
	return be.message
}

//Unwrap implements Error interface
func (be baseError) Unwrap() error {
	return be.err
}

//UnmarshalKazooError returns error from Kazoo side
func UnmarshalKazooError(body io.ReadCloser) Error {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return NewError("ErrReadResponseBody", "", err)
	}

	e := &genericKazooError{}
	jsonErr := json.Unmarshal(bodyBytes, e)
	if err != nil {
		return NewError("UnmarshalError", string(bodyBytes), jsonErr)
	}

	return NewError(e.Status, e.Error, nil)
}
