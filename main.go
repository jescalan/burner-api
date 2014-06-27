// Package main is the command line interface to the burner api. It can be used
// to stop and start the server.
package main

import (
  "fmt"
  "net/http"
  "bytes"
  "os"
  "log"
  "path/filepath"
  "github.com/nu7hatch/gouuid"
)

func hostFile(res http.ResponseWriter, req *http.Request) {
  // get the contents of the POSTed file
  contents := GetBody(req)

  // create a new uuid
  id, err := uuid.NewV4()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(id)

  // create a new file named the uuid .tar.gz in the 'files' directory
  dirname, _ := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatal(err)
  }
  fPath := filepath.Join(dirname, "files", id.String() + ".tar.gz")
  file, err := os.Create(fPath)
  if err != nil {
    log.Fatal(err)
  }

  // write the contents of the POSTed file to it
  file.Write(contents)

  // return the id as the response
  fmt.Fprint(res, id.String())
}

func main() {
	http.HandleFunc("/new", hostFile)
  http.ListenAndServe(":1111", nil)
}

// Given a reader, extract a string and parse it into a struct
func GetBody(req *http.Request) []byte {
  buf := new(bytes.Buffer)
  buf.ReadFrom(req.Body)
  return buf.Bytes()
}
