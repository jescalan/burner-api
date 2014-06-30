package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nu7hatch/gouuid"
)

// - If not a POST, returns a 404
// - Gets the contents of a POSTed file
// - Creates a UUID associated with it
// - Downloads the file locally, named UUID.tar.gz
// - Responds with the UUID
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

// - Gets the "file" param from the request body
// - If there is a file with that name + .tar.gz, respond with that file
// - Then delete that file. Yolo!
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
}

func main() {
	http.HandleFunc("/new", HostFile)
	http.HandleFunc("/", ServeFile)
	http.ListenAndServe(":1111", nil)
}

// writes a 404 response to a passed in (pointer to a) response
func fourohfour(res http.ResponseWriter) {
	res.WriteHeader(404)
	fmt.Fprint(res, "not found")
}

// given a request, read the body and return as a byte slice
func getFileContents(req *http.Request) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	return buf.Bytes()
}

// struct used to return from getFilename, below
type params struct {
	File string
}

// Given a request, extract the body and pull the "file" param into a struct
func getFilename(req *http.Request) (p params, err error) {
	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(content, &p)
	return
}

// given a file name, create ./files/NAME.tar.gz
func createFile(id string) (file *os.File, err error) {
	dirname, err := dirname()
	if err != nil {
		return
	}

	file, err = os.Create(filepath.Join(dirname, "files", id+".tar.gz"))

	return
}

// get the name of the directory this file is in
func dirname() (dirname string, err error) {
	dirname, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}
