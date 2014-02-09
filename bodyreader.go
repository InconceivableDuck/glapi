package glapi

import (
  "io/ioutil"
)

func BodyReader(contentType string) Handler {

  return func(req *Request, res *Response, next func(err error)) {

    // Only care about text-based content.
    if req.Headers.Get("content-type") != contentType {
      next(nil)
      return
    }

    content, err := ioutil.ReadAll(req.RawRequest.Body)

    if err != nil {
      next(err)
      return
    }

    req.Body = content

    next(nil)
  }
}
