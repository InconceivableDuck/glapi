package glapi

func (app *App) Get(route string, handler Handler) {
	addRoute(app, "GET", route, handler)
}

func (app *App) Post(route string, handler Handler) {
	addRoute(app, "POST", route, handler)
}

func (app *App) Put(route string, handler Handler) {
	addRoute(app, "PUT", route, handler)
}

func (app *App) Delete(route string, handler Handler) {
	addRoute(app, "DELETE", route, handler)
}

func addRoute(app *App, method string, route string, handler Handler) {
	router := NewRouter(route)

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
}
