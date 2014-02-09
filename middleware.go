package glapi

import (
  "log"
)

// The main middleware interface.
// All middlewares must implement this interface.
type Middleware struct {
  handler Handler
  next    func()
}

// Function definition for middleware handler functions.
// These functions are called by a middleware's Invoke function.
type Handler func(req *Request, res *Response, next func(err error))

// Base middleware invoke call. Simply calls the current handler and invokes the
// next middleware. Does not do any processing on the request.
func (this *Middleware) Invoke(req *Request, res *Response, app *App, nextIdx int) {

  // Call the handler. This is typically code provided by the user.
  this.handler(req, res, func(err error) {

    if err != nil {
      // TODO: look for error middleware and invoke, otherwise return a 500.
      log.Print(err)
      return
    }

    if len(app.middleware) > nextIdx {
      nextNextIdx := nextIdx + 1
      // Invoke the next middleware in the list.
      app.middleware[nextIdx].Invoke(req, res, app, nextNextIdx)
    } else {
      res.Send("Cannot get " + req.URL.Path)
    }
  })
}
