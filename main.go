// Package main is the command line interface to the burner api. It can be used
// to stop and start the server.
package main

import (
	"bytes"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// - if not a POST, return a 404
// - Gets the contents of a POSTed file
// - Creates a UUID associated with it
// - Downloads the file locally, named UUID.tar.gz
// - Responds with the UUID
func hostFile(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(404)
		fmt.Fprint(res, "not found")
		return
	}

	contents := GetBody(req)

	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	file, err := CreateFile(id.String())
	if err != nil {
		log.Fatal(err)
	}

	file.Write(contents)

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

// given a file name, create ./files/NAME.tar.gz
func CreateFile(id string) (*os.File, error) {
	dirname, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}

	fPath := filepath.Join(dirname, "files", id+".tar.gz")
	file, err := os.Create(fPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
