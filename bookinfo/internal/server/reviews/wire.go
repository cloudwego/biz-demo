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

func NewServer(ctx context.Context) (*Server, error) {
	panic(wire.Build(
		configparser.Default,
		Configure,

		reviews.New,

		injectors.ProvideRatingsClient,

		wire.FieldsOf(new(*Options),
			"Server",
			"Ratings",
		),
		wire.Struct(new(Server), "*"),
	))
}
