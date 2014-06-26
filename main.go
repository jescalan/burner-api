// Package main is the command line interface to the burner api. It can be used
// to stop and start the server.
package main

import (
  "fmt"
  "net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
  fmt.Fprint(res, "hello world!")
}

func main() {
	http.HandleFunc("/", handler)
  http.ListenAndServe(":1111", nil)
}
