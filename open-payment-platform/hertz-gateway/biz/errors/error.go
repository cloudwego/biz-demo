// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
