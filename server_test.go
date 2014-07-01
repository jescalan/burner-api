package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var id, testPath string

func TestHostFile(t *testing.T) {
	data := strings.NewReader("wow")
	req, err := http.NewRequest("POST", "http://example.com", data)
	assert.Nil(t, err)
	res := httptest.NewRecorder()

	HostFile(res, req)
	assert.Equal(t, res.Code, 200)
	assert.IsType(t, res.Body, new(bytes.Buffer))

	content, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	id = string(content)
	testPath = filepath.Join(dir(), "files", id+".tar.gz")

	_, err = os.Stat(testPath)
	assert.False(t, os.IsNotExist(err))

	content, err = ioutil.ReadFile(testPath)
	assert.Nil(t, err)
	assert.Equal(t, string(content), "wow")
}

func TestHostFileNotPost(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	assert.Nil(t, err)
	res := httptest.NewRecorder()

	HostFile(res, req)
	assert.Equal(t, res.Code, 404)
}

func TestServeFile(t *testing.T) {
	data := strings.NewReader("{\"File\": \"" + id + "\" }")
	req, err := http.NewRequest("GET", "http://example.com", data)
	assert.Nil(t, err)
	res := httptest.NewRecorder()

	ServeFile(res, req)

	assert.Equal(t, res.Code, 200)
	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, string(body), "wow")

	_, err = os.Stat(testPath)
	assert.True(t, os.IsNotExist(err))
}

// dir is a utility function used to quickly grab the current directory name
func dir() string {
	_, filename, _, _ := runtime.Caller(1)
	d, _ := filepath.Abs(filepath.Dir(filename))
	return d
}
