# Hertz Gateway

## Principle

- Traverse `idl` directory and build `generic client`,when gateway app run.(see [router.go](router.go))
- All requests of router group `/gateway` must be verified authentication using GatewayAuth.(
  see [gateway_auth.go](biz/middleware/gateway_auth.go))
- Accept the requests with the query path of prefix `/gateway`.
- `Gateway` handler process these requests and route them to backend server by `generic call` client.(
  see [gateway.go](biz/handler/gateway.go))

## References

- [Generic Call](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/)
- [IDL Definition Specification for Mapping between Thrift and HTTP](https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/thrift_idl_annotation_standards/)
