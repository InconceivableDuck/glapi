package glapi

import (
  "net/url"
  "strings"
)

func NewURL(url *url.URL) *URL {
  newURL := &URL{}
  newURL.RawURL = url
  newURL.Parts = SplitURL(url)
  newURL.Path = url.Path

  return newURL
}

type URL struct {

  // The raw URL object from the HTTP server.
  RawURL *url.URL

  // Array of URL parts.
  // e.g. /user/1234 [0]="user", [1]="1234"
  Parts []string

  // The URL path. Does not include query params or domain.
  Path string
}

func SplitURL(url *url.URL) []string {
  // Split the string after removing the leading slash.
  return SplitPath(url.Path)
}

func SplitPath(path string) []string {
  return strings.Split(strings.TrimLeft(path, "/"), "/")
}
