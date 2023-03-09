package errno

import (
	"errors"
	"fmt"
	"mydouyin/kitex_gen/douyinuser"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(int64(douyinuser.ErrCode_SuccessCode), "Success")
	ServiceErr             = NewErrNo(int64(douyinuser.ErrCode_ServiceErrCode), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int64(douyinuser.ErrCode_ParamErrCode), "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(int64(douyinuser.ErrCode_UserAlreadyExistErrCode), "User already exists")
	AuthorizationFailedErr = NewErrNo(int64(douyinuser.ErrCode_AuthorizationFailedErrCode), "Authorization failed")
	QueryErr               = NewErrNo(int64(douyinuser.ErrCode_QueryErrCode), "Query database error")
	ExceedAuthorization    = NewErrNo(10006, "Exceed your authority")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
