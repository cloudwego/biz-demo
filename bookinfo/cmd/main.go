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
	goflag "flag"
	"fmt"

	"github.com/cloudwego/biz-demo/bookinfo/cmd/details"
	"github.com/cloudwego/biz-demo/bookinfo/cmd/productpage"
	"github.com/cloudwego/biz-demo/bookinfo/cmd/ratings"
	"github.com/cloudwego/biz-demo/bookinfo/cmd/reviews"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/configparser"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/version"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/cobra"
)

// Runnable server interface
type Runnable interface {
	Run(ctx context.Context)
}

var (
	moduleName = version.Get().Module
	rootCmd    = &cobra.Command{
		Use:   moduleName,
		Short: fmt.Sprintf("%s module", moduleName),
	}
)

func main() {
	rootCmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
	rootCmd.AddCommand(
		productpage.NewCommand(),
		reviews.NewCommand(),
		ratings.NewCommand(),
		details.NewCommand(),
	)

	configparser.Flags(rootCmd.PersistentFlags())
	if err := rootCmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
