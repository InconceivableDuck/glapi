package glapi

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Response struct {
  RawResponse http.ResponseWriter
}

func (res *Response) Send(r interface{}) {

  switch str := r.(type) {
  case string:
    res.sendText(str)
    return
  }

  res.sendJson(r)
}

func (res *Response) sendText(text string) {
  fmt.Fprintf(res.RawResponse, text)
}

func (res *Response) sendJson(r interface{}) {
  j, _ := json.MarshalIndent(r, "", "  ")
  res.sendText(string(j[:]))
}
