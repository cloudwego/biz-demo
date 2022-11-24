package errors

const (
	ErrCodeBadRequest           int32 = 10001
	ErrCodeUnauthorized         int32 = 10002
	ErrCodeServerNotFound       int32 = 10003
	ErrCodeServerMethodNotFound int32 = 10004
	ErrCodeServerHandleFail     int32 = 10005
	ErrCodeResponseUnableParse  int32 = 10006
)

var errMsgMap = map[int32]string{
	ErrCodeBadRequest:           "BadRequest",
	ErrCodeUnauthorized:         "Unauthorized",
	ErrCodeServerNotFound:       "ServerNotFound",
	ErrCodeServerMethodNotFound: "ServerMethodNotFound",
	ErrCodeServerHandleFail:     "ServerHandleFail",
	ErrCodeResponseUnableParse:  "ResponseUnableParse",
}

type Err struct {
	ErrCode int32  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func New(errCode int32) Err {
	return Err{
		ErrCode: errCode,
		ErrMsg:  errMsgMap[errCode],
	}
}

func (e Err) Error() string {
	return e.ErrMsg
}
