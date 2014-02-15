package glapi

import (
  "log"
  "net/url"
  "testing"
)

func TestRouterNoParams(t *testing.T) {

  r := NewRouter("/user")

  incomingRequest, _ := url.Parse("http://example.com/not-user")

  match, params := r.IsMatch(NewURL(incomingRequest))

  if match {
    t.Errorf("IsMatch boolean Expected %s Actual %s", false, match)
  }

  if len(params) > 0 {
    t.Errorf("IsMatch params length Expected %d Actual %d", 0, len(params))
  }

  incomingRequest, _ = url.Parse("http://example.com/user")

  match, params = r.IsMatch(NewURL(incomingRequest))

  log.Print(params)

  if !match {
    t.Errorf("IsMatch boolean Expected %s Actual %s", true, match)
  }

  if len(params) > 0 {
    t.Errorf("IsMatch params length Expected %d Actual %d", 0, len(params))
  }

  r = NewRouter("/user/project")

  incomingRequest, _ = url.Parse("http://example.com/user")

  match, params = r.IsMatch(NewURL(incomingRequest))

  if match {
    t.Errorf("IsMatch boolean Expected %s Actual %s", false, match)
  }

  if len(params) > 0 {
    t.Errorf("IsMatch params length Expected %d Actual %d", 0, len(params))
  }

  incomingRequest, _ = url.Parse("http://example.com/user/project")

  match, params = r.IsMatch(NewURL(incomingRequest))

  if !match {
    t.Errorf("IsMatch boolean Expected %s Actual %s", true, match)
  }

  if len(params) > 0 {
    t.Errorf("IsMatch params length Expected %d Actual %d", 0, len(params))
  }

}

func TestRouterOneParam(t *testing.T) {

  r := NewRouter("/user/{id}")

  incomingRequest, _ := url.Parse("http://example.com/not-user")

  match, params := r.IsMatch(NewURL(incomingRequest))

  if match {
    t.Errorf("IsMatch boolean Expected %s Actual %s", false, match)
  }

  if len(params) > 0 {
    t.Errorf("IsMatch params length Expected %d Actual %d", 0, len(params))
  }

  incomingRequest, _ = url.Parse("http://example.com/user")

  match, params = r.IsMatch(NewURL(incomingRequest))

  if match {
    t.Errorf("IsMatch boolean Expected %s Actual %s", false, match)
  }

  if len(params) > 0 {
    t.Errorf("IsMatch params length Expected %d Actual %d", 0, len(params))
  }

  incomingRequest, _ = url.Parse("http://example.com/user/something")

  match, params = r.IsMatch(NewURL(incomingRequest))

  if !match {
    t.Errorf("IsMatch boolean Expected %s Actual %s", true, match)
  }

  if params["id"] != "something" {
    t.Errorf("IsMatch id param Expected %s Actual %s", "something", params["id"])
  }

}

func TestRouterTwoParams(t *testing.T) {

  r := NewRouter("/user/{id}/project/{pid}")

  incomingRequest, _ := url.Parse("http://example.com/not-user")

  match, params := r.IsMatch(NewURL(incomingRequest))

  if match {
    t.Errorf("IsMatch boolean Expected %s Actual %s", false, match)
  }

  if len(params) > 0 {
    t.Errorf("IsMatch params length Expected %d Actual %d", 0, len(params))
  }

  incomingRequest, _ = url.Parse("http://example.com/user/userId/project/projectId")

  match, params = r.IsMatch(NewURL(incomingRequest))

  if !match {
    t.Errorf("IsMatch boolean Expected %s Actual %s", true, match)
  }

  if params["id"] != "userId" {
    t.Errorf("IsMatch id param Expected %s Actual %s", "userId", params["id"])
  }

  if params["pid"] != "projectId" {
    t.Errorf("IsMatch id param Expected %s Actual %s", "projectId", params["id"])
  }

}
