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

package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cloudwego/biz-demo/open-payment-platform/hertz-gateway/biz/errors"
	"github.com/cloudwego/biz-demo/open-payment-platform/hertz-gateway/biz/types"
	"github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/common"
	"github.com/cloudwego/biz-demo/open-payment-platform/pkg/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// AuthParam authentication required parameters
type AuthParam struct {
	Sign       string `form:"sign,required" json:"sign"`
	SignType   string `form:"sign_type,required" json:"sign_type"`
	MerchantId string `form:"merchant_id,required" json:"merchant_id"`
	NonceStr   string `form:"nonce_str,required" json:"nonce_str"`
}

// GatewayAuth is the middleware for authentication
func GatewayAuth() []app.HandlerFunc {
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {
		var authParam AuthParam

		// verify
		if err := c.BindAndValidate(&authParam); err != nil {
			hlog.Error(err)
			c.JSON(http.StatusOK, errors.New(common.Err_BadRequest))
			c.Abort()
			return
		}
		// TODO get key in the right way
		// actual business, the key should be different depending on merchant
		key := "123"
		p, err := auth.NewSignProvider(authParam.SignType, key)
		if err != nil {
			hlog.Error(err)
			c.JSON(http.StatusOK, errors.New(common.Err_Unauthorized))
			c.Abort()
			return
		}

		if !p.Verify(authParam.Sign, map[string]interface{}{}) {
			hlog.Error(err)
			c.JSON(http.StatusOK, errors.New(common.Err_Unauthorized))
			c.Abort()
			return
		}

		c.Next(ctx)

		// build signature
		data := make(utils.H)
		if err = json.Unmarshal(c.Response.Body(), &data); err != nil {
			dataJson, _ := json.Marshal(errors.New(common.Err_RequestServerFail))
			c.Response.SetBody(dataJson)
			return
		}
		data[types.ResponseNonceStr] = authParam.NonceStr
		data[types.ResponseSignType] = authParam.SignType
		data[types.ResponseSign] = p.Sign(data)
		dataJson, _ := json.Marshal(data)
		c.Response.SetBody(dataJson)
	}}
}
