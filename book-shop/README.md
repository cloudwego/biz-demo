# Book Shop

## Introduction
An e-commerce demo built with `Kitex` and `Hertz`.

| Service Name       | Usage          | Framework    | protocol | Path       | IDL              |
|--------------------|----------------|--------------|----------|------------|------------------|
| facade             | HTTP interface | kitex/hertz  | http     | app/facade |                  |
| cwg.bookshop.user  | user service   | kitex/gorm   | thrift   | app/user   | idl/user.thrift  |
| cwg.bookshop.order | order service  | kitex/gorm   | thrift   | app/order  | idl/order.thrift |
| cwg.bookshop.item  | item service   | kitex/gorm   | thrift   | app/item   | idl/item.thrift  |

* components used
  * ElasticSearch
  * Kibana
  * MySQL
  * Redis
  * ETCD
* Hertz middlewares used
  * [swagger](http://github.com/hertz-contrib/swagger)
  * [JWT](http://github.com/hertz-contrib/jwt)
  * [pprof](https://github.com/hertz-contrib/pprof)
  * [gzip](https://github.com/hertz-contrib/gzip)

## Architecture
### Technology Architecture
![](./pics/arch.png)
### Service Calling Relations
![](./pics/relation.png)
## Quick Start

### Setup Environment
```shell
$ make start
```

### Run Services
```shell
$ make user
$ make item
$ make order
$ make facade
```

### Stop Environment
```shell
$ make stop
```

### Get Documents & Run Test
browse to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Examples
### pprof
```shell
$ go tool pprof -http=:1234 http://localhost:8080/debug/pprof/heap
```
![](./pics/pprof.png)
### User Service
#### User Register
![](./pics/register.png)
#### User Login
![](./pics/login.png)
#### Shop Login
![](./pics/shop_login.png)
#### JWT Auth
![](./pics/auth.png)
### Item Service
#### Add Item
![](./pics/item_add.png)
#### Edit Item
![](./pics/item_edit.png)
#### Delete Item
![](./pics/item_del.png)
#### Offline Item
![](./pics/item_offline.png)
#### Online Item
![](./pics/item_online.png)
#### Get Item
![](./pics/item_get.png)
#### List Items
![](./pics/item_list.png)
#### Batch Get Items (2C Interface)
![](./pics/item_2c_get.png)
#### Search Items (2C Interface)
![](./pics/item_search.png)
### Order Service
#### Create Order
![](./pics/order_create.png)
#### Cancel Order
![](./pics/order_cancel.png)
#### List Orders
![](./pics/order_list.png)
#### Get Order
![](./pics/order_get.png)
