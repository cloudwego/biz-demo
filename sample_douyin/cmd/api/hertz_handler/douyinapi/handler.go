// Copyright 2023 CloudWeGo Authors
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

package douyinapi

import (
	"reflect"

	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, resp interface{}) {
	Err := errno.ConvertErr(err)
	// r.SetErr(Err)
	if resp == nil {
		c.JSON(consts.StatusOK, Response{
			StatusCode: Err.ErrCode,
			StatusMsg:  Err.ErrMsg,
		})
		return
	}
	ref_r := reflect.ValueOf(resp)
	// fmt.Println(ref_r.Elem().FieldByName("StatusCode").CanSet(), ref_r.Elem().FieldByName("StatusMsg").CanSet(), Err)
	if !ref_r.Elem().FieldByName("StatusCode").CanSet() || !ref_r.Elem().FieldByName("StatusMsg").CanSet() {
		c.JSON(consts.StatusOK, Response{
			StatusCode: Err.ErrCode,
			StatusMsg:  Err.ErrMsg,
		})
		return
	}

	ref_r.Elem().FieldByName("StatusCode").SetInt(Err.ErrCode)
	ref_r.Elem().FieldByName("StatusMsg").SetString(Err.ErrMsg)

	c.JSON(consts.StatusOK, resp)
	// r.Send(c)
}
