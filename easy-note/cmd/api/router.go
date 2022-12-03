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

package main

import (
	"context"

	"easy-note/cmd/api/router"
	"easy-note/cmd/api/router/note"
	"easy-note/cmd/api/router/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// register registers all routers.
func register(r *server.Hertz) {
	user.Register(r)
	note.Register(r)
	r.POST("/v2/user/login", router.JwtMiddleware.LoginHandler)
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 404
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 405
		c.String(consts.StatusOK, "no method")
	})
}
