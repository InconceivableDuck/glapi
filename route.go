package glapi

type Route struct {
}

func (app *App) Get(route string, handler Handler) *Route {
  return addRoute(app, "GET", route, handler)
}

func (app *App) Post(route string, handler Handler) *Route {
  return addRoute(app, "POST", route, handler)
}

func (app *App) Put(route string, handler Handler) *Route {
  return addRoute(app, "PUT", route, handler)
}

func (app *App) Delete(route string, handler Handler) *Route {
  return addRoute(app, "DELETE", route, handler)
}

func addRoute(app *App, method string, path string, handler Handler) *Route {
  router := NewRouter(path)
  route := &Route{}

  app.Use(func(req *Request, res *Response, next func(err error)) {

    if req.Method != method {
      next(nil)
      return
    }

    match, params := router.IsMatch(req.URL)

    if !match {
      next(nil)
      return
    }

    req.Params = params

    handler(req, res, next)
  })

  return route
}
