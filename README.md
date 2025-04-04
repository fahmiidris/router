## Golang Router

### Group

This is a simple router for golang. It is forked from [httprouter](https://github.com/julienschmidt/httprouter) and modified to support router group from PR [#89](https://github.com/julienschmidt/httprouter/pull/89).

### Middleware

This package includes middleware support. Middleware functions can be chained to execute before or after the main handler function for a route. This allows you to perform tasks like authentication, logging, or request modification.

### Usage

```go
go get github.com/fahmiidris/router
```
