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

package entity

type Order struct {
	ID              int
	MerchantID      string
	Channel         string
	PayWay          string
	OrderStatus     int8
	OutOrderNo      string
	TotalAmount     uint64
	Body            string
	AuthCode        string
	WxAppid         string
	SubOpenid       string
	JumpURL         string
	NotifyURL       string
	ClientIP        string
	Attach          string
	OrderExpiration string
	ExtendParams    string
}

func NewOrder() *Order {
	return &Order{}
}
