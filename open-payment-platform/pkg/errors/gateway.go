package errors

const (
	ErrCodeAuthFail = 401
)

var (
	ErrAuthFail = &Error{
		Code: ErrCodeAuthFail,
	}
)
