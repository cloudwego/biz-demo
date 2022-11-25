package handler

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/biz-demo/open-payment-platform/hertz-gateway/biz/errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type requiredParams struct {
	Method     string `form:"method,required" json:"method,required"`
	MerchantId string `form:"merchant_id,required" json:"merchant_id,required"`
	BizParams  string `form:"biz_params,required" json:"biz_params,required"`
}

var (
	SvcMap = make(map[string]genericclient.Client)
)

func Gateway(ctx context.Context, c *app.RequestContext) {
	svcName := c.Param("svc")
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, errors.New(errors.ErrCodeServerNotFound))
		return
	}
	var params requiredParams
	if err := c.BindAndValidate(&params); err != nil {
		hlog.Error(err)
		c.JSON(http.StatusOK, errors.New(errors.ErrCodeServerMethodNotFound))
		return
	}

	req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(params.BizParams)))
	if err != nil {
		hlog.Fatalf("new http request failed: %v", err)
	}
	req.URL.Path = fmt.Sprintf("/%s/%s", svcName, params.Method)

	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		hlog.Errorf("convert request failed: %v", err)
		c.JSON(http.StatusOK, errors.New(errors.ErrCodeServerHandleFail))
		return
	}
	resp, err := cli.GenericCall(ctx, "", customReq)
	respMap := make(map[string]interface{})
	if err != nil {
		hlog.Errorf("GenericCall err:%v", err)
		bizErr, ok := kerrors.FromBizStatusError(err)
		if !ok {
			c.JSON(http.StatusOK, errors.New(errors.ErrCodeServerHandleFail))
			return
		}
		respMap["err_code"] = bizErr.BizStatusCode()
		respMap["err_message"] = bizErr.BizMessage()
		c.JSON(http.StatusOK, respMap)
		return
	}
	realResp, ok := resp.(*generic.HTTPResponse)
	if !ok {
		c.JSON(http.StatusOK, errors.New(errors.ErrCodeServerHandleFail))
		return
	}
	realResp.Body["err_code"] = 0
	realResp.Body["err_message"] = "ok"
	c.JSON(http.StatusOK, realResp.Body)
}
