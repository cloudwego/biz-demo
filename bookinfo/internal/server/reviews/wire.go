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

package reviews

import (
	"context"

	"github.com/cloudwego/biz-demo/bookinfo/internal/service/reviews"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/configparser"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/injectors"
	"github.com/google/wire"
)

// NewServer build server with wire
func NewServer(ctx context.Context) (*Server, error) {
	panic(wire.Build(
		configparser.Default,
		Configure,

		injectors.ProvideRatingsClient,

		reviews.New,

		wire.FieldsOf(new(*Options),
			"Server",
			"Ratings",
		),
		wire.Struct(new(Server), "*"),
	))
}
