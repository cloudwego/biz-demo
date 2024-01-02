# Gomall
[EN](README.md)

新人学习 CloudWeGo 的教学项目

## 技术栈
| 技术            | 介绍 |
|---------------|----|
| cwgo          | -  |
| kitex         | -  |
| [bootstrap](https://getbootstrap.com/docs/5.3/getting-started/introduction/) | Bootstrap is a powerful, feature-packed frontend toolkit. Build anything—from prototype to production—in minutes.  |
| Hertz         | -  |
| MySQL         | -  |
| Redis         | -  |
| ES            | -  |
| Prometheus    | -  |
| Jaeger        | -  |
| Docker        | -  |

## 业务逻辑
- [x] 页面访问认证检查
- [x] 注册
- [x] 登录
- [x] 退出
- [x] 产品分类
- [x] 产品
- [x] 加购
- [x] 购物车数量角标
- [x] 下单
- [x] 支付
- [x] 订单中心

## 如何使用
### 准备
必备清单
- Go
- IDE / Code Editor
- Docker
- [cwgo](https://github.com/cloudwego/cwgo)
- kitex `go install github.com/cloudwego/kitex/tool/cmd/kitex@latest`
- [Air](https://github.com/cosmtrek/air)
- ...

### 克隆项目
```
git clone ...
```

### 拷贝 `.env` 文件
```
make init
```
*Note:*`必须生成并输入 SESSION_SECRET 值供 session 功能正常使用`
### 下载 Go 依赖
```
make tidy
```

### 启动容器
```
make env-start
```
if you want to stop their docker application,you can run `make env-stop`.

### 启动某服务
该命令必须执行一个服务

*注意:* `我们使用 air 运行并热加载，必须先安装好`
```
make run svc=`svcName`
```
### 浏览 Gomall 站点
```
make open-gomall
```
### 查看注册中心
```
make open-consul
```
### Make 用法
```
make
```
## 贡献者
- [rogerogers](https://github.com/rogerogers)
- [baiyutang](https://github.com/baiyutang)
