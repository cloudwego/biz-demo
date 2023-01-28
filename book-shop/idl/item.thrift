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
namespace go cwg.bookshop.item

enum Status {
    Online // 上架
    Offline // 下架
    Delete // 删除
}

struct BookProperty {
    1: string isbn // ISBN
    2: string spu_name // 书名
    3: i64 spu_price // 定价
}

struct Product {
    1: i64 product_id
    2: string name // 商品名
    3: string pic // 主图
    4: string description // 详情
    5: BookProperty property // 属性
    6: i64 price // 价格
    7: i64 stock // 库存
    8: Status status // 商品状态
}

struct AddReq {
    1: required string name // 商品名
    2: required string pic // 主图
    3: required string description // 详情
    4: required BookProperty property // 属性
    5: required i64 price // 价格
    6: required i64 stock // 库存
}

struct AddResp {
    1: i64 product_id
    255: base.BaseResp BaseResp
}

struct EditReq {
    1: required i64 product_id
    2: optional string name // 商品名
    3: optional string pic // 主图
    4: optional string description // 详情
    5: optional BookProperty property // 属性
    6: optional i64 price // 价格
    7: optional i64 stock // 库存
}

struct EditResp {
    255: base.BaseResp BaseResp
}

struct DeleteReq {
    1: required i64 product_id
}

struct DeleteResp {
    255: base.BaseResp BaseResp
}

struct OnlineReq {
    1: required i64 product_id
}

struct OnlineResp {
    255: base.BaseResp BaseResp
}

struct OfflineReq {
    1: required i64 product_id
}

struct OfflineResp {
    255: base.BaseResp BaseResp
}

struct GetReq {
    1: required i64 product_id
}

struct GetResp {
    1: Product product
    255: base.BaseResp BaseResp
}

struct MGet2CReq {
    1: required list<i64> product_ids
}

struct MGet2CResp {
    1: map<i64, Product> product_map
    255: base.BaseResp BaseResp
}

struct SearchReq {
    1: optional string name
    2: optional string description
    3: optional string spu_name
}

struct SearchResp {
    1: list<Product> products
    255: base.BaseResp BaseResp
}

struct ListReq {
    1: optional string name
    2: optional string spu_name
    3: optional Status status
}

struct ListResp {
    1: list<Product> products
    255: base.BaseResp BaseResp
}

struct DecrStockReq {
    1: required i64 product_id
    2: required i64 stock_num
}

struct DecrStockResp {
    255: base.BaseResp BaseResp
}

service ItemService {
    AddResp Add(1: AddReq req) // 添加商品
    EditResp Edit(1: EditReq req) // 编辑商品
    DeleteResp Delete(1: DeleteReq req) // 删除商品
    OnlineResp Online(1: OnlineReq req) // 上架商品
    OfflineResp Offline(1: OfflineReq req) // 下架商品
    GetResp Get(1: GetReq req) // 查询商品 2B
    MGet2CResp MGet2C(1: MGet2CReq req) // 批量查询商品 2C
    SearchResp Search(1: SearchReq req) // 搜索商品 c端
    ListResp List(1: ListReq req) // 商品列表 b端
    DecrStockResp DecrStock(1: DecrStockReq req) // 扣减库存
    DecrStockResp DecrStockRevert(1: DecrStockReq req) // 库存返还
}



