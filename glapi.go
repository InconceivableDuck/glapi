package glapi

import (
  "net/http"
)

func NewApp() App {
  app := App{}
  app.middleware = make([]Middleware, 0)
  app.errorMiddleware = make([]ErrorMiddleware, 0)
  app.Use(Bootstrap())
  return app
}

type App struct {
  // All actions are defined by middleware.
  // The request will flow through every middleware in the order
  // they were added as long as next() continues to be called.
  middleware []Middleware

  // Error handlers. Just like regular middleware, these will be
  // called in order as long as next() continues to be called.
  // If next is called with an error, the default error handler will be called.
  errorMiddleware []ErrorMiddleware
}

func (app *App) Use(handler Handler) {
  newMid := Middleware{}
  newMid.handler = handler
  app.middleware = append(app.middleware, newMid)
}

func (app *App) Error(handler ErrorHandler) {
  newMid := ErrorMiddleware{}
  newMid.handler = handler
  app.errorMiddleware = append(app.errorMiddleware, newMid)
}

func (app *App) HandleError(req *Request, res *Response, err error) {
  if len(app.errorMiddleware) > 0 {
    app.errorMiddleware[0].Invoke(req, res, app, err, 1)
  } else {
    app.DefaultErrorHandler(req, res, err)
  }
}

func (app *App) HandleRequest(req *Request, res *Response) {
  app.middleware[0].Invoke(req, res, app, 1)
}

func (app *App) DefaultHandler(req *Request, res *Response) {
  res.Send("Cannot " + req.Method + " " + req.URL.Path)
}

func (app *App) DefaultErrorHandler(req *Request, res *Response, err error) {
  res.Code(http.StatusInternalServerError)
  res.Send(err.Error())
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  req := &Request{}
  req.RawRequest = r

  res := NewResponse()
  res.RawResponse = w

  app.HandleRequest(req, res)
}

func (app *App) Listen(addr string) {
  http.ListenAndServe(addr, app)
}
