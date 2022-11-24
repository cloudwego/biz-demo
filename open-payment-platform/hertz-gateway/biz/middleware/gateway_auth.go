package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cloudwego/biz-demo/open-payment-platform/hertz-gateway/biz/errors"
	"github.com/cloudwego/biz-demo/open-payment-platform/pkg/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type AuthParam struct {
	Sign       string `form:"sign,required" json:"sign,required"`
	SignType   string `form:"sign_type,required" json:"sign_type,required"`
	MerchantId string `form:"merchant_id" json:"merchant_id,required"`
	NonceStr   string `form:"nonce_str,required" json:"nonce_str,required"`
}

func GatewayAuth() []app.HandlerFunc {
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {
		var authParam AuthParam

		if err := c.BindAndValidate(&authParam); err != nil {
			hlog.Error(err)
			c.JSON(http.StatusOK, errors.New(errors.ErrCodeBadRequest))
			c.Abort()
			return
		}
		// TODO get key in the right way
		key := "123"
		p, err := auth.NewSignProvider(authParam.SignType, key)
		if err != nil {
			hlog.Error(err)
			c.JSON(http.StatusOK, errors.New(errors.ErrCodeUnauthorized))
			c.Abort()
			return
		}

		if !p.Verify(authParam.Sign, map[string]interface{}{}) {
			hlog.Error(err)
			c.JSON(http.StatusOK, errors.New(errors.ErrCodeUnauthorized))
			c.Abort()
			return
		}

		c.Next(ctx)

		var data = make(utils.H)
		if err = json.Unmarshal(c.Response.Body(), &data); err != nil {
			dataJson, _ := json.Marshal(errors.ErrCodeResponseUnableParse)
			c.Response.SetBody(dataJson)
			return
		}
		data["nonce_str"] = authParam.NonceStr
		data["sign_type"] = authParam.SignType
		data["sign"] = p.Sign(data)
		dataJson, _ := json.Marshal(data)
		c.Response.SetBody(dataJson)
	}}
}
