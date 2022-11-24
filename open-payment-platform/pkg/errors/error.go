package errors

import "fmt"

type Error struct {
	Code int64
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d message = %s", e.Code, e.Msg)
}
