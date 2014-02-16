package glapi

import (
  "net/http"
  "testing"
)

// Empty response writer that implements http.ResponseWriter interface.
type TResponseWriter struct {
}

func (t TResponseWriter) Header() http.Header {
  return make(http.Header)
}

func (t TResponseWriter) Write(d []byte) (i int, err error) {
  return 0, nil
}

func (t TResponseWriter) WriteHeader(i int) {

}

func TestGetHandler(t *testing.T) {

  app := NewApp()

  called := false

  app.Get("/test", func(req *Request, res *Response, next func(err error)) {
    called = true
  })

  w := TResponseWriter{}
  r, _ := http.NewRequest("GET", "http://example.com/test", nil)

  app.ServeHTTP(w, r)

  if !called {
    t.Errorf("Get Handler Called: Expected %s Actual %s", true, called)
  }

  called = false

  w = TResponseWriter{}
  r, _ = http.NewRequest("POST", "http://example.com/test", nil)

  app.ServeHTTP(w, r)

  if called {
    t.Errorf("Get Handler Called: Expected %s Actual %s", false, called)
  }

}

func TestGetHandlerParams(t *testing.T) {

  app := NewApp()

  called := false

  app.Get("/test/{id}", func(req *Request, res *Response, next func(err error)) {
    called = true
  })

  w := TResponseWriter{}
  r, _ := http.NewRequest("GET", "http://example.com/test/123", nil)

  app.ServeHTTP(w, r)

  if !called {
    t.Errorf("Get Handler Called: Expected %s Actual %s", true, called)
  }

  w = TResponseWriter{}
  r, _ = http.NewRequest("GET", "http://example.com/not-match", nil)

  called = false

  app.ServeHTTP(w, r)

  if called {
    t.Errorf("Get Handler Called: Expected %s Actual %s", false, called)
  }
}

func TestPostHandler(t *testing.T) {

  app := NewApp()

  called := false

  app.Post("/test", func(req *Request, res *Response, next func(err error)) {
    called = true
  })

  w := TResponseWriter{}
  r, _ := http.NewRequest("POST", "http://example.com/test", nil)

  app.ServeHTTP(w, r)

  if !called {
    t.Errorf("Get Handler Called: Expected %s Actual %s", true, called)
  }
}

func TestPutHandler(t *testing.T) {

  app := NewApp()

  called := false

  app.Put("/test", func(req *Request, res *Response, next func(err error)) {
    called = true
  })

  w := TResponseWriter{}
  r, _ := http.NewRequest("PUT", "http://example.com/test", nil)

  app.ServeHTTP(w, r)

  if !called {
    t.Errorf("Get Handler Called: Expected %s Actual %s", true, called)
  }
}

func TestDeleteHandler(t *testing.T) {

  app := NewApp()

  called := false

  app.Delete("/test", func(req *Request, res *Response, next func(err error)) {
    called = true
  })

  w := TResponseWriter{}
  r, _ := http.NewRequest("DELETE", "http://example.com/test", nil)

  app.ServeHTTP(w, r)

  if !called {
    t.Errorf("Get Handler Called: Expected %s Actual %s", true, called)
  }
}
