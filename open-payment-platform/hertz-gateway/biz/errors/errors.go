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

import (
	"github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/common"
)

type Err struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

// New Error, the error_code must be defined in IDL.
func New(errCode common.Err) Err {
	return Err{
		ErrCode: int64(errCode),
		ErrMsg:  errCode.String(),
	}
}

func (e Err) Error() string {
	return e.ErrMsg
}
