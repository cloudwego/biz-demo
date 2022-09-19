# Bookinfo
> Rewrite **[Bookinfo](https://istio.io/latest/en/docs/examples/bookinfo/)** project using `hertz`, `kitex`

## 架构 
![img.png](./docs/bookinfo-arch.png)


### 流量路由示例

### 配置按比例路由
```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews
spec:
  hosts:
    - reviews
  http:
  - route:
    - destination:
        host: reviews
        subset: v1
      weight: 80
    - destination:
        host: reviews
        subset: v2
      weight: 20

```

### 按比例路由生效
![weight](./docs/weight-routing.png)

### 配置灰度流量路由规则

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
```

### 携带 baggage 发起请求
```shell
curl -H "baggage: env=dev" http://localhost/api/v1/products/1
```

### 路由已经生效
![img_1.png](./docs/canary-routing.png)
