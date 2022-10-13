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
Rewrite **[Bookinfo](https://istio.io/latest/docs/examples/bookinfo/)** project using `hertz`, `kitex`, same as the **[Bookinfo](https://istio.io/latest/docs/examples/bookinfo/)**

> The application displays information about a book, similar to a single catalog entry of an online book store. Displayed on the page is a description of the book, book details (ISBN, number of pages, and so on), and a few book reviews.

The Bookinfo application is broken into four separate microservices:

- **productpage**. The productpage microservice calls the details and reviews **microservices** to populate the page.
- **details**. The details microservice contains book information.
- **reviews**. The reviews microservice contains book reviews. It also calls the ratings microservice.
- **ratings**. The ratings microservice contains book ranking information that accompanies a book review.

##### What is the business logic?

- [x] Use `wire` for dependency injection
- [x] Use `opentelemetry` for tracing
- [x] Implement proxyless `flow lane` with [`Kitex-xds`](https://github.com/kitex-contrib/xds) and `opentelemetry baggage`
- [x] Implement a bookinfo ui using arco-design react

##### Which CloudWeGo subprojects are used? List all technologies used.
- [kitex](https://github.com/cloudwego/kitex)
- [hertz](https://github.com/cloudwego/hertz)
- [kitex-xds](https://github.com/kitex-contrib/xds)
- [kitex-opentelemetry](https://github.com/kitex-contrib/obs-opentelemetry)
- [hertz-opentelemetry](https://github.com/hertz-contrib/obs-opentelemetry)

##### Detailed documentation
[bookinfo](./bookinfo/README.md)

#### Contributors & Maintainers
[@CoderPoet](https://github.com/CoderPoet)

