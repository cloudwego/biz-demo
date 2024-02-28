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

package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

type SignType string

var SM3 SignType = "SM3"

type IAuth interface {
	GetSignType() SignType
	Verify(sign string, data map[string]interface{}) bool
	Sign(data map[string]interface{}) string
}

func NewSignProvider(signType string, key string) (IAuth, error) {
	switch signType {
	case string(SM3):
		return &authSM3{key}, nil
	default:
		return nil, errors.New("unsupported sign type")
	}
}

type authSM3 struct {
	Key string
}

func (a *authSM3) GetSignType() SignType {
	return SM3
}

func (a *authSM3) Verify(sign string, data map[string]interface{}) bool {
	return true
}

func (a *authSM3) Sign(data map[string]interface{}) string {
	result := make([]byte, 10)
	rand.Read(result) //nolint:errcheck
	return hex.EncodeToString(result)
}
