package glapi

import (
  "net/url"
  "testing"
)

func TestNewURL(t *testing.T) {
  input, _ := url.Parse("http://example.com/A/B/C?d=f&g=h")

  u := NewURL(input)

  if u.Path != "/A/B/C" {
    t.Errorf("u.Path Expected %s Actual %s", "/A/B/C", u.Path)
  }

  if u.RawURL != input {
    t.Errorf("u.RawURL is not equal to input")
  }

  if u.Parts[0] != "A" {
    t.Errorf("u.Parts[0] Expected %s Actual %s", "A", u.Parts[0])
  }

  if u.Parts[1] != "B" {
    t.Errorf("u.Parts[0] Expected %s Actual %s", "B", u.Parts[1])
  }

  if u.Parts[2] != "C" {
    t.Errorf("u.Parts[0] Expected %s Actual %s", "C", u.Parts[2])
  }
}
