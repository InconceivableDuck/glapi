package glapi

import (
  "net/http"
  "testing"
)

type TestObject struct {
  A string
  B string
}

type TestResponseWriter struct {
  Code *int
}

func (t TestResponseWriter) Header() http.Header {
  return make(http.Header)
}

func (t TestResponseWriter) Write(p []byte) (int, error) {
  return 0, nil
}

func (t TestResponseWriter) WriteHeader(h int) {
  *t.Code = h
}

func TestCode(t *testing.T) {

  rw := TestResponseWriter{}
  code := 0
  rw.Code = &code

  r := NewResponse()
  r.RawResponse = rw
  r.Headers = make(http.Header)

  r.Code(500)
  r.Send("Hello")

  if r.Headers.Get("content-type") != "text/plain" {
    t.Errorf("Send: Expected %s Actual %s", "text/plain", r.Headers.Get("content-type"))
  }

  if *rw.Code != 500 {
    t.Errorf("Send Code: Expected %d Actual %d", 500, rw.Code)
  }
}

func TestSendText(t *testing.T) {

  rw := TestResponseWriter{}
  code := 0
  rw.Code = &code

  r := NewResponse()
  r.RawResponse = rw
  r.Headers = make(http.Header)

  r.Send("Hello")

  if r.Headers.Get("content-type") != "text/plain" {
    t.Errorf("Send: Expected %s Actual %s", "text/plain", r.Headers.Get("content-type"))
  }

  if *rw.Code != 200 {
    t.Errorf("Send Code: Expected %d Actual %d", 200, rw.Code)
  }
}

func TestSendJson(t *testing.T) {

  rw := TestResponseWriter{}
  code := 0
  rw.Code = &code

  r := NewResponse()
  r.RawResponse = rw
  r.Headers = make(http.Header)

  tObj := TestObject{"1", "2"}

  r.Send(tObj)

  if r.Headers.Get("content-type") != "application/json" {
    t.Errorf("Send: Expected %s Actual %s", "application/json", r.Headers.Get("content-type"))
  }

  if *rw.Code != 200 {
    t.Errorf("Send Code: Expected %d Actual %d", 200, rw.Code)
  }
}
