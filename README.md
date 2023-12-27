# gomall

## 演示步骤

### 环境准备

`make env-start`

### 启动服务

`make run svc=user`

`make run svc=product`

`make run svc=frontend`

`make run svc=cart`

`make run svc=checkout`

`make run svc=order`

`make run svc=payment`

### 检查服务
`make open.consul`

### 默认数据
- 分类 2 个
- 产品 多个
- 账号 1 个

### 业务流程演示

1. 注册
1. 登录
1. 分类浏览
1. 产品浏览
1. 加购
1. 购物车商品数量角标
1. 下单
1. 订单中心
1. 退出

### 可观测性
