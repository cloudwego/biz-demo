# CloudWeGo Demo for Business

This repo contains CloudWeGo demos with business logic, providing valuable references for enterprise user using in production.

Each demo should/may contain multiple CloudWeGo subprojects, e.g. Kitex & Hertz, and demonstrate an individual business scenario.

## Demo List

### 1. Bookinfo

#### Description
##### What is it about and what problem does it solve?
- How to use kitex proxyless in istio?
- How to implement full-process traffic lane using CloudWeGo?

##### What is the business scenario?
Rewrite **[Bookinfo](https://istio.io/latest/en/docs/examples/bookinfo/)** project using `hertz`, `kitex`, same as the **[Bookinfo](https://istio.io/latest/en/docs/examples/bookinfo/)**
![bookinfo_4.png](bookinfo/docs/bookinfo_canary.png)

##### What is the business logic?
![lane.png](./bookinfo/docs/lane.png)

##### Which CloudWeGo subprojects are used? List all technologies used.
- [kitex](https://github.com/cloudwego/kitex)
- [hertz](https://github.com/cloudwego/hertz)
- [kitex-xds](https://github.com/kitex-contrib/xds)
- [kitex-opentelemetry](https://github.com/kitex-contrib/obs-opentelemetry)
- [hertz-opentelemetry](https://github.com/hertz-contrib/obs-opentelemetry)

##### Provide a link to detailed documentation(readme in subdirectory).
[bookinfo](./bookinfo/README.md)

#### Contributors & Maintainers
[@CoderPoet](https://github.com/CoderPoet)

