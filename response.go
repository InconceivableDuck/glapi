package glapi

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Response struct {
  RawResponse http.ResponseWriter
  Headers     http.Header
  code        int
}

func NewResponse() *Response {
  res := &Response{}
  res.code = http.StatusOK
  return res
}

func (res *Response) Send(r interface{}) {

  switch str := r.(type) {
  case string:
    res.sendText(str)
    return
  }

  res.sendJson(r)
}

func (res *Response) Code(code int) {
  res.code = code
}

func (res *Response) sendText(text string) {
  res.Header.Set("content-type", "text/plain")
  res.RawResponse.WriteHeader(res.code)
  fmt.Fprintf(res.RawResponse, text)
}

func (res *Response) sendJson(r interface{}) {
  res.Header.Set("content-type", "application/json")
  res.RawResponse.WriteHeader(res.code)
  j, _ := json.MarshalIndent(r, "", "  ")
  fmt.Fprintf(res.RawResponse, string(j[:]))
}
