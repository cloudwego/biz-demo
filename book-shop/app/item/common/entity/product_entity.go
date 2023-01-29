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

import "github.com/jinzhu/copier"

type ProductEntity struct {
	ProductId   int64
	Name        string
	Pic         string
	Description string
	Property    *PropertyEntity
	Price       int64
	Stock       int64
	Status      int64
}

func (entity *ProductEntity) Clone() (*ProductEntity, error) {
	ret := &ProductEntity{}
	err := copier.Copy(ret, entity)
	return ret, err
}

type PropertyEntity struct {
	ISBN     string
	SpuName  string
	SpuPrice int64
}
