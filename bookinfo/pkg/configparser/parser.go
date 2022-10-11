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

package configparser

import (
	"fmt"

	"github.com/spf13/viper"
)

// Parser loads configuration.
type Parser struct {
	v *viper.Viper
}

func newViper() *viper.Viper {
	return viper.NewWithOptions()
}

// UnmarshalExact unmarshals the config into a struct, erroring if a field is nonexistent.
func (l *Parser) UnmarshalExact(intoCfg interface{}) error {
	return l.v.UnmarshalExact(intoCfg)
}

// NewParserFromFile creates a new Parser by reading the given file.
func NewParserFromFile(fileName string) (*Parser, error) {
	v := newViper()
	v.SetConfigFile(fileName)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read the file %v: %w", fileName, err)
	}
	return &Parser{v: v}, nil
}
