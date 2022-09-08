package reviews

import (
	"github.com/cloudwego/biz-demo/bookinfo/pkg/configparser"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/injectors"
)

type Options struct {
	Server  *ServerOptions                  `mapstructure:"server"`
	Ratings *injectors.RatingsClientOptions `mapstructure:"ratings"`
}

func DefaultOptions() *Options {
	return &Options{
		Server:  DefaultServerOptions(),
		Ratings: injectors.DefaultRatingsClientOptions(),
	}
}

func Configure(configProvider configparser.Provider) (*Options, error) {
	opt := DefaultOptions()

	cp, err := configProvider.Get()
	if err != nil {
		return nil, err
	}

	if err = cp.UnmarshalExact(opt); err != nil {
		return nil, err
	}

	return opt, nil
}
