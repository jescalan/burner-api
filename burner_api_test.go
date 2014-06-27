package main

import (
  "testing"
  "fmt"
  "strings"
  "net/http"
)

func TestBasicLogic(t *testing.T) {
  req, err := http.NewRequest("POST", "http://example.com", strings.NewReader("wow"))
  if err != nil {
    t.Fail()
  }
  fmt.Println(req)
}
