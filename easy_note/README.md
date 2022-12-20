# Easy Note (dev-version)

## Introduction

Add a demo for `kitex` and `hertz` which implements a simple note service, the demo is divided into three main sections.

| Service Name | Usage                | Framework   | protocol | Path     | IDL             |
|--------------|----------------------|-------------|----------|----------|-----------------|
| demoapi      | http interface       | kitex/hertz | http     | cmd/api  |                 |
| demouser     | user data management | kitex/gorm  | thrift   | cmd/user | idl/user.thrift |
| demonote     | note data management | kitex/gorm  | thrift   | cmd/note | idl/note.thrift |

### Call Relations

```
                                    http
                           ┌────────────────────────┐
 ┌─────────────────────────┤                        ├───────────────────────────────┐
 │                         │         demoapi        │                               │
 │      ┌──────────────────►                        │◄──────────────────────┐       │
 │      │                  └───────────▲────────────┘                       │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                           resolve                                 │       │
 │      │                              │                                    │       │
req    resp                            │                                   resp    req
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                   ┌──────────▼─────────┐                          │       │
 │      │                   │                    │                          │       │
 │      │       ┌───────────►        etcd        ◄─────────────────┐        │       │
 │      │       │           │                    │                 │        │       │
 │      │       │           └────────────────────┘                 │        │       │
 │      │       │                                                  │        │       │
 │      │     register                                           register   │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
┌▼──────┴───────┴───┐                                           ┌──┴────────┴───────▼─┐
│                   │───────────────── req ────────────────────►│                     │
│       demonote    │                                           │        demouser     │
│                   │◄──────────────── resp ────────────────────│                     │
└───────────────────┘                                           └─────────────────────┘
      thrift                                                           thrift
```

### Basic Features

- Hertz
  - Use `thrift` IDL to define `HTTP` interface
  - Use `hz` to generate code
  - Use `Hertz` binding and validate
  - Middleware
    - Use `requestid`, `jwt`, `obs-opentelemetry`, `recovery`
- Kitex
  - Use `thrift` IDL to define `RPC` interface
  - Use `kitex` to generate code
  - Middleware
    - Rate Limiting, Request Retry, Timeout Control, Connection Multiplexing
  - Tracing
    - use `obs-opentelemetry` and `jaeger` to tracing
  - Service Discovery and Register
    - Use [registry-etcd](https://github.com/kitex-contrib/registry-etcd) to discovery and register service

### Catalog Introduce

| catalog     | introduce               |
|-------------|-------------------------|
| handler     | HTTP handler            |
| service     | business logic          |
| rpc         | RPC call logic          |
| dal         | DB operation            |
| pack        | data pack               |
| pkg/mw      | RPC middleware          |
| pkg/consts  | constants               |
| pkg/errno   | customized error number |
| pkg/configs | SQL and Tracing configs |

## Quick Start

### Setup Basic Dependence

```shell
docker-compose up
```

### Run User RPC Server

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### Run Note RPC Server

```shell
cd cmd/note
sh build.sh
sh output/bootstrap.sh
```

### Run API Server

```shell
cd cmd/api
go run .
```

#### Snapshots

### Jaeger

visit `http://127.0.0.1:16686/` on  browser

#### Snapshots

![jaeger-tracing](./images/jarger-tracing.png)

![jaeger-architecture](./images/jaeger-architecture.png)

## API Requests

The following is a list of API requests and partial responses.

### Register

```shell
curl --location --request POST '127.0.0.1:8080/v2/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"lorain",
    "password":"123456"
}'
```

#### response

```javascript
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}

// failed
{
    "code": 10003,
    "message": "User already exists",
    "data": null
}
```

### Login

#### will return jwt token

```shell
curl --location --request POST '127.0.0.1:8080/v2/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"lorain",
    "password":"123456"
}'
```

#### response

```javascript
// successful
{
    "code": 0,
    "expire": "2022-12-3T01:56:46+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI1Mjg2MDYsImlkIjoxLCJvcmlnX2lhdCI6MTY0MjUyNTAwNn0.k7Ah9G4Enap9YiDP_rKr5HSzF-fc3cIxwMZAGeOySqU"
}

// failed
{
    "code": 10004,
    "message": "Authorization failed",
    "data": null
}
```

### Create Note

```shell
curl --location --request POST '127.0.0.1:8080/v2/note' \
--header 'Authorization: Bearer $token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"test title",
    "content":"test content"
}'
```

#### response

```javascript
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}

// failed
{
    "code": 10002,
    "message": "Wrong Parameter has been given",
    "data": null
}
```

### Query Note

```shell
curl --location --request GET '127.0.0.1:8080/v1/note/query?offset=0&limit=20&search_key=test' \
--header 'Authorization: Bearer $token'
```

#### response

```javascript
// successul
{
    "code": 0,
    "message": "Success",
    "data": {
        "notes": [
            {
                "note_id": 1,
                "user_id": 1,
                "username": "lorain",
                "user_avatar": "test",
                "title": "test title",
                "content": "test content",
                "create_time": 1642525063
            }
        ],
        "total": 1
    }
}

// failed
{
    "code":10002,
    "message":"Wrong Parameter has been given",
    "data":null
}
```

### Update Note

```shell
curl --location --request PUT '127.0.0.1:8080/v2/note/$note_id' \
--header 'Authorization: Bearer $token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"test",
    "content":"test"
}'
```

#### response

```javascript
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}

// failed
{
    "code":10001,
    "message":"strconv.ParseInt: parsing \"$note_id\": invalid syntax",
    "data":null
}
```

### Delete Note

```shell
curl --location --request DELETE '127.0.0.1:8080/v2/note/$note_id' \
--header 'Authorization: Bearer $token'
```

#### response

```javascript
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}

// failed
{
    "code":10001,
    "message":"strconv.ParseInt: parsing \"$note_id\": invalid syntax",
    "data":null
}
```
