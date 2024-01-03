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

include "base.thrift"
namespace go cwg.bookinfo.details


struct Product {
    1: required string ID,
    2: required string Title,
    3: required string Author,
    4: required string Description,
}

struct GetProductReq {
    1: required string ID,
}

struct GetProductResp {
    1: required Product Product,
}

service DetailsService {
    GetProductResp GetProduct(1: GetProductReq req)
}
