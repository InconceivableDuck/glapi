package glapi

func Bootstrap() Handler {

  return func(req *Request, res *Response, next func(err error)) {

    req.Headers = req.RawRequest.Header
    req.URL = NewURL(req.RawRequest.URL)
    req.Query = req.RawRequest.URL.Query()
    req.Method = req.RawRequest.Method

    next(nil)
  }
}
