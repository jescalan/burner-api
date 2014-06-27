// Package main is the command line interface to the burner api. It can be used
// to stop and start the server.
package main

import (
  "fmt"
  "net/http"
  "bytes"
  "io"
  "encoding/json"
)

func handler(res http.ResponseWriter, req *http.Request) {
  params := GetParams(req.Body)
  fmt.Println(params.File)
  fmt.Fprint(res, "hello world!")
}

func main() {
	http.HandleFunc("/", handler)
  http.ListenAndServe(":1111", nil)
}

type Params struct {
  File string
}

// Given a reader, extract a string and parse it into a struct
func GetParams(body io.Reader) Params {
  buf := new(bytes.Buffer)
  buf.ReadFrom(body)

  var i Params
  json.Unmarshal(buf.Bytes(), &i)
  return i
}
