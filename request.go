package glapi

import (
  "net/http"
  "net/url"
)

type Request struct {
  RawRequest *http.Request
  Headers    http.Header
  URL        *URL
  Query      url.Values
  Params     URLParams
  Method     string
}
