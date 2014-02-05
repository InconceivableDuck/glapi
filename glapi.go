package glapi

import (
	"log"
	"net/http"
)

func NewApp() App {
	app := App{}
	app.middleware = make([]Middleware, 0)
	return app
}

type App struct {
	middleware []Middleware
}

type Request struct {
	RawRequest *http.Request
}

type Handler func(req *Request, res *Response, next func(err error))

type Middleware struct {
	handler Handler
	next    func()
}

func (mid *Middleware) Invoke(req *Request, res *Response, app *App, nextIdx int) {

	mid.handler(req, res, func(err error) {

		if err != nil {
			log.Print(err)
			return
		}

		if len(app.middleware) > nextIdx {
			nextNextIdx := nextIdx + 1
			app.middleware[nextIdx].Invoke(req, res, app, nextNextIdx)
		}
	})
}

func (app *App) Get(route string, handler Handler) {

	newMid := Middleware{}
	newMid.handler = handler

	app.middleware = append(app.middleware, newMid)
}

func (app *App) HandleRequest(req *Request, res *Response) {
	app.middleware[0].Invoke(req, res, app, 1)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	req := &Request{r}
	res := &Response{w}

	app.HandleRequest(req, res)
}

func (app *App) Listen(addr string) {
	http.ListenAndServe(addr, app)
}
