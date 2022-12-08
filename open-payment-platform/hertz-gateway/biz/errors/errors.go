package errors

import (
	"github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/common"
)

type Err struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func New(errCode common.Err) Err {
	return Err{
		ErrCode: int64(errCode),
		ErrMsg:  errCode.String(),
	}
}

func (e Err) Error() string {
	return e.ErrMsg
}
