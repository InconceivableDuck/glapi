package glapi

// The main middleware interface.
// All middlewares must implement this interface.
type Middleware struct {
  handler Handler
  next    func()
}

type ErrorMiddleware struct {
  handler ErrorHandler
  next    func()
}

// Function definition for middleware handler functions.
// These functions are called by a middleware's Invoke function.
type Handler func(req *Request, res *Response, next func(err error))

// Function definition for error middlware.
type ErrorHandler func(req *Request, res *Response, err error, next func())

// Invokes the error middlewares.
func (this *ErrorMiddleware) Invoke(req *Request, res *Response, app *App, err error, nextIdx int) {

  this.handler(req, res, err, func() {
    if len(app.errorMiddleware) > nextIdx {
      nextNextIdx := nextIdx + 1
      app.errorMiddleware[nextIdx].Invoke(req, res, app, err, nextNextIdx)
    } else {
      app.DefaultErrorHandler(req, res, err)
    }
  })
}

// Base middleware invoke call. Simply calls the current handler and invokes the
// next middleware. Does not do any processing on the request.
func (this *Middleware) Invoke(req *Request, res *Response, app *App, nextIdx int) {

  // Call the handler. This is typically code provided by the user.
  this.handler(req, res, func(err error) {

    if err != nil {
      app.HandleError(req, res, err)
      return
    }

    if len(app.middleware) > nextIdx {
      nextNextIdx := nextIdx + 1
      // Invoke the next middleware in the list.
      app.middleware[nextIdx].Invoke(req, res, app, nextNextIdx)
    } else {
      app.DefaultHandler(req, res)
    }
  })
}
