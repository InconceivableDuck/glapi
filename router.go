package glapi

import (
  "log"
  "strings"
)

type Router struct {
  Route string
  Parts []URLParts
}

type URLParts struct {
  Name  string
  Match string
}

type URLParams map[string]string

func NewRouter(route string) *Router {
  r := &Router{}
  r.Route = route
  r.Parts = parseRoute(route)

  return r
}

func parseRoute(route string) []URLParts {
  parts := strings.Split(route, "/")
  urlParts := make([]URLParts, len(parts))

  for idx, part := range parts {

    urlParts[idx].Match = part

    if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
      urlParts[idx].Match = "*"
      urlParts[idx].Name = part
      name := part
      name = strings.TrimLeft(name, "{")
      name = strings.TrimRight(name, "}")
      urlParts[idx].Name = name
    }
  }

  log.Print(urlParts)

  return urlParts
}

func (this *Router) IsMatch(url *URL) (isMatch bool, params URLParams) {

  params = make(URLParams)

  if len(url.Parts) != len(this.Parts) {
    return false, params
  }

  for idx, part := range url.Parts {
    if this.Parts[idx].Match == "*" {
      params[this.Parts[idx].Name] = part
      continue
    }

    if this.Parts[idx].Match == part {
      params[this.Parts[idx].Name] = part
      continue
    }

    return false, params
  }

  return true, params
}
