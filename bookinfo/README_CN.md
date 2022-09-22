# Bookinfo

> Rewrite **[Bookinfo](https://istio.io/latest/en/docs/examples/bookinfo/)** project using `hertz`, `kitex`

## 架构

![img.png](./docs/bookinfo-arch.png)

## 泳道示意图
![lane.png](./docs/lane.png)

## 流量路由示例

#### 定义路由规则

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews
spec:
  hosts:
    - reviews
  http:
    - match:
        - headers:
            baggage:
              exact: "env=dev"
      route:
        - destination:
            host: reviews
            subset: v2
          weight: 100
    - route:
        - destination:
            host: reviews
          weight: 100

---

apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: ratings
spec:
  hosts:
    - ratings
  http:
    - match:
        - headers:
            baggage:
              exact: "env=dev"
      route:
        - destination:
            host: ratings
            subset: v2
          weight: 100
    - route:
        - destination:
            host: ratings
          weight: 100
```

#### 请求基准泳道服务，评分为 0 或 1 随机
![bookinfo_1.png](docs/bookinfo_1.png)
![bookinfo_2.png](docs/bookinfo_2.png)

#### 通过浏览器 mod-header 插件，设置灰度标识 header
![bookinfo_3.png](docs/bookinfo_3.png)

#### 再点击刷新按钮，可以发现请求打到了分支泳道
![bookinfo_4.png](docs/bookinfo_4.png)


### 查看 Tracing 
![tracing-topo](docs/coa-tracing-topo.png)
![tracing](docs/coa-tracing.png)

### 查看拓扑
![operation-topo](docs/upstream-operation-topo.png)