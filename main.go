// Package main provides a simple http server that can store a file and serve it
// for download once before deleting it.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/nu7hatch/gouuid"
)

// main parses a single port flag, sets up the routes, and starts the server on
// the specified port or 1111 by default.
func main() {
	var port = flag.Int("p", 1111, "port on which the server should start")
	flag.Parse()

	http.HandleFunc("/new", HostFile)
	http.HandleFunc("/", ServeFile)
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

// HostFile expects a POST with a file. It grabs the file's contents, generates
// an id, saves the file locally as that id, then returns the id.
func HostFile(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fourohfour(res)
		return
	}

	contents := getFileContents(req)

	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	file, err := createFile(id.String())
	if err != nil {
		log.Fatal(err)
	}

	file.Write(contents)

	fmt.Fprint(res, id.String())
}

// ServeFile catches all other requests. It expects a "file" param in the
// request body, specifying an id. It searches for a file named with that id,
// and if it exists, serves that file then deletes it. If not, 404.
func ServeFile(res http.ResponseWriter, req *http.Request) {
	params, err := getFilename(req)
	if err != nil {
		fourohfour(res)
		return
	}

	dirname, err := dirname()
	if err != nil {
		log.Fatal(err)
	}

	fPath := filepath.Join(dirname, "files", params.File+".tar.gz")
	content, err := ioutil.ReadFile(fPath)
	if err != nil {
		fourohfour(res)
		return
	}

	fmt.Fprint(res, string(content))

	err = os.Remove(fPath)
	if err != nil {
		log.Fatal(err)
	}
}

// fourohfour writes a 404 response to a passed in http.ResponseWriter.
func fourohfour(res http.ResponseWriter) {
	res.WriteHeader(404)
	fmt.Fprint(res, "not found")
}

// getFileContents reads the given request body and returns it as a byte slice.
func getFileContents(req *http.Request) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	return buf.Bytes()
}

// params is just a struct used to return from the getFilename function below.
type params struct {
	File string
}

// getFilename extracts the request body and pull the "file" param into a struct
// containing the file name requested as a string.
func getFilename(req *http.Request) (p params, err error) {
	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(content, &p)
	return
}

// createFile creates a blank file using the given name at ./files/NAME.tar.gz.
func createFile(id string) (file *os.File, err error) {
	dirname, err := dirname()
	if err != nil {
		return
	}

	file, err = os.Create(filepath.Join(dirname, "files", id+".tar.gz"))

	return
}

// dirname gets the name of the directory this file is in.
func dirname() (dirname string, err error) {
	_, filename, _, _ := runtime.Caller(1)
	dirname, err = filepath.Abs(filepath.Dir(filename))
	return
}
