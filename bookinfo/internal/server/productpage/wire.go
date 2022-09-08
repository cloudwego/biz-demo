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

//go:build wireinject
// +build wireinject

package productpage

import (
	"context"

	"github.com/cloudwego/biz-demo/bookinfo/internal/handler/productpage"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/configparser"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/injectors"
	"github.com/google/wire"
)

func NewServer(ctx context.Context) (*Server, error) {
	panic(wire.Build(
		configparser.Default,
		Configure,

		productpage.New,

		injectors.ProvideReviewClient,
		injectors.ProvideDetailsClient,

		wire.FieldsOf(new(*Options),
			"Server",
			"Reviews",
			"Details",
		),
		wire.Struct(new(Server), "*"),
	))
}
