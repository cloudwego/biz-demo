package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
	"time"
)

func main() {
	h := server.Default(server.WithDisablePrintRoute(false), server.WithAutoReloadRender(true, time.Second))
	hlog.SetLevel(hlog.LevelTrace)
	h.LoadHTMLGlob("template/*")
	h.Delims("{{", "}}")
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(http.StatusOK, "home", utils.H{
			"title": "hello cloudwego",
		})
	})
	h.Static("/static", "./")
	h.Spin()
}
