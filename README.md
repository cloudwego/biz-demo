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

##### What are the core technologies/projects used?

- [x] Use `istiod` as xDS server for CRD configuration and distribution
- [x] Use `wire` for dependency injection
- [x] Use `opentelemetry` for tracing
- [x] Implement proxyless `flow lane` with [`Kitex-xds`](https://github.com/kitex-contrib/xds) and `opentelemetry baggage`
- [x] Implement a bookinfo ui using `arco-design` react

##### Which CloudWeGo subprojects are used? List all technologies used.
- [Kitex](https://github.com/cloudwego/kitex)
- [Hertz](https://github.com/cloudwego/hertz)
- [kitex-xds](https://github.com/kitex-contrib/xds)
- [kitex-opentelemetry](https://github.com/kitex-contrib/obs-opentelemetry)
- [hertz-opentelemetry](https://github.com/hertz-contrib/obs-opentelemetry)

##### Detailed documentation
[bookinfo](./bookinfo/README.md)

#### Contributors & Maintainers
[@CoderPoet](https://github.com/CoderPoet)


### 2. Open Payment Platform

#### Description
##### What is it about and what problem does it solve?
- How to use kitex generic call as the http gateway?
- How to implement Clean Structure of `Go` using kitex?

##### What is the business scenario?
> The application demonstrates the usage of kitex generic call.

- We build the `generic call` client for each backend server by traversing IDL in Hertz.
- Hertz will accept the requests with the query path of prefix `/gateway` .
- `Gateway` handler will process these requests and route them to backend server by `generic call` client.
- `payment` server is just a sample server using kitex.This service design pkg with `Clean Structure`.

##### What are the core technologies/projects used?

- [x] Use `Hertz` as Gateway.
- [x] Use `Kitex` generic call client route requests.
- [x] Use `Kitex` as RPC framework to build micro-services.
- [x] Use `Clean Architecture` for design pkg and code layout.
- [x] Use `ent` entity framework for implementing repository.
- [x] Use `wire` for dependency injection
- [x] Use `Nacos` as service registry.
- [x] Use `MySQL` as RDBMS.

##### Which CloudWeGo subprojects are used? List all technologies used.
- [Kitex](https://github.com/cloudwego/kitex)
- [Hertz](https://github.com/cloudwego/hertz)
- kitex-layout(coming soon)

##### Detailed documentation
[Open Payment Platform](./open-payment-platform/README.md)

#### Contributors & Maintainers
[@baiyutang](https://github.com/baiyutang)

