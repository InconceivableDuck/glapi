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

app.Post("/user", func(req *glapi.Request, res *glapi.Response, next func(err error)) {
  res.Send("Posted to /user")
})
```

### BodyReader Middleware
The BodyReader middleware makes it easy to read the body content of incoming requests into memory.

```go
app.Use(glapi.BodyReader("application/json")
```

BodyReader only reads the body if the incoming content-type matches the specified value. If you'd like to read multiple types of content, apply the middleware multiple times.

```go
app.Use(glapi.BodyReader("application/json")
app.Use(glapi.BodyReader("text/plain")
```

BodyReader places the body contents in the `Request.Body` field.

### JSON Content
Combine the BodyReader middlware with Go's built-in [JSON package](http://golang.org/pkg/encoding/json/) for simple handling of JSON data.

```go
import(
  "encoding/json"
  "github.com/inconceivableduck/glapi"
)

type User struct {
  FirstName string
  LastName  string
}

func main() {
  app := glapi.NewApp()

  app.Use(glapi.BodyReader("application/json"))

  app.Post("/user", func(req *glapi.Request, res *glapi.Response, next func(err error)) {
    user := User{}
    err := json.Unmarshal(req.Body, &user)
    if err != nil {
      next(err)
    } else {
      res.send("OK")
    }
  })
}
```

### Errors

Every middleware includes a `next` function. If `next` is ever called with an error, propagation is immediately stopped and error handlers will be invoked. By default the application will respond with a 500 and the content of the error. Custom error handlers can be added much like regular middlware.

```go
app.Error(func(req *glapi.Request, res *glapi.Response, err error, next func()) {
  // Handle error.
  // Call next() to continue to next error handler if needed.
})
```

### Types

#### glapi.Request
The Request object holds useful information about the current request.

`Request.RawRequest` [*http.Request](http://golang.org/pkg/net/http/#Request) - The raw request object return from the underlying HTTP server.<br />
`Request.Headers` [http.Header](http://golang.org/pkg/net/http/#Header) - The request headers.<br />
`Request.URL` glapi.URL - The URL of the request.<br />
`Request.Query` [url.Values](http://golang.org/pkg/net/url/#Values) - Query string values.<br />
`Request.Params` glapi.URLParams - Map of named parameters and their values.<br />
`Request.Method` string - The request method: GET, POST, PUT, DELETE.<br />
`Request.Body` []byte - The request's body. Populated if using the BodyReader middleware.

#### glapi.Response
The Response objects provides convenience utilities to response to the incoming request.

`Response.Headers` [http.Header](http://golang.org/pkg/net/http/#Header) - The response headers. <br />
`Response.Send(r interface{})` - Sends content to the response. If the parameter is a string it is sent as plain text. If is is any other type it is [marshalled](http://golang.org/pkg/encoding/json/#MarshalIndent) and sent as JSON. <br />
`Response.Code(code int)` - Sets the HTTP status code of the response. Defaults to 200 (OK). <br />
