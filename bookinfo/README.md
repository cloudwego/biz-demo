# Bookinfo

> Rewrite **[Bookinfo](https://istio.io/latest/en/docs/examples/bookinfo/)** project using `hertz`, `kitex`

## Architecture
![img.png](./docs/bookinfo-arch.png)

## Lane
![lane.png](./docs/lane.png)

## Traffic routing example

#### Define routing rules

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
            subset: v1
          weight: 80
        - destination:
            host: reviews
            subset: v3
          weight: 20

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
            subset: v1
          weight: 100
```

#### Request base lane service, rating `0` or `1` randomly
![bookinfo_1.png](docs/bookinfo_rating_1.png)
![bookinfo_2.png](docs/bookinfo_without_rating.png)

#### Set the request coloring flag through the browser mod-header plugin
![bookinfo_3.png](docs/bookinfo_header.png)

#### Click the refresh button again, you can find that the request hits the branch lane, and the rating becomes `5`
![bookinfo_4.png](docs/bookinfo_canary.png)


### View Tracing
![tracing-topo](docs/coa-tracing-topo.png)
![tracing](docs/coa-tracing.png)

### View Topology
![operation-topo](docs/upstream-operation-topo.png)
