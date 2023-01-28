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
namespace go cwg.bookshop.order

enum Status {
    Finish
    Cancel
    Pending
}

struct OrderItem {
    1: i64 order_id
    2: i64 user_id
    3: string user_name
    4: string address
    5: i64 product_id
    6: i64 stock_num
    7: string product_snapshot
    8: Status status
    9: i64 create_time
    10: i64 update_time
}
struct CreateOrderReq {
    1: required i64 user_id
    2: required string address
    3: required i64 product_id
    4: required i64 stock_num
}

struct CreateOrderResp {
    255: base.BaseResp BaseResp
}

struct CancelOrderReq {
    1: required i64 order_id
}

struct CancelOrderResp {
    255: base.BaseResp BaseResp
}

struct ListOrderReq {
    1: required i64 user_id
    2: optional Status status
}

struct ListOrderResp {
    1: list<OrderItem> orders
    255: base.BaseResp BaseResp
}

struct GetOrderByIdReq {
    1: required i64 order_id
}

struct GetOrderByIdResp {
    1: OrderItem order
    255: base.BaseResp BaseResp
}

service OrderService {
    CreateOrderResp CreateOrder(1: CreateOrderReq req) // 创建订单
    CancelOrderResp CancelOrder(1: CancelOrderReq req) // 取消订单
    ListOrderResp ListOrder(1: ListOrderReq req) // 订单列表
    GetOrderByIdResp GetOrderById(1: GetOrderByIdReq req) // 订单详情
}
