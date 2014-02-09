package glapi

import (
	"net/http"
)

func NewApp() App {
	app := App{}
	app.middleware = make([]Middleware, 0)
	app.Use(Bootstrap())
	return app
}

type App struct {
	// All actions are defined by middleware.
	// The request will flow through every middleware in the order
	// they were added as long as next() continues to be called.
	middleware []Middleware
}

func (app *App) Use(handler Handler) {
	newMid := BaseMiddleware{}
	newMid.handler = handler
	app.middleware = append(app.middleware, newMid)
}

func (app *App) HandleRequest(req *Request, res *Response) {
	app.middleware[0].Invoke(req, res, app, 1)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	req := &Request{}
	req.RawRequest = r

	res := &Response{}
	res.RawResponse = w

	app.HandleRequest(req, res)
}

func (app *App) Listen(addr string) {
	http.ListenAndServe(addr, app)
}
