glapi
=====

[Express](http://expressjs.com) inspired web application framework for Go.

```go
package main

import(
  "github.com/inconceivableduck/glapi"
)

func main() {
  app := glapi.NewApp()
  
  app.Get("/", func(req *glapi.Request, res *glapi.Response, next func(err error)) {
    res.Send("Hello World")
  })
  
  app.Listen(":8080")
}
```

### Middleware
glapi is based around the concept of middleware. Each incoming request will pass through each middleware in the order they were added as long as the `next` function is called. If `next` is ever called with a non-nil error parameter propagation is immediately stopped.

```go
app.Use(func(req *Request, res *Response, next func(err error)) {
  // Do something.
  next(nil)
})
```

### Routing
Route handlers are syntactic sugar for adding middleware to the glapi application. Routes are added to the application by specifying the absolute path for matching. Named parameters can be specified by enclosing the value in `{}`. Routes are matched in the order they are added to the application. Typically `next` will not be called in route handlers.

```go
app.Get("/user/{id}", func(req *glapi.Request, res *glapi.Response, next func(err error)) {
  id := req.Params["id"]
  res.Send("Hello " + id)
})

app.Get("/user/{id}/project/{pid}", func(req *glapi.Request, res *glapi.Response, next func(err error)) {
  id := req.Params["id"]
  pid := req.Params["pid"]
  res.Send("Hello " + id + " with project " + pid)
})

app.Post("/user", func(req *glapi.Request, res *glapi.Response, next func(err error)) {
  res.Send("Posted to /user")
})
```
