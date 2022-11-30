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
	"errors"
	"fmt"
)

type fileProvider struct{}

// NewFile create file provider
func NewFile() Provider {
	return &fileProvider{}
}

// Get file provider
func (fl *fileProvider) Get() (*Parser, error) {
	fileName := getConfigFlag()
	if fileName == "" {
		return nil, errors.New("config file not specified")
	}

	cp, err := NewParserFromFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error loading config file %q: %v", fileName, err)
	}

	return cp, nil
}
